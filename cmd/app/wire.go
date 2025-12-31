//go:build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/acepanel/backup-template/internal/app"
	"github.com/acepanel/backup-template/internal/bootstrap"
	"github.com/acepanel/backup-template/internal/route"
	"github.com/acepanel/backup-template/internal/service"
)

// initCli init command line.
func initCli() (*app.Cli, error) {
	panic(wire.Build(bootstrap.ProviderSet, route.ProviderSet, service.ProviderSet, app.NewCli))
}
