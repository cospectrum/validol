package validol

import (
	"fmt"
	"strings"
)

//nolint:errname // struct is private
type failedExpr struct {
	expr string
}

var _ error = &failedExpr{}

func (e failedExpr) Error() string {
	return e.expr + " failed"
}

func failed(expr string) error {
	return failedExpr{expr: expr}
}

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
