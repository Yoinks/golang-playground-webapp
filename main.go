package main

import (
	"log"

	"github.com/pokeman-service/cmd"
)

func main() {
}

func init() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
