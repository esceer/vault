package main

import (
	"sync"

	"github.com/esceer/vault/cli"
	"github.com/esceer/vault/storage"
)

var (
	once   sync.Once
	engine *cli.Engine
)

func GetEngine() *cli.Engine {
	once.Do(func() {
		store := storage.New()
		engine = cli.NewEngine(store)
	})
	return engine
}

func main() {
	cli.DisplayMenu(GetEngine())
}
