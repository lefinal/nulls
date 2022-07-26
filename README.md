# nulls

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
![Go](https://github.com/LeFinal/nulls/workflows/Go/badge.svg?branch=main)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/lefinal/nulls)
[![GoReportCard example](https://goreportcard.com/badge/github.com/lefinal/nulls)](https://goreportcard.com/report/github.com/lefinal/nulls)
[![codecov](https://codecov.io/gh/lefinal/nulls/branch/main/graph/badge.svg?token=ema8Z2HEk5)](https://codecov.io/gh/lefinal/nulls)
[![GitHub issues](https://img.shields.io/github/issues/lefinal/nulls)](https://github.com/lefinal/nulls/issues)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/lefinal/nulls)

This package uses the principle of NULL-types from the `sql`-package. However, it also adds JSON (un)marshalling
functionality while supporting NULL-values.
It's based on the idea of [gobuffalo/nulls](https://github.com/gobuffalo/nulls), but uses native (un)marshalling
functionality of the `json`-package.

# Installation

In order to use this package, run:

```shell
go get github.com/lefinal/nulls
```

# Predefined Datatypes

- `bool` (`nulls.Bool`)
- `[]byte` (`nulls.ByteSlice`)
- `float32` (`nulls.Float32`)
- `float64` (`nulls.Float64`)
- `int` (`nulls.Int`)
- `int16` (`nulls.Int16`)
- `int32` (`nulls.Int32`)
- `int64` (`nulls.Int64`)
- `json.RawMessage` (`nulls.JSONRawMessage`)
- `string` (`nulls.String`)
- `time.Time` (`nulls.Time`)

# Support for Generics

Any datatype implementing the required interface can be used as `Nullable` offering the same functionality as predefined
ones.
If no SQL-support is required, you can also use `JSONNullable`.

# Usage

All datatype feature a value and `Valid`-field. The latter one is `false` when a NULL-value is represented. Otherwise,
the actual value is found in the value-field. Constructors for creating non-NULL-values are available in the form of for
example `NewString(str)`. As the zero-value for the `Valid`-field is `false`, you do not need to create NULL-values
explicitly.