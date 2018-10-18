## DED (Disposable Email Detector)

[![Build Status](https://www.travis-ci.org/namreg/ded.svg?branch=master)](https://www.travis-ci.org/namreg/ded)
[![Go Report Card](https://goreportcard.com/badge/github.com/namreg/ded)](https://goreportcard.com/report/github.com/namreg/ded)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/namreg/ded/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/namreg/ded?status.svg)](https://godoc.org/github.com/namreg/ded)

Go package that detects disposable emails based on the [list](https://github.com/wesbos/burner-email-providers).

_Notice: DED does not check that email address is correct._

### Installation
```bash
go get github.com/namreg/ded
```

### How to use
```go
package main

import (
	"fmt"

	"github.com/namreg/ded"
)

func main() {
	disposable, _ := ded.IsDisposableEmail("temp@mail.wtf")
	fmt.Println(disposable) // true

	disposable, _ = ded.IsDisposableEmail("temp@google.com")
	fmt.Println(disposable) // false

	disposable, _ := ded.IsDisposableDomain("mail.wtf")
	fmt.Println(disposable) // true

	disposable, _ = ded.IsDisposableDomain("google.com")
	fmt.Println(disposable) // false
}
```
