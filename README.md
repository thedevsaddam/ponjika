Package ponjika
================
[![Build Status](https://travis-ci.org/thedevsaddam/ponjika.svg?branch=master)](https://travis-ci.org/thedevsaddam/ponjika)
[![Project status](https://img.shields.io/badge/version-1.0-green.svg)](https://github.com/thedevsaddam/ponjika/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/thedevsaddam/ponjika)](https://goreportcard.com/report/github.com/thedevsaddam/ponjika)
[![Coverage Status](https://coveralls.io/repos/github/thedevsaddam/ponjika/badge.svg?branch=master)](https://coveralls.io/github/thedevsaddam/ponjika?branch=master)
[![GoDoc](https://godoc.org/github.com/thedevsaddam/ponjika?status.svg)](https://godoc.org/github.com/thedevsaddam/ponjika)
[![License](https://img.shields.io/dub/l/vibe-d.svg)](https://github.com/thedevsaddam/ponjika/blob/dev/LICENSE.md)

Tiny bengali ponjika based on Gregorian date

### Installation

Install the package using
```go
$ go get github.com/thedevsaddam/ponjika
```

### Usage

To use the package import it in your `*.go` code
```go
import "github.com/thedevsaddam/ponjika"
```
### Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/thedevsaddam/ponjika"
)

func main() {
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, "2018-04-14 14:18:00")
	p := ponjika.New(t)
	fmt.Println(p)
	fmt.Println(p.Phonetic())
}
// output:
// ১ বৈশাখ ১৪২৫ রোজ শনিবার
// 1 Boisakh 1425 Roj Shonibar
```

### Credit
This package is directly ported from [Nuhil Mehdy's](https://github.com/nuhil) [bangla-calendar](https://github.com/nuhil/bangla-calendar)

Special thanks to [Ahmed shamim](https://github.com/me-shaon)

### See all [contributors](https://github.com/thedevsaddam/ponjika/graphs/contributors)

### Read [API doc](https://godoc.org/github.com/thedevsaddam/ponjika) to know about ***Available options and Methods***

### **License**
The **ponjika** is an open-source software licensed under the [MIT License](LICENSE.md).
