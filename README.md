# Go Type-Agnostic Collection Partitioning

[![GoDoc](https://godoc.org/github.com/meirf/gopart?status.png)](https://godoc.org/github.com/meirf/gopart) [![Travis](https://travis-ci.org/meirf/gopart.svg?branch=master)](https://travis-ci.org/meirf/gopart)

Type-agnostic partitioning for anything that can be indexed in Go - slices, arrays,`string`s. Inspired by [Guava's Lists.partition](http://guava-libraries.googlecode.com/svn/tags/release09/javadoc/com/google/common/collect/Lists.html#partition(java.util.List, int)).

## Usage

```go
	for idxRange := range gopart.Partition(len(bigList), partitionSize) {
	    bulkOperation(bigList[idxRange.Low:idxRange.High])
	}
```

## Installation

    # install the library:
    go get github.com/meirf/gopart
    
    // use in your .go code:
    import (
        "github.com/meirf/gopart"
    )
