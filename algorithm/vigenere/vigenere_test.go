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

package vigenere

import (
	"strings"
	"testing"
)

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
	var sample = getSampleWithWhitespaces()
	var actual = Encrypt(sample.msg, sample.key)

	if actual != sample.enc {
		t.Fatalf(`Encryption of %q wrong: %q`, sample.msg, actual)
	}
}

func TestWordsWithWhitespacesDecrypt(t *testing.T) {
	var sample = getSampleWithWhitespaces()

	// Remember this program deletes information about whitespaces and uppercase
	// when the message was encrypted
	var expected = strings.ReplaceAll(strings.ToLower(sample.msg), " ", "")

	var actual = Decrypt(sample.enc, sample.key)

	if actual != expected {
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

func getSampleWithWhitespaces() sample {
	return sample{
		key: []byte{'C', 'I', 'P', 'H', 'E', 'R'},
		msg: "This crypto system is not secure",
		enc: "vpxzgiaxivwpubttmjpwizitwzt",
	}
}
