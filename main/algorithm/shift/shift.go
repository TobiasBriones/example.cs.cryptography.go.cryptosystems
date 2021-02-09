package shift

const asciiInitialIndex = 65

type alphabet struct {
	chars []byte
}

type position struct {
	key  int
	len  uint
	curr uint
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
