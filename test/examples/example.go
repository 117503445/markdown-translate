package examples

import (
	_ "embed"
)

//go:embed mempool.md
var Mempool string

//go:embed all.md
var All string

var Examples = map[string]string{
	"mempool.md": Mempool,
	"all.md":     All,
}
