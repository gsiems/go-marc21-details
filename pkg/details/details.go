// Copyright 2017-2018 Gregory Siems. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

// Package marc21-details extracts detailed information from MARC21
// records.

package details

// CodeValue contains a code and it's corresponding descriptive label
// for a controlfield entry.
type CodeValue struct {
	Code   string
	Label  string
	Offset int
	Width  int
}

// pluckByte extracts a single-byte from a string and returns
// the string result.
func pluckByte(b string, i int) (s string) {

	if len(b) > i {
		s = string(b[i])
	}

	return s
}

// pluckBytes extracts one or more bytes from a string and returns
// the string result.
func pluckBytes(b string, i, w int) (s string) {

	if w == 0 {
		return pluckByte(b, i)
	}
	if len(b) >= i+w {
		s = b[i : i+w]
	}
	return s
}

func codeLookup(codeList map[string]string, b string, i, w int) (code, label string) {

	code = pluckBytes(b, i, w)

	if code != "" {
		label = codeList[code]
	}

	return code, label
}
