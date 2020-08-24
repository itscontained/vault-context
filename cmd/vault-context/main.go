package main

import (
	"github.com/itscontained/vault-context/internal/commands"
)

var (
	version = "0.0.0"
	date    = "1970-01-01T00:00:00Z"
	commit  = ""
)

func main() {
	commands.Version = &version
	commands.Date = &date
	commands.Commit = &commit
	commands.Execute()
}
