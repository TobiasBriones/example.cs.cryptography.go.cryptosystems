package vigenere

import "strings"
import "../../algorithm"

func Encrypt(msg string, key []byte) string {
	var enc = ""
	var length = len(key)
	var input = strings.ToUpper(msg)
	var alphabet = algorithm.GetAlphabet()
	var m = byte(alphabet.Length())

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
