// args: -Econtainedctx
package testdata

import "context"

type ok struct {
	i int
	s string
}

type ng struct {
	ctx context.Context // ERROR "found a struct that contains a context.Context field"
}

type empty struct{}
