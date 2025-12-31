//go:build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/acepanel/backup-ftp/internal/app"
	"github.com/acepanel/backup-ftp/internal/bootstrap"
	"github.com/acepanel/backup-ftp/internal/route"
	"github.com/acepanel/backup-ftp/internal/service"
)

// initCli init command line.
func initCli() (*app.Cli, error) {
	panic(wire.Build(bootstrap.ProviderSet, route.ProviderSet, service.ProviderSet, app.NewCli))
}
