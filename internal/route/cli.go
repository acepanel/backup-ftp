package route

import (
	"github.com/leonelquinteros/gotext"
	"github.com/urfave/cli/v3"

	"github.com/acepanel/backup-template/internal/service"
)

type Cli struct {
	t   *gotext.Locale
	cli *service.CliService
}

func NewCli(t *gotext.Locale, cli *service.CliService) *Cli {
	return &Cli{
		t:   t,
		cli: cli,
	}
}

func (route *Cli) Commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:   "check",
			Usage:  route.t.Get("Check the configuration"),
			Action: route.cli.Check,
		},
		{
			Name:   "mkdir",
			Usage:  route.t.Get("Make new directory"),
			Action: route.cli.Mkdir,
		},
		{
			Name:   "deldir",
			Usage:  route.t.Get("Delete directory"),
			Action: route.cli.Deldir,
		},
		{
			Name:   "put",
			Usage:  route.t.Get("Upload file"),
			Action: route.cli.Put,
		},
		{
			Name:   "del",
			Usage:  route.t.Get("Delete file"),
			Action: route.cli.Del,
		},
		{
			Name:   "get",
			Usage:  route.t.Get("Get file data"),
			Action: route.cli.Get,
		},
		{
			Name:   "ls",
			Usage:  route.t.Get("List directory or file"),
			Action: route.cli.Ls,
		},
	}
}
