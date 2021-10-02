package main

import (
	"log"
	"os"
)

func main() {
	if err := MakeGolden(os.Stdout, os.Stderr); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
