package affine

import (
	math "../../algorithm"
	"errors"
	"strings"
)

type E struct {
	Pair     Pair
	alphabet math.Alphabet
}

func (e *E) Set(pair Pair) error {
	if math.Gcd(int(pair.A), 26) != 1 {
		return errors.New("A and 26 must be relatively prime")
	}
	e.Pair = pair
	e.alphabet = math.GetAlphabet()
	return nil
}

func (e *E) Apply(x byte) byte {
	return (e.Pair.A*x + e.Pair.B) % 26
}

func (e *E) Inverse() D {
	return D{e.Pair, math.GetAlphabet()}
}

type D struct {
	Pair     Pair
	alphabet math.Alphabet
}

func (d *D) Apply(y byte) byte {
	var aInverse = 26 - d.Pair.A
	return aInverse * (y - d.Pair.B) % 26
}

type Pair struct {
	A byte
	B byte
}

func Encrypt(msg string, fn E) string {
	var enc = ""
	var input = strings.ToUpper(msg)

	for _, ch := range input {
		var x = byte(fn.alphabet.CanonicalPositionOf(byte(ch)))
		var y = fn.Apply(x)
		enc += string(fn.alphabet.Chars[y])
	}
	enc = strings.ToLower(enc)
	return enc
}
