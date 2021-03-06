package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"strings"
	"unicode"
)

// Token represents a lexical token.
type Token int

const (
	// Special tokens
	EOF Token = iota
	WHITESPACE

	// Number
	NUMBER

	// Punctuation
	SLASH
	COMMA
	PUNC

	// Units and Measures
	UNIT     // g, grams, litres, etc.
	FRACTION // half, third, quarter

	// Food words
	FOOD // Bread, chicken, salt

	// KEYWORDS
	OF // eg 1L of water
	A  // eg half a cup of water

	// Other
	WORD
)

func (t Token) String() string {
	switch t {
	case EOF:
		return "EOF"
	case WHITESPACE:
		return "WHITESPACE"
	case NUMBER:
		return "NUMBER"
	case SLASH:
		return "SLASH"
	case COMMA:
		return "COMMA"
	case PUNC:
		return "PUNC"
	case UNIT:
		return "UNIT"
	case FRACTION:
		return "FRACTION"
	case FOOD:
		return "FOOD"
	case OF:
		return "OF"
	case A:
		return "A"
	case WORD:
		return "WORD"
	default:
		log.Fatalf("Error: you forgot to add the string for Type %d\n", t)
	}
	return ""
}

const eof = rune(0)

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// scanWhitespace consumes all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		}
	}

	return WHITESPACE, " "
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	ident := buf.String()
	if unit := getUnitName(buf.String()); unit != "" {
		return UNIT, unit
	}

	switch strings.ToLower(ident) {
	case "a":
		return A, "a"
	case "of":
		return OF, "of"
	case "half":
		return FRACTION, "half"
	case "third", "thirds":
		return FRACTION, "third"
	case "quarter", "quarters":
		return FRACTION, "quarter"
	case "eighth", "eighths":
		return FRACTION, "eighth"
	case "one":
		return NUMBER, "1"
	case "two":
		return NUMBER, "2"
	case "three":
		return NUMBER, "3"
	case "four":
		return NUMBER, "4"
	case "five":
		return NUMBER, "5"
	case "six":
		return NUMBER, "6"
	case "seven":
		return NUMBER, "7"
	case "eight":
		return NUMBER, "8"
	case "nine":
		return NUMBER, "9"
	case "ten":
		return NUMBER, "10"
	case "eleven":
		return NUMBER, "11"
	case "twelve":
		return NUMBER, "12"
	default:
		if IsFoodTerm(ident) {
			return FOOD, ident
		}
		// Otherwise return as a regular identifier.
		return WORD, ident
	}
}

// scanIdent consumes the current rune and all contiguous digit runes
func (s *Scanner) scanNumber() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !unicode.IsDigit(ch) && ch != '.' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// Otherwise return as a regular identifier.
	return NUMBER, buf.String()
}

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an ident or reserved word.
	if isWhitespace(ch) {
		return s.scanWhitespace()
	} else if unicode.IsDigit(ch) {
		s.unread()
		return s.scanNumber()
	} else if unicode.IsLetter(ch) {
		s.unread()
		return s.scanIdent()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return EOF, ""
	case '/':
		return SLASH, string(ch)
	case ',':
		return COMMA, string(ch)
	}

	return PUNC, string(ch)
}
