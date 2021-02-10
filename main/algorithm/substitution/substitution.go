package substitution

import (
	"../../algorithm"
)

type E struct {
	Image algorithm.Alphabet
}

func (e E) Apply(x byte) byte {
	return e.Image.Chars[x]
}

func (e E) Inverse() D {
	var d = D{}
	var image = e.Image
	var alphabet = algorithm.GetAlphabet()
	var inverseImage = algorithm.GetAlphabet()

	for i, ch := range image.Chars {
		var pos = image.CanonicalPositionOf(ch)
		inverseImage.Chars[pos] = alphabet.Chars[i]

		d.Image = inverseImage
	}
	return d
}

type D struct {
	Image algorithm.Alphabet
}

func (d D) Apply(y byte) byte {
	return d.Image.Chars[y]
}
