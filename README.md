# nulls

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
- `string` (`nulls.String`)
- `time.Time` (`nulls.Time`)

# Support for Generics

Any datatype implementing the required interface can be used as `Nullable` offering the same functionality as predefined
ones.

# Usage

All datatype feature a value and `Valid`-field. The latter one is `false` when a NULL-value is represented. Otherwise,
the actual value is found in the value-field. Constructors for creating non-NULL-values are available in the form of for
example `NewString(str)`. As the zero-value for the `Valid`-field is `false`, you do not need to create NULL-values
explicitly.