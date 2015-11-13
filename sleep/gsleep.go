package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func usage() {
	fmt.Println("Usage: ", os.Args[0], "seconds")
}
func main() {
	if len(os.Args) != 2 {
		usage()
		return
	}
	dur, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
	}
	time.Sleep(time.Duration(dur) * time.Second)
}
