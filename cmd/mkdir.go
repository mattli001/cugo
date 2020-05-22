// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by an MIT-style license that
// may be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/mkdir"
	"github.com/jcmdln/flagger"
)

type mkdirCmd struct {
	name        string
	usage       string
	description string

	help bool
	mkdir.Options
}

func (u *mkdirCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "mkdir", "[-pv] [-m MODE] DIRECTORY ..."
	u.description = "Make directories"

	flags.UintVar(&u.Mode, 0755, "Set permissions to MODE value", "-m", "--mode")
	flags.BoolVar(&u.Parents, "Create missing parent directories", "-p", "--parents")
	flags.BoolVar(&u.Verbose, "Display each directory after it was created", "-v", "--verbose")
	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *mkdirCmd) Action(s []string, flags *flagger.Flags) error {
	var (
		data []string
		err  error
	)

	if data, err = flags.Parse(s); err != nil {
		err = fmt.Errorf("%s: %s", u.name, err)
		return err
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	if _, err = u.Mkdir(data); err != nil {
		return err
	}

	return nil
}

func init() {
	Command.Add("mkdir", &mkdirCmd{})
}
