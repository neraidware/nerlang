package main

import (
	"fmt"
	"log"
	"os"

	"neraid.live/mbl/tokenizer"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("usage: %v script.mbl\n", os.Args[0])
		os.Exit(1)
	}

	file := os.Args[1]

	script_contents, err := os.ReadFile(file)

	if err != nil {
		log.Fatalf("%v: failed to open file", os.Args[0])
		os.Exit(1)
	}

	log.Printf("beginning parsing")

	script_tokenizer := tokenizer.Tokenizer{}

	script_tokenizer.Tokenize(string(script_contents))
}
