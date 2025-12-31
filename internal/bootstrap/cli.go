package bootstrap

import (
	"github.com/leonelquinteros/gotext"
	"github.com/urfave/cli/v3"

	"github.com/acepanel/backup-ftp/internal/route"
)

func NewCli(t *gotext.Locale, cmd *route.Cli) *cli.Command {
	return &cli.Command{
		Name:     "backup-plugin-ftp",
		Usage:    t.Get("FTP backup plugin for AcePanel"),
		Commands: cmd.Commands(),
	}
}
