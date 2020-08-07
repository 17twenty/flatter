# go-flatten

Really simple library and demo app to flatten JSON into dot notation. [Forked][1] from [Nick Glynn][2]

[1]: https://github.com/17twenty/flatter
[2]: https://github.com/17twenty


[![GoDoc][3]][4]
[![GoCard][5]][6]
[![Build][7]][8]

[3]: https://godoc.org/github.com/cameronnewman/go-flatten?status.svg
[4]: https://godoc.org/github.com/cameronnewman/go-flatten
[5]: https://goreportcard.com/badge/github.com/cameronnewman/go-flatten
[6]: https://goreportcard.com/report/github.com/cameronnewman/go-flatten
[7]: https://travis-ci.org/cameronnewman/go-flatten.svg?branch=master
[8]: https://travis-ci.org/cameronnewman/go-flatten

## Getting Started

Simple stuff - use go get and it will also install the `flatten` demo app.

```
go get -u github.com/cameronnewman/go-flatten/...
```

```
$ flatten -s simple.json -w foo.txt
$ flatten -s simple.json
sample.someitem.thesearecool.1.eggs false
sample.someitem.thesearecool.1.neat tubular
sample.someitem.thesearecool.0.neat wow
sample.someitem.thesearecool.1.sausage true
```

You can use the library as you see fit. Not sure how useful it will be tho :D


## Issues
 * None

## License
MIT Licensed
