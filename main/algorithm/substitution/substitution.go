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

// Package substitution provides the data structures and implementation of the classical Substitution cipher algorithm.
//
// Author: Tobias Briones
package substitution

import (
	"github.com/tobiasbriones/ep-cryptosystems/main/algorithm"
	"strings"
)

type E struct {
	Image algorithm.Alphabet
}

func (e E) Apply(x byte) byte {
	return e.Image.Chars[x]
}

func (e E) Inverse() D {
	var d = D{}
	var image = e.Image
	var alphabet = algorithm.GetAlphabet()
	var inverseImage = algorithm.GetAlphabet()

	for i, ch := range image.Chars {
		var pos = image.CanonicalPositionOf(ch)
		inverseImage.Chars[pos] = alphabet.Chars[i]

		d.Image = inverseImage
	}
	return d
}

type D struct {
	Image algorithm.Alphabet
}

func (d D) Apply(y byte) byte {
	return d.Image.Chars[y]
}

func Encrypt(msg string, e E) string {
	var enc = ""
	var input = strings.ToUpper(msg)
	input = strings.ReplaceAll(input, " ", "")

	for _, ch := range input {
		var x = byte(e.Image.CanonicalPositionOf(byte(ch)))
		enc += string(e.Apply(x))
	}
	enc = strings.ToLower(enc)
	return enc
}

func Decrypt(enc string, d D) string {
	var msg = ""
	var input = strings.ToUpper(enc)

	for _, ch := range input {
		var y = byte(d.Image.CanonicalPositionOf(byte(ch)))
		msg += string(d.Apply(y))
	}
	msg = strings.ToLower(msg)
	return msg
}
