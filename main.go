package main

import (
	"embed"
	"log"

	cmd "github.com/rosikui/code-quote/cmd/code-quote"
)

//go:embed quotes/*.json
var embeddedQuotes embed.FS

func main() {
	if err := cmd.Execute(embeddedQuotes); err != nil {
		log.Fatal(err)
	}
}
