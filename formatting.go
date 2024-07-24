package validol

import (
	"fmt"
	"strings"
)

func fmtVarargs[T any](elems []T) string {
	return strings.Join(stringSlice(elems), ", ")
}

func stringSlice[T any](elems []T) []string {
	out := make([]string, 0, len(elems))
	for _, el := range elems {
		out = append(out, fmt.Sprintf("%+v", el))
	}
	return out
}
