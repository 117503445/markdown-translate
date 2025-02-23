package examples

import (
	_ "embed"
)

//go:embed mempool.md
var Mempool string

//go:embed basic.md
var Basic string

var Examples = map[string]string{
	"mempool.md": Mempool,
	"basic.md":   Basic,
}
