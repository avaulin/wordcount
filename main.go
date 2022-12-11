package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}
	if src == "" {
		fmt.Println(0)
		return
	}
	words := strings.Split(src, " ")
	fmt.Println(len(words))
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string, err error) {
	if len(os.Args) <= 1 {
		return "", err
	}
	return os.Args[1], nil
}

func fail(err error) {
	fmt.Println(err)
	os.Exit(1)
}
