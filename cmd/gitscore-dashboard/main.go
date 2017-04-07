package main

import (
	"log"
	"os"

	"github.com/oem/gitscore/ansi"
)

func main() {
	err := ansi.Draw()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
