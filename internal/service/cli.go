package service

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/knadh/koanf/v2"
	"github.com/leonelquinteros/gotext"
	"github.com/urfave/cli/v3"

	"github.com/acepanel/backup-template/internal/helper"
	"github.com/acepanel/backup-template/pkg/types"
)

type CliService struct {
	t    *gotext.Locale
	conf *koanf.Koanf
}

func NewCliService(t *gotext.Locale, conf *koanf.Koanf) *CliService {
	return &CliService{
		t:    t,
		conf: conf,
	}
}

func (s *CliService) Check(ctx context.Context, cmd *cli.Command) error {
	// TODO

	helper.Success(nil)

	return nil
}

func (s *CliService) Mkdir(ctx context.Context, cmd *cli.Command) error {
	for _, dir := range cmd.Args().Slice() {
		if dir == "" {
			return errors.New(s.t.Get("directory path is required"))
		}
		// TODO
	}

	helper.Success(nil)

	return nil
}

func (s *CliService) Deldir(ctx context.Context, cmd *cli.Command) error {
	for _, dir := range cmd.Args().Slice() {
		if dir == "" {
			return errors.New(s.t.Get("directory path is required"))
		}
		// TODO
	}

	helper.Success(nil)

	return nil
}

func (s *CliService) Put(ctx context.Context, cmd *cli.Command) error {
	localPath := cmd.Args().Get(0)
	remotePath := cmd.Args().Get(1)
	if localPath == "" || remotePath == "" {
		return errors.New(s.t.Get("local and remote file paths are required"))
	}

	// TODO

	helper.Success(nil)

	return nil
}

func (s *CliService) Del(ctx context.Context, cmd *cli.Command) error {
	for _, file := range cmd.Args().Slice() {
		if file == "" {
			return errors.New(s.t.Get("file path is required"))
		}
		// TODO
	}

	helper.Success(nil)

	return nil
}

func (s *CliService) Get(ctx context.Context, cmd *cli.Command) error {
	// TODO

	helper.Success(nil)

	return nil
}

func (s *CliService) Ls(ctx context.Context, cmd *cli.Command) error {
	path := cmd.Args().First()

	if filepath.Dir(path) != "." {
		var files types.ListDir
		// TODO
		helper.Success(files)
	} else {
		var file types.ListFile
		// TODO
		helper.Success(file)
	}

	return nil
}
