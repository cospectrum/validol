package validol

import (
	"fmt"
)

var _ Validator[string] = Email

func Email(s string) error {
	ok := emailRegex().MatchString(s)
	if !ok {
		return failed(fmt.Sprintf("validol.Email(%q)", s))
	}
	return nil
}

var _ Validator[string] = UUID4

func UUID4(s string) error {
	ok := uuid4Regex().MatchString(s)
	if !ok {
		return failed(fmt.Sprintf("validol.UUID4(%q)", s))
	}
	return nil
}
