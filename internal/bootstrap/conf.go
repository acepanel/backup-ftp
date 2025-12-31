package bootstrap

import (
	"strings"

	"github.com/knadh/koanf/providers/env/v2"
	"github.com/knadh/koanf/v2"
)

func NewConf() (*koanf.Koanf, error) {
	k := koanf.New(".")

	if err := k.Load(env.Provider(".", env.Opt{
		TransformFunc: func(k, v string) (string, any) {
			k = strings.ReplaceAll(strings.ToLower(strings.TrimPrefix(k, "TEMPLATE_")), "_", ".")
			return k, v
		},
	}), nil); err != nil {
		return nil, err
	}

	return k, nil
}
