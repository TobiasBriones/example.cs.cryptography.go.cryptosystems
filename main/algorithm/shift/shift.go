package shift

import (
	"github.com/tobiasbriones/ep-cryptosystems/main/algorithm"
	"strings"
)

type position struct {
	key  int
	len  uint
	curr uint
}

// Encrypts the message with the given key using the Shift algorithm.
// The implemented shift algorithm given by (P, C, K, E, D):
//
// P := string [A-Za-z]
//
// C := string [A-Za-z]
//
// K := uint K in [0, 25]
//
// E and D are not bijective functions and hence one is not the inverse of the
// another. But they are bijective if and only if the input message is lowercase
// and not contains any whitespaces.
//
// The algorithm flow is as follows:
//
// Start
//
// 1. Transform msg into uppercase and remove whitespaces.
//
// 2. Shift right the amount of time indicated by the given key. It uses a
// circular alphabet.
//
// 3. Transform into lowercase
//
// End
func Encrypt(msg string, key uint) string {
	var input = getInput(msg)
	var shifted = shiftRight(input, key)
	return strings.ToLower(shifted)
}

// Decrypts the given encrypted message using the Shift algorithm.
// The decrypted message may not be like the exact original message since the
// whitespaces and capitalization information is lost when encrypting. The
// decrypted message is return in lowercase.
//
// See Encrypt
func Decrypt(enc string, key uint) string {
	var input = getInput(enc)
	var shifted = shiftLeft(input, key)
	return getOutput(shifted)
}

func getInput(msg string) string {
	var input = strings.ToUpper(msg)
	return strings.ReplaceAll(input, " ", "")
}

func getOutput(msg string) string {
	return strings.ToLower(msg)
}

func shiftRight(msg string, key uint) string {
	return shift(msg, int(key))
}

func shiftLeft(msg string, key uint) string {
	return shift(msg, -int(key))
}

func shift(str string, key int) string {
	var alphabet = algorithm.GetAlphabet()
	var chars = alphabet.Chars
	var shifted = ""
	var position = position{
		key:  key,
		len:  alphabet.Length(),
		curr: 0,
	}
	for _, ch := range str {
		position.curr = alphabet.CanonicalPositionOf(byte(ch))
		var newPos = getPosition(position)

		shifted += string(chars[newPos])
	}
	return shifted
}

func getPosition(pos position) uint {
	var length = int(pos.len)
	var newPos = int(pos.curr) + pos.key

	if newPos >= length {
		newPos %= length
	} else if newPos < 0 {
		newPos *= -1
		newPos %= length
		newPos = length - newPos
	}
	return uint(newPos)
}
