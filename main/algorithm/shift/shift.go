package shift

import "strings"

const asciiInitialIndex = 65

type alphabet struct {
	chars []byte
}

type position struct {
	key  int
	len  uint
	curr uint
}

func Encrypt(msg string, key uint) string {
	var input = getInput(msg)
	var shifted = shiftRight(input, key)
	return strings.ToLower(shifted)
}

func Decrypt(enc string, key uint) string {
	var input = getInput(enc)
	var shifted = shiftLeft(input, key)
	return getOutput(shifted)
}

func (al alphabet) length() uint {
	return uint(len(al.chars))
}

func (al alphabet) positionOf(ch byte) uint {
	return uint(ch) - asciiInitialIndex
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
	var alphabet = getAlphabet()
	var chars = alphabet.chars
	var shifted = ""
	var position = position{
		key:  key,
		len:  alphabet.length(),
		curr: 0,
	}
	for _, ch := range str {
		position.curr = alphabet.positionOf(byte(ch))
		var newPos = getPosition(position)

		shifted += string(chars[newPos])
	}
	return shifted
}

func getAlphabet() alphabet {
	// byte array represents ASCII chars
	// A = 65, ... , Z = 90
	return alphabet{
		chars: []byte{
			'A', 'B', 'C', 'D', 'E', 'F', 'G',
			'H', 'I', 'J', 'K', 'L', 'M', 'N',
			'O', 'P', 'Q', 'R', 'S', 'T', 'U',
			'V', 'W', 'X', 'Y', 'Z',
		},
	}
}

func getPosition(pos position) uint {
	var length = int(pos.len)
	var newPos = int(pos.curr) + pos.key

	if newPos >= length {
		newPos -= length
	} else if newPos < 0 {
		newPos = length + newPos
	}
	return uint(newPos)
}
