package main

import (
	"github.com/claerhead/go_blockchain/cli"
	"github.com/claerhead/go_blockchain/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
