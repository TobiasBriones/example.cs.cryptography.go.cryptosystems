package shift

import "testing"

func TestOneWordMsgEncrypt(t *testing.T) {
	const key = 3
	const msg = "message"
	const expected = "phvvdjh"
	var actual = Encrypt(msg, key)

	if actual != expected {
		t.Fatalf(`Encryption of %q wrong: %q`, msg, actual)
	}
}

func TestTwoWordMsgEncrypt(t *testing.T) {
	const key = 5
	const msg = "Hello World"
	const expected = "mjqqtbtwqi"
	var actual = Encrypt(msg, key)

	if actual != expected {
		t.Fatalf(`Encryption of %q wrong: %q`, msg, actual)
	}
}
