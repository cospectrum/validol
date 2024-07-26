# validol
[![github]](https://github.com/cospectrum/validol)
[![goref]](https://pkg.go.dev/github.com/cospectrum/validol)

[github]: https://img.shields.io/badge/github-cospectrum/validol-8da0cb?logo=github
[goref]: https://pkg.go.dev/badge/github.com/cospectrum/validol

Validation library for golang.

## Content
- [Install](#install)
- [Usage](#usage)
- [Validators](#validators)
- [Combinators](#combinators)

## Install
```sh
go get -u github.com/cospectrum/validol
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
		vd.Len[string](vd.Gt(5)),
		vd.Len[string](vd.Lte(100)),
		vd.Email,
	)(string(e))
}

type User struct {
	Email Email
	Sex   Sex
	age   int
}

func (u User) Validate() error {
	return errors.Join(
		vd.Gte(18)(u.age),
		vd.Walk(u), // to continue the `Walk` using the descendants' `Validate` methods
	)
}

func main() {
	users := []User{
		{Email: "first_user@mail.com", age: 22},
		{Email: "second_user@mail.com", Sex: "male"},
	}
	if err := vd.Validate(users); err != nil {
		panic(err)
	}
}
```

## Validators
Type `Validator[T]` is equivalent to `func(T) error`. \
`Validatable` interface requires `Validate() error` method.

| Name | Input | Description |
| - | - | - |
| Validate | T | If `T` is `Validatable`, then it will call `Validate` method, otherwise will call `Walk` |
| Walk | T | Recursively calls `Validate` method for `descendants` of `T`. The descendants of the `Validatable` descendant will not be checked automatically, instead the type must continue `Walk` manually (inside its own `Validate`). The `descendants` are public struct fields, embedded types, slice/array elements, map keys/values. |
| Required | T | Checks that the value is different from `default` |
| Empty | T | Checks that the value is initialized as `default` |
| NotNil | T | Checks that the value is different from `nil` |
| Nil | T | Checks that the value is `nil` |
| Email | string | Email string |
| UUID4 | string | Universally Unique Identifier UUID v4 |

## Combinators
Functions that create a `Validator[T]`.

| Name | Input | Output | Description |
| - | - | - | - |
| All | ...Validator[T] | Validator[T] | Checks whether `all` validations have been completed successfully |
| Any | ...Validator[T] | Validator[T] | Checks that `at least one` validation has been completed successfully |
| Not | Validator[T] | Validator[T] | Logical not |
| OneOf | ...T | Validator[T] | Checks that the value is equal to one of the arguments | 
| Eq | T comparable | Validator[T] | == |
| Ne | T comparable | Validator[T] | != |
| Gt | T cmp.Ordered | Validator[T] | > |
| Gte | T cmp.Ordered | Validator[T] | >= |
| Lt | T cmp.Ordered | Validator[T] | < |
| Lte | T cmp.Ordered | Validator[T] | <= |
| Len | Validator[int] | Validator[T] | Checks whether the `len` of the object passes the specified `Validator[int]` |
| StartsWith | string | Validator[string] | Checks if the string starts with the specified prefix |
| EndsWith | string | Validator[string] | Checks whether the string ends with the specified suffix |
| Contains | string | Validator[string] | Checks whether the specified substr is within string |
| ContainsRune | rune | Validator[string] | Checks whether the specified Unicode code point is within string |
