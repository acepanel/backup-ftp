package bootstrap

import (
	"github.com/leonelquinteros/gotext"
	"github.com/urfave/cli/v3"

	"github.com/acepanel/backup-template/internal/route"
)

func NewCli(t *gotext.Locale, cmd *route.Cli) *cli.Command {
	return &cli.Command{
		Name:     "backup-plugin-template",
		Usage:    t.Get("Template backup plugin for AcePanel"),
		Commands: cmd.Commands(),
	}
}
