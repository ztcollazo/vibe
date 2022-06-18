package main

import (
	"log"

	"github.com/ztcollazo/vibe/example"
)

func main() {
	log.Fatal(example.RunApp().Listen(":3000"))
}
