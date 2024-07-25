package main

import (
	"fmt"

	vd "github.com/cospectrum/validol"
)

type email string

func (e email) Validate() error {
	return vd.Email(string(e))
}

func main() {
	emails := []email{"test@gmail.com", "test2@gmail.com"}

	if err := vd.Walk(emails); err != nil {
		panic(err)
	}
	fmt.Println("emails are valid")
}
