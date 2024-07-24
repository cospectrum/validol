# validol
Validation library for golang

## Install
```sh
go get github.com/cospectrum/validol
```

## Usage

```go
import "github.com/cospectrum/validol"

type Sex string

func (s Sex) Validate() error {
    validate := validol.Any(
        validol.OneOf("male", "female", "other"),
        validol.OneOf("", nil)
    )
    return validate(string(s))
}

func main() {
    var in Sex
    if err := in.Validate(); err != nil {
        panic(err)
    }
}
```
