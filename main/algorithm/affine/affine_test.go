package affine

import (
	"testing"
)

func TestOneWordEncrypt(t *testing.T) {
	const msg = "hot"
	const expected = "axg"
	var fn, err = getTestEncryptFunction()

	if err != nil {
		t.Fatal("Couldn't set valid Pair to E function")
	}
	var actual = Encrypt(msg, fn)

	if actual != expected {
		t.Fatalf(`Encryption of %q wrong: %q`, msg, actual)
	}
}

func TestOneWordDecrypt(t *testing.T) {
	const enc = "axg"
	const expected = "hot"
	var e, err = getTestEncryptFunction()

	if err != nil {
		t.Fatal("Couldn't set valid Pair to E function")
	}
	var fn = e.Inverse()
	var actual = Decrypt(enc, fn)

	if actual != expected {
		t.Fatalf(`Decryption of %q wrong: %q`, enc, actual)
	}
}

func TestWordsWithWhitespacesEncrypt(t *testing.T) {
	const msg = "Hot dog"
	const expected = "axgyxt"
	var fn, err = getTestEncryptFunction()

	if err != nil {
		t.Fatal("Couldn't set valid Pair to E function")
	}
	var actual = Encrypt(msg, fn)

	if actual != expected {
		t.Fatalf(`Encryption of %q wrong: %q`, msg, actual)
	}
}

func TestWordsWithWhitespacesDecrypt(t *testing.T) {
	const enc = "axgyxt"
	const expected = "hotdog"
	var e, err = getTestEncryptFunction()

	if err != nil {
		t.Fatal("Couldn't set valid Pair to E function")
	}
	var fn = e.Inverse()
	var actual = Decrypt(enc, fn)

	if actual != expected {
		t.Fatalf(`Decryption of %q wrong: %q`, enc, actual)
	}
}

func TestNotAcceptInvalidPair(t *testing.T) {
	var _, err = getInvalidTestEncryptFunction()

	if err == nil {
		t.Fatal("An invalid Pair to E function was accepted")
	}
}

func getTestEncryptFunction() (E, error) {
	var e = E{}
	var err = e.Set(Pair{7, 3})
	return e, err
}

func getInvalidTestEncryptFunction() (E, error) {
	var e = E{}
	var err = e.Set(Pair{13, 3})
	return e, err
}
