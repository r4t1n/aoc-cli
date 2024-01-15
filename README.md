# aoc-cli

[![Go Report Card](https://goreportcard.com/badge/github.com/r4t1n/aoc-cli)](https://goreportcard.com/report/github.com/r4t1n/aoc-cli)

Command line interface for [Advent of Code](https://adventofcode.com/about)

## Installing

### Building from source

**[Make sure go is installed for building aoc-cli](https://go.dev/doc/install)**

```
git clone https://github.com/r4t1n/aoc-cli.git && cd aoc-cli && make && make install
```

## Session cookie

Everyone gets a different puzzle input, to download your input you need to provide your session cookie.

To obtain your session cookie, login to the [Advent of Code](https://adventofcode.com) website and inspect the `session` cookie to get it's value - [see instructions](https://www.cookieyes.com/blog/how-to-check-cookies-on-your-website-manually).

The session cookie (a long hex string) must be provided in a file called `.adventofcode.session` inside your home directory. 

For example:

`/home/ratin/.adventofcode.session`

## Command line flags

### `-day (int)`

Sets the day used for the date, replace int with the day

### `-year (int)`

Sets the year used for the date, replace int with the year
