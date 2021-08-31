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

package main

import (
	"fmt"
	shift2 "github.com/tobiasbriones/ep-cryptosystems/algorithm/shift"
)

// Encrypts a message with the shift cipher.
// More use cases available on the test files.
func main() {
	const key = 3
	const msg = "cryptosystems"
	var enc = shift2.Encrypt(msg, key)

	fmt.Println("The message \"" + msg + "\" encrypted by the shift algorithm with key=3 is: " + enc)
}
