package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const latinAlphabet = "abcdefghijklmnopqrstuvwxyz"

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %v <input>\n"+
		"    Transform a string into a numeric representation.\n",
		os.Args[0])
	os.Exit(1)
}

// alphabetOrder return a representation of a string using letter's position in the alphabet.
// Should be extended to use another base.
func alphabetOrder(str string) (string, error) {
	var num string
	lower := strings.ToLower(str)
	for _, c := range lower {
		i := strings.IndexRune(latinAlphabet, c)
		if i == -1 {
			return "", errors.New("invalid input: latin aplhabet only")
		}
		i++
		num += strconv.Itoa(i)
	}
	return num, nil
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	str := os.Args[1]

	num, err := alphabetOrder(str)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}
	fmt.Println(num)
}
