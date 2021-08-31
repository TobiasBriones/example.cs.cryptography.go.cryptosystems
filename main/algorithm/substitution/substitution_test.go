/*
 * Copyright (c) 2021 Tobias Briones. All rights reserved.
 *
 * SPDX-License-Identifier: BSD-3-Clause
 *
 * This file is part of Example Project: Cryptosystems.
 *
 * This source code is licensed under the BSD-3-Clause License found in the
 * LICENSE file in the root directory of this source tree or at
 * https://opensource.org/licenses/BSD-3-Clause.
 */

package substitution

import (
	"github.com/tobiasbriones/ep-cryptosystems/main/algorithm"
	"testing"
)

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

func TestWordsWithWhitespaceEncrypt(t *testing.T) {
	const msg = "This cipher text cannot be decrypted"
	const expected = "mgzvyzlghcmhjmyxssfmnhahycdlmha"
	var fn = getTestEncryptFunction()
	var actual = Encrypt(msg, fn)

	if actual != expected {
		t.Fatalf(`Encryption of %q wrong: %q`, msg, actual)
	}
}

func TestWordsWithWhitespaceDecrypt(t *testing.T) {
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
