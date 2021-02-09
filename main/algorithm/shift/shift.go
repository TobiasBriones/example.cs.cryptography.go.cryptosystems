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
