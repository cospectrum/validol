package main

import (
	"fmt"

	vd "github.com/cospectrum/validol"
)

type userStatus string

const (
	offlie userStatus = "offline"
	online userStatus = "online"
	banned userStatus = "banned"
)

func (a userStatus) Validate() error {
	return vd.OneOf(offlie, online, banned)(a)
}

type uuid string

func (u uuid) Validate() error {
	return vd.UUID4(string(u))
}

func main() {
	Map := map[uuid]userStatus{
		"6926c346-c9d9-431c-854f-a6a3574511e4": offlie,
		"28ef9ebb-1653-4a34-80e0-fc6e3ae90fa2": banned,
		"127c3be4-e475-4219-8bcc-12fba6fd6c53": online,
	}
	if err := vd.Validate(Map); err != nil {
		panic(err)
	}
	fmt.Println("Map is valid")
}
