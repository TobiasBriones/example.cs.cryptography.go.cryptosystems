package main

import (
	"fmt"
	"github.com/TobiasBriones/example.cs.cryptography.go.cryptosystems/main/algorithm"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	var a = []byte{
		21, 15, 23, 25, 6, 8, 0, 23, 8, 21, 22, 15, 20, 1, 19, 19, 12, 9, 15, 22, 8, 25, 8, 19, 22, 25, 19,
	}
	var al = algorithm.GetAlphabet()
	var enc = ""

	for _, v := range a {
		enc += string(al.Chars[v])
	}
	enc = strings.ToLower(enc)
	fmt.Println(enc)
}
