package vigenere

import "testing"

type sample struct {
	key []byte
	msg string
	enc string
}

func (s *sample) length() int {
	return len(s.key)
}

func TestOneWordEncrypt(t *testing.T) {
	var sample = getSample()
	var actual = Encrypt(sample.msg, sample.key)

	if actual != sample.enc {
		t.Fatalf(`Encryption of %q wrong: %q`, sample.msg, actual)
	}
}

func TestOneWordDecrypt(t *testing.T) {
	var sample = getSample()
	var actual = Decrypt(sample.enc, sample.key)

	if actual != sample.msg {
		t.Fatalf(`Decryption of %q wrong: %q`, sample.enc, actual)
	}
}

func TestWordsWithWhitespacesEncrypt(t *testing.T) {
	var sample = getSampleWithWhistespaces()
	var actual = Encrypt(sample.msg, sample.key)

	if actual != sample.enc {
		t.Fatalf(`Encryption of %q wrong: %q`, sample.msg, actual)
	}
}

func TestWordsWithWhitespacesDecrypt(t *testing.T) {
	var sample = getSampleWithWhistespaces()
	var actual = Decrypt(sample.enc, sample.key)

	if actual != sample.msg {
		t.Fatalf(`Decryption of %q wrong: %q`, sample.enc, actual)
	}
}

func getSample() sample {
	return sample{
		key: []byte{'C', 'I', 'P', 'H', 'E', 'R'},
		msg: "thiscryptosystemisnotsecure",
		enc: "vpxzgiaxivwpubttmjpwizitwzt",
	}
}

func getSampleWithWhistespaces() sample {
	return sample{
		key: []byte{'C', 'I', 'P', 'H', 'E', 'R'},
		msg: "This cryptosystem is not secure",
		enc: "vpxzgiaxivwpubttmjpwizitwzt",
	}
}
