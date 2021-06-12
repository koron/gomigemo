# Go/Migemo module

[![PkgGoDev](https://pkg.go.dev/badge/github.com/koron/gomigemo)](https://pkg.go.dev/github.com/koron/gomigemo)
[![GoDoc](https://godoc.org/github.com/koron/gomigemo?status.svg)](https://godoc.org/github.com/koron/gomigemo)
[![Actions/Go](https://github.com/koron/gomigemo/workflows/Go/badge.svg)](https://github.com/koron/gomigemo/actions?query=workflow%3AGo)
[![Go Report Card](https://goreportcard.com/badge/github.com/koron/gomigemo)](https://goreportcard.com/report/github.com/koron/gomigemo)

## Usage

To load dictionary files from file system:

```go
// Import migemo package.
import "github.com/koron/gomigemo/migemo"

// Load dictionary files.
dict, err := migemo.LoadDefault()

// Compile to get *regexp.Regexp.
re, err := migemo.Compile(dict, "aiueo")
```

To embedded dictionary to the executable file:

```go
// Import migemo and embedict package.
import (
    "github.com/koron/gomigemo/embedict"
    "github.com/koron/gomigemo/migemo"
)

// Load embedded dictionary.
dict, err := embedict.Load()

// Compile to get *regexp.Regexp.
re, err := migemo.Compile(dict, "aiueo")
```

## LICENSE

Distributed under MIT License,
except for `_dict/SKK-JISYO.utf-8.L` and `embedict/bindata.go` which is GPL.

See LICENSE.
