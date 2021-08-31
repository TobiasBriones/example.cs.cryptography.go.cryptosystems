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

func TestOneWordMsgDecrypt(t *testing.T) {
	const key = 3
	const enc = "phvvdjh"
	const expected = "message"
	var actual = Decrypt(enc, key)

	if actual != expected {
		t.Fatalf(`Decryption of %q wrong: %q`, enc, actual)
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

func TestTwoWordMsgDecrypt(t *testing.T) {
	const key = 5
	const enc = "mjqqtbtwqi"
	const expected = "helloworld"
	var actual = Decrypt(enc, key)

	if actual != expected {
		t.Fatalf(`Decryption of %q wrong: %q`, enc, actual)
	}
}
