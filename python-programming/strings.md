# Strings
{id: python-strings}


## Single quoted and double quoted strings
{id: single-quote-double-quote}

![](examples/strings/quotes.py)
![](examples/strings/quotes.out)

## Long lines
{id: long-lines}

![](examples/strings/other.py)
![](examples/strings/other.out)


## Triple quoted strings (multiline)
{id: triple-quoted-strings}
{i: """}
{i: '''}

![](examples/strings/triple_quotes.py)

Can spread multiple lines.

```
first row
second row
third row
```


## String length (len)
{id: string-length}
{i: len}

![](examples/strings/length.py)


## String repetition
{id: string-repetition}
{i: *}
{i: +}

![](examples/strings/repetition.py)


## String slice (instead of substr)
{id: string-slice}
{i: slice}
{i: substr}
{i: [:]}
{i: :}

![](examples/strings/string_slice.py)


## Change a string
{id: string-change}
{i: immutable}

![](examples/strings/string_change.py)

{aside}
Replace part of a string
{/aside}


* Strings in Python are **immutable** - they never change.


## How to change a string
{id: how-to-change-a-string}

![](examples/strings/change_a_string.py)


## String copy
{id: string-copy}

![](examples/strings/string_copy.py)

{aside}
When assigning a variable pointing a string, the new variable is pointing to the same string..
If we then assign some other string to either of the variables, then they will point to two different strings.
{/aside}


## String functions and methods (len, upper, lower)
{id: string-functions}
{i: len}
{i: upper}
{i: lower}

![](examples/strings/functions.py)

* Type **dir("")** in the REPL to get the list of string methods.
* List of [built-in functions](http://docs.python.org/library/functions.html).
* List of [string methods](http://docs.python.org/library/string.html).


## index in string
{id: index-in-string}
{i: index}
{i: ValueError}

![](examples/strings/index.py)
![](examples/strings/index.py.out)


## index in string with range
{id: index-in-string-range}
{i: index}

![](examples/strings/index2.py)
![](examples/strings/index2.py.out)


## rindex in string with range
{id: rindex-in-string-range}
{i: rindex}

![](examples/strings/rindex.py)
![](examples/strings/rindex.py.out)


## find in string
{id: find-in-string}
{i: find}
{i: rfind}

Alternatively use find and rfind that will return -1 instead of raising an exception.

![](examples/strings/find.py)


## Find all in the string
{id: find-all-in-the-string}

Later, when we learned loops.


## in string
{id: in-syntax-on-string}
{i: in}

Check if a substring is **in** the string?

![](examples/strings/in.py)
![](examples/strings/in.out)


## index if in string
{id: location-if-in-string}
{i: index}
{i: in}

![](examples/strings/location_if_in.py)


## Encodings: ASCII, Windows-1255, Unicode
{id: ascii-unicode}

* [ASCII](https://en.wikipedia.org/wiki/ASCII)
* [Hebrew Character](https://en.wikipedia.org/wiki/Hebrew_character)
* [Windows-1255](https://en.wikipedia.org/wiki/Windows-1255)
* [Unicode (UTF-8)](https://en.wikipedia.org/wiki/Unicode)



## raw strings
{id: raw-strings}
{i: r}
![](examples/strings/raw.py)

Escape sequences are kept intact and not escaped. Used in regexes.

## ord
{id: ord}
{i: ord}

* [ord](https://docs.python.org/3/library/functions.html#ord)

![](examples/strings/ord_chars.py)

## ord in a file
{id: ord-in-file}
{i: ord}
![](examples/strings/ord.py)


## chr - number to character
{id: chr}
{i: chr}

* [chr](https://docs.python.org/3/library/functions.html#chr)

![](examples/strings/chr.py)
![](examples/strings/chr.out)


## Exercise: one string in another string
{id: exercise-one-string-in-another-string}

Write script that accepts two strings and tells if one of them can be found in the other and where?


## Exercise: to ASCII CLI
{id: exercise-to-ascii-cli}

Write script that gets a character on the command line and prints out the ascii code of it.

Maybe even:

Write script that gets a string on the command line and prints out the ascii code of each character.


## Exercise: from ASCII CLI
{id: exercise-from-ascii-cli}

Write script that accepts a number on the command line and prints the character represented by that number.


## Solution: one string in another string
{id: solution-one-string-in-another-string}

![](examples/strings/in_string.py)


## Solution: to ASCII CLI
{id: solution-to-ascii-cli}


![](examples/strings/to_ascii_char.py)
![](examples/strings/to_ascii.py)

## Solution: from ASCII CLI
{id: solution-from-ascii-cli}

![](examples/strings/from_ascii.py)

