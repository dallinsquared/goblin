package main

import (
	"io"
	"log"
	"os"
)

func cherror(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) > 1 {
		for _, val := range os.Args[1:] {
			if val == "-" {
				_, err := io.Copy(os.Stdout, os.Stdin)
				cherror(err)
			} else {
				f, err := os.Open(val)
				cherror(err)
				_, err = io.Copy(os.Stdout, f)
				cherror(err)
			}
		}
	} else {
		_, err := io.Copy(os.Stdout, os.Stdin)
		cherror(err)

	}
}
