# aoc-cli

[![Go Report Card](https://goreportcard.com/badge/github.com/r4t1n/aoc-cli)](https://goreportcard.com/report/github.com/r4t1n/aoc-cli)

Command line interface for Advent of Code

## Installing

### Building from source

**Make sure go is installed for building aoc-cli**

```
git clone https://github.com/r4t1n/aoc-cli.git && ./aoc-cli/build.sh
```

or

```
git clone https://github.com/r4t1n/aoc-cli.git && cd aoc-cli && go build
```

**then** move the binary to the desired location (make sure that location is in the PATH)

## Session cookie

Everyone gets a different puzzle input, to download your input you need provide your session cookie.

To obtain your session cookie, login to the [Advent of Code](https://adventofcode.com) website and inspect the `session` cookie to get it's value - [see instructions](https://www.cookieyes.com/blog/how-to-check-cookies-on-your-website-manually).

The session cookie (a long hex string) must be provided in a file called `.adventofcode.session` inside your home directory. For example:

`/home/r4t1n/.adventofcode.session`