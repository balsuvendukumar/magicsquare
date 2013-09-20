magicsquare
===========

Package `magicsquare` is a [Go](http://golang.org) library and command line
utility, to generate [magic
squares](http://en.wikipedia.org/wiki/Magic_square).


## Status

[![Build Status](https://travis-ci.org/jmcvetta/magicsquare.png)](https://travis-ci.org/jmcvetta/magicsquare)
[![Build Status](https://drone.io/github.com/jmcvetta/magicsquare/status.png)](https://drone.io/github.com/jmcvetta/magicsquare/latest)
[![Coverage Status](https://coveralls.io/repos/jmcvetta/magicsquare/badge.png)](https://coveralls.io/r/jmcvetta/magicsquare)


## Usage

A working [Go](http://golang.org) installation is required to build `magicsquare`.

```bash
# We go get 'magicsquare/magicsquare' to install the command line utility
$ go get github.com/jmcvetta/magicsquare/magicsquare

$ magicsquare 7
   20   12    4   45   37   29   28
   11    3   44   36   35   27   19
    2   43   42   34   26   18   10
   49   41   33   25   17    9    1
   40   32   24   16    8    7   48
   31   23   15   14    6   47   39
   22   21   13    5   46   38   30

```

## Documentation

See GoDoc for [automatically generated API
documentation](http://godoc.org/github.com/jmcvetta/magicsquare).


## License

This is Free Software, released under the terms of the [GPL
v3](http://www.gnu.org/copyleft/gpl.html).
