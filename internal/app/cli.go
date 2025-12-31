package app

import (
	"context"
	"os"

	"github.com/acepanel/backup-ftp/internal/helper"
	"github.com/urfave/cli/v3"
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
