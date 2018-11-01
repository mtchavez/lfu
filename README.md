# LFU Cache

[![Latest Version](http://img.shields.io/github/release/mtchavez/lfu.svg?style=flat-square)](https://github.com/mtchavez/lfu/releases)
[![Build Status](https://travis-ci.org/mtchavez/lfu.svg)](https://travis-ci.org/mtchavez/lfu)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/mtchavez/lfu)
[![Go Report Card](https://goreportcard.com/badge/github.com/mtchavez/lfu)](https://goreportcard.com/report/github.com/mtchavez/lfu)
[![Maintainability](https://api.codeclimate.com/v1/badges/97808771ae80f01c8c65/maintainability)](https://codeclimate.com/github/mtchavez/lfu/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/97808771ae80f01c8c65/test_coverage)](https://codeclimate.com/github/mtchavez/lfu/test_coverage)

LFU cache in Go offering O(1) Get and Insert [outlined here](http://dhruvbird.com/lfu.pdf).

## Install

`go get -u github.com/mtchavez/lfu`

## Usage

```go
package main

import (
    "fmt"
    "github.com/mtchavez/lfu"
)

func main() {
    cache := lfu.NewLFU()

    // Insert
    cache.Insert(42, []byte("user:42:user@example.com"))

    // Insert existing key
    success, err := cache.Insert(42, "Nope")
    if !success {
        fmt.Println("Error inserting:", err)
    }

    var data interface{}
    var e error
    // Get
    data, e = cache.Get(42)
    fmt.Println("Data for 42 is", data)

    // Get not found
    data, e = cache.Get(987654321)
    if e != nil {
        fmt.Println("Error on get:", e)
    }
}

```

## Tests

`go test --cover`

## Benhcmarks

Get and insert methods are benchmarked. Results from OS X with
a 2.3 GHz Intel Core i7 on `go version go1.8.3 darwin/amd64`

```
# Updated: 2017-08-15

BenchmarkInsert-8                1000000              1860 ns/op
BenchmarkParallelInsert-8        1000000              1861 ns/op
BenchmarkGet_EmptyCache-8        5000000               362 ns/op
BenchmarkGet_AllMisses-8         3000000               732 ns/op
BenchmarkGet_AllHits-8           1000000              1417 ns/op
BenchmarkParallelGet-8           2000000              1405 ns/op
```

## TODO

* Some kind of eviction
