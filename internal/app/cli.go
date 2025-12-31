package app

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"

	"github.com/acepanel/backup-template/internal/helper"
)

type Cli struct {
	cmd *cli.Command
}

func NewCli(cmd *cli.Command) *Cli {
	return &Cli{
		cmd: cmd,
	}
}

func (r *Cli) Run() {
	if err := r.cmd.Run(context.TODO(), os.Args); err != nil {
		helper.Error(err.Error())
	}
}
