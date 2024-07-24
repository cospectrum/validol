# validol
[![github]](https://github.com/cospectrum/validol)

[github]: https://img.shields.io/badge/github-cospectrum/validol-8da0cb?logo=github

Validation library for golang.

## Content
- [Install](#install)
- [Usage](#usage)
- [Validators](#validators)
- [Combinators](#combinators)

## Install
```sh
go get github.com/cospectrum/validol@latest
```
Requires Go version `1.22.0` or greater.

## Usage
```go
import (
	"errors"
	vd "github.com/cospectrum/validol"
)

type Sex string

func (s Sex) Validate() error {
	return vd.OneOf[Sex]("male", "female", "other")(s)
}

type Email string

func (e Email) Validate() error {
	return vd.All(
		vd.Len[string](vd.All(vd.Gt(5), vd.Lte(100))),
		vd.Email,
	)(string(e))
}

type Info struct {
	Email email
	Sex   Sex
	age   uint
}

func (info Info) Validate() error {
	return errors.Join(
		vd.Walk(info),
		vd.Gte(uint(18))(info.age),
	)
}

func main() {
	var info Info
	if err := info.Validate(); err != nil {
		panic(err)
	}
}
```

## Validators
Type `validol.Validator[T]` is equivalent to `func(T) error`.

| Name | Input | Description | 
| - | - | - |
| Walk | any | Recursively checks all "descendants" that have the `Validate() error` method |
| Email | string | Email string |
| UUID4 | string | Universally Unique Identifier UUID v4 |

## Combinators
Functions that create a `Validator[T]`.

| Name | Input | Output | Description |
| - | - | - | - |
| All | ...Validator[T] | Validator[T] | Logical and |
| Any | ...Validator[T] | Validator[T] | Logical or |
| Not | ...Validator[T] | Validator[T] | Logical not |
| OneOf | ...T | Validator[T] | Checks that the value is equal to one of the arguments | 
| Eq | T comparable | Validator[T] | == |
| Ne | T comparable | Validator[T] | != |
| Gt | T cmp.Ordered | Validator[T] | > |
| Gte | T cmp.Ordered | Validator[T] | >= |
| Lt | T cmp.Ordered | Validator[T] | < |
| Lte | T cmp.Ordered | Validator[T] | <= |
| Len | Validator[int] | Validator[T] | Checks that the len passes the specified Validator[int] |
| StartsWith | `prefix string` | Validator[string] | Checks if the string starts with the specified prefix |
| EndsWith | `suffix string` | Validator[string] | Checks whether the string ends with the specified suffix |
