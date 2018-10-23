// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package sleep

import (
	"strings"
	"time"

	er "github.com/jcmdln/cugo/lib/error"
)

func Sleep(args string) {
	a := strings.Split(args, " ")

	var (
		s   []string
		t   time.Duration
		err error
	)

	for _, i := range a {
		s = strings.Split(i, "")
		if len(s) < 2 {
			i += "s"
		}

		t, err = time.ParseDuration(i)
		if !er.Error("cugo", err) {
			time.Sleep(t)
		}
	}
}
