NoshGobble
================================================================================

NoshGobble is a calorie counter. I know, I know, there are a million and one
calorie counters out there already. But this one is a little bit different. My
goal with this project was to be able to enter your ingredient list in free form
and have the app parse the ingredients and give you back the aggregated
nutritional data. Think of it like Shazam for recipes, except, instead of
retrieving the song name and artist info, you're getting the nutrition label for
a recipe you like.

![Screenshot](/public/images/screencast.gif?raw=true "Screenshot")

This project still has a long way to go. Finding the nutritional data that the
user would expect is harder than I expected. I plan to add some machine learning
to get this right.

Statistics
--------------------------------------------------------------------------------
```
http://cloc.sourceforge.net v 1.60  T=0.91 s (24.3 files/s, 1513.8 lines/s)
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                              13            140             63            977
Javascript                       2              4             16             65
SQL                              3              1              0             39
HTML                             1              0              0             25
CSS                              2              0             16             16
Bourne Shell                     1              2              1              7
-------------------------------------------------------------------------------
SUM:                            22            147             96           1129
-------------------------------------------------------------------------------
```

Not bad for 1 weekend's worth of coding in a new language.

Installation
--------------------------------------------------------------------------------

```
$ go get github.com/dbalmain/go-sqlite3
$ go get github.com/a2800276/porter
$ go get github.com/gophergala/noshgobble
$ ./noshgobble -reset
$ ./noshgobble -serve
```
