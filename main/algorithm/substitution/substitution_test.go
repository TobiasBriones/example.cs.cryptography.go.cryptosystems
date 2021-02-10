package substitution

import "testing"
import "../../algorithm"

func TestOneWordEncrypt(t *testing.T) {
	const msg = "thisciphertextcannotbedecrypted"
	const expected = "mgzvyzlghcmhjmyxssfmnhahycdlmha"
	var fn = getTestEncryptFunction()
	var actual = Encrypt(msg, fn)

	if actual != expected {
		t.Fatalf(`Encryption of %q wrong: %q`, msg, actual)
	}
}

func TestOneWordDecrypt(t *testing.T) {
	const enc = "mgzvyzlghcmhjmyxssfmnhahycdlmha"
	const expected = "thisciphertextcannotbedecrypted"
	var fn = getTestEncryptFunction().Inverse()
	var actual = Decrypt(enc, fn)

	if actual != expected {
		t.Fatalf(`Decryption of %q wrong: %q`, enc, actual)
	}
}

func getTestEncryptFunction() E {
	return E{
		Image: algorithm.Alphabet{
			Chars: []byte{
				'X', 'N', 'Y', 'A', 'H', 'P', 'O', 'G',
				'Z', 'Q', 'W', 'B', 'T', 'S', 'F', 'L',
				'R', 'C', 'V', 'M', 'U', 'E', 'K', 'J',
				'D', 'I',
			},
		},
	}
}
