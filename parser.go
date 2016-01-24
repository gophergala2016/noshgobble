package main

import (
	"errors"
	"io"
	"strconv"
	"strings"
)

// Parser represents a parser.
type Parser struct {
	s   *Scanner
	buf struct {
		tok Token  // last read token
		lit string // last read literal
		n   int    // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) ignoreOfA() {
	tok, _ := p.scan()
	for tok == OF || tok == A || tok == WHITESPACE {
		tok, _ = p.scan()
	}

	p.unscan()
}

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnore() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WHITESPACE {
		tok, lit = p.scan()
	}
	return
}

func (p *Parser) scanFraction() (quantity float64, err error) {
	tok, lit := p.scanIgnore()
	if tok == FRACTION {
		switch lit {
		case "half":
			quantity = 0.5
		case "third", "thirds":
			quantity = 0.333
		case "quarter", "quarters":
			quantity = 0.25
		case "eighth", "eighths":
			quantity = 0.125
		default:
			return 0.0, errors.New("Couldn't parse quantity")
		}
		p.ignoreOfA()
	} else if tok == A {
		return p.scanFraction()
	} else {
		return 0.0, errors.New("Couldn't parse quantity")
	}
	return
}

func (p *Parser) scanQuantity() (quantity float64, err error) {
	tok, lit := p.scanIgnore()
	if tok == NUMBER {
		quantity, err = strconv.ParseFloat(lit, 64)
		checkErr(err)
		fraction, e := p.scanFraction()
		if e == nil {
			quantity *= fraction
		} else {
			p.unscan()
		}
	} else {
		p.unscan()
		quantity, err = p.scanFraction()
	}

	return
}

func (p *Parser) scanUnit() (unit Unit, err error) {
	tok, lit := p.scanIgnore()
	if tok == UNIT {
		unit = getUnitModel(lit).id
	} else {
		unit = COUNT
	}

	return
}

func (p *Parser) scanTerms() string {
	terms := make([]string, 0)
	for tok, lit := p.scan(); tok != EOF; tok, lit = p.scan() {
		if tok == FOOD {
			terms = append(terms, strings.ToLower(lit))
		}
	}
	return strings.Join(terms, " OR ")
}

func (p *Parser) Parse() (item *FoodItem, err error) {
	item = &FoodItem{}
	if item.quantity, err = p.scanQuantity(); err != nil {
		return nil, err
	}
	if item.unit, err = p.scanUnit(); err != nil {
		return nil, err
	}
	if item.terms = p.scanTerms(); item.terms == "" {
		return nil, errors.New("No food terms recognized")
	}
	return item, nil
}
