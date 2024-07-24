package validol

import (
	"fmt"
)

var _ Validator[string] = Email

func Email(s string) error {
	ok := emailRegex().MatchString(s)
	if !ok {
		return fmt.Errorf("validol.Email(%s) failed", s)
	}
	return nil
}
