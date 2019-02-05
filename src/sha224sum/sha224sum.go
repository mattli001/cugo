// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// compute and check SHA224 message digest
//
// SYNOPSIS
//
//     sha224sum [-bct] [FILE ...]
//
// DESCRIPTION
//
// sha224sum prints or checks SHA224 checksums. If no file is given in the
// operands, read standard input.
//
// SEE ALSO
//
//     tbd
//
// REFERENCES
//
//     https://linux.die.net/man/1/sha224sum
package sha224sum

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	// Binary is a boolean that, when true, enables reading in binary
	// mode.
	Binary bool
	// Check is a boolean that, when true, reads sums from FILEs and
	// checks them.
	Check bool
	// Text is a bool that, when true, enables reading in binary mode,
	// though this is the default behavior.
	Text bool

	data []byte
)

// Sha224sum ...
func Sha224sum(args []string) {
	for _, file := range args {
		if contents, err := ioutil.ReadFile(file); err != nil {
			fmt.Printf("cugo: %s\n", err)
			os.Exit(1)
		} else {
			data = []byte(contents)
			fmt.Printf("%x  %s\n", sha256.Sum224(data), file)
		}
	}

	os.Exit(0)
}
