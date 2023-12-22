# aoc-cli

[![Go Report Card](https://goreportcard.com/badge/github.com/r4t1n/aoc-cli)](https://goreportcard.com/report/github.com/r4t1n/aoc-cli)

Command line interface for Advent of Code

## Installing

### Compile from source

```
git clone https://github.com/r4t1n/aoc-cli.git
```

```
cd aoc-cli
```

```
go build -o aoc
```

## Session cookie

Everyone gets a different puzzle input, to download your input you need provide your session cookie.

To obtain your session cookie, login to the [Advent of Code](https://adventofcode.com) website and inspect the `session` cookie to get it's value - [see instructions](https://www.cookieyes.com/blog/how-to-check-cookies-on-your-website-manually).

The session cookie (a long hex string) must be provided in a single line (no
line breaks) in a file called `.adventofcode.session` inside your home directory. For example:

`/home/r4t1n/.adventofcode.session`