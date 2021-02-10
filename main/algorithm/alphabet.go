package algorithm

const asciiInitialIndex = 65

type Alphabet struct {
	Chars []byte
}

func (al Alphabet) Length() uint {
	return uint(len(al.Chars))
}

func (al Alphabet) CanonicalPositionOf(ch byte) uint {
	return uint(ch) - asciiInitialIndex
}

func GetAlphabet() Alphabet {
	// byte array represents ASCII chars
	// A = 65, ... , Z = 90
	return Alphabet{
		Chars: []byte{
			'A', 'B', 'C', 'D', 'E', 'F', 'G',
			'H', 'I', 'J', 'K', 'L', 'M', 'N',
			'O', 'P', 'Q', 'R', 'S', 'T', 'U',
			'V', 'W', 'X', 'Y', 'Z',
		},
	}
}
