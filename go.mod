module github.com/acepanel/backup-template

go 1.25.0

require (
	github.com/google/wire v0.7.0
	github.com/jlaffaye/ftp v0.2.0
	github.com/knadh/koanf/providers/env/v2 v2.0.0
	github.com/knadh/koanf/v2 v2.3.0
	github.com/leonelquinteros/gotext v1.7.2
	github.com/urfave/cli/v3 v3.6.1
)

require (
	github.com/go-viper/mapstructure/v2 v2.4.0 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/knadh/koanf/maps v0.1.2 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/tools v0.31.0 // indirect
)

tool (
	github.com/google/wire
	github.com/leonelquinteros/gotext/cli/xgotext
)
