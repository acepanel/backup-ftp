package bootstrap

import (
	"github.com/knadh/koanf/v2"
	"github.com/leonelquinteros/gotext"

	"github.com/acepanel/backup-ftp/pkg/embed"
)

func NewT(conf *koanf.Koanf) (*gotext.Locale, error) {
	locale := conf.String("locale")
	l := gotext.NewLocaleFSWithPath(locale, embed.LocalesFS, "locales")
	l.AddDomain("app")

	return l, nil
}
