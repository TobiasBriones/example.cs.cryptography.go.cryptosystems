package vigenere

import (
	"github.com/TobiasBriones/example.cs.cryptography.go.cryptosystems/main/algorithm"
	"strings"
)

func Encrypt(msg string, key []byte) string {
	var enc = ""
	var length = len(key)
	var input = strings.ToUpper(msg)
	var alphabet = algorithm.GetAlphabet()
	var m = byte(alphabet.Length())
	input = strings.ReplaceAll(input, " ", "")

	for i, ch := range input {
		var code = alphabet.CanonicalPositionOf(byte(ch))
		var groupIndex = i % length
		var keyCh = key[groupIndex]
		var add = alphabet.CanonicalPositionOf(keyCh)
		var summation = (byte(code) + byte(add)) % m
		enc += string(alphabet.Chars[summation])
	}
	enc = strings.ToLower(enc)
	return enc
}

func Decrypt(enc string, key []byte) string {
	var msg = ""
	var input = strings.ToUpper(enc)
	var alphabet = algorithm.GetAlphabet()
	var length = len(key)

	for i, ch := range input {
		var groupIndex = i % length
		var keyCh = key[groupIndex]
		var subtract = int(alphabet.CanonicalPositionOf(keyCh))
		var summation = int(alphabet.CanonicalPositionOf(byte(ch)))
		var code = norm(summation - subtract)
		msg += string(alphabet.Chars[code])
	}
	msg = strings.ToLower(msg)
	return msg
}

// Used to fix negative results, I took this function from the shift algorithm
// since I solved that problem there before
func norm(newPos int) int {
	if newPos < 0 {
		newPos *= -1
		newPos %= 26
		newPos = 26 - newPos
	}
	return newPos
}
