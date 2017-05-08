# Flatter

Really simple library and demo app to flatten JSON into dot notation.

## Getting Started

Simple stuff - use go get and it will also install the `flatten` demo app.

```
go get -u github.com/17twenty/flatter/...
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
