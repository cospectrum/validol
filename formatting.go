package validol

import (
	"fmt"
	"strings"
)

func fmtVarargs[T any](elems []T) string {
	toStr := func(el T) string {
		return fmt.Sprintf("%+v", el)
	}
	return strings.Join(mapF(elems, toStr), ", ")
}

func mapF[T any, U any](elems []T, f func(T) U) []U {
	out := make([]U, 0, len(elems))
	for _, el := range elems {
		out = append(out, f(el))
	}
	return out
}
