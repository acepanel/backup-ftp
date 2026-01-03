package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/acepanel/backup-ftp/internal/helper"
	"github.com/acepanel/backup-ftp/pkg/types"
	"github.com/jlaffaye/ftp"
	"github.com/knadh/koanf/v2"
	"github.com/leonelquinteros/gotext"
	"github.com/urfave/cli/v3"
)

type CliService struct {
	t        *gotext.Locale
	conf     *koanf.Koanf
	host     string // FTP 服务器地址
	port     int    // FTP 端口，默认 21
	username string // 用户名
	password string // 密码
	basePath string // 基础路径
}

func NewCliService(t *gotext.Locale, conf *koanf.Koanf) *CliService {
	return &CliService{
		t:        t,
		conf:     conf,
		host:     conf.String("ftp.host"),
		port:     conf.Int("ftp.port"),
		username: conf.String("ftp.username"),
		password: conf.String("ftp.password"),
		basePath: conf.String("ftp.basepath"),
	}
}

func (s *CliService) Check(ctx context.Context, cmd *cli.Command) error {
	if s.port == 0 {
		return errors.New(s.t.Get("ftp port is required"))
	}

	conn, err := s.connect()
	if err != nil {
		return err
	}

	defer func(conn *ftp.ServerConn) { _ = conn.Quit() }(conn)

	helper.Success(nil)

	return nil
}

func (s *CliService) Mkdir(ctx context.Context, cmd *cli.Command) error {
	conn, err := s.connect()
	if err != nil {
		return err
	}

	defer func(conn *ftp.ServerConn) { _ = conn.Quit() }(conn)

	for _, dir := range cmd.Args().Slice() {
		if dir == "" {
			return errors.New(s.t.Get("directory path is required"))
		}
		if err = conn.MakeDir(dir); err != nil {
			return err
		}
	}

	helper.Success(nil)

	return nil
}

func (s *CliService) Deldir(ctx context.Context, cmd *cli.Command) error {
	conn, err := s.connect()
	if err != nil {
		return err
	}

	defer func(conn *ftp.ServerConn) { _ = conn.Quit() }(conn)

	for _, dir := range cmd.Args().Slice() {
		if dir == "" {
			return errors.New(s.t.Get("directory path is required"))
		}
		if err = conn.RemoveDir(dir); err != nil {
			return err
		}
	}

	helper.Success(nil)

	return nil
}

func (s *CliService) Put(ctx context.Context, cmd *cli.Command) error {
	conn, err := s.connect()
	if err != nil {
		return err
	}

	defer func(conn *ftp.ServerConn) { _ = conn.Quit() }(conn)

	localPath := cmd.Args().Get(0)
	remotePath := cmd.Args().Get(1)
	if localPath == "" || remotePath == "" {
		return errors.New(s.t.Get("local and remote file paths are required"))
	}

	// 确保远程目录存在
	if filepath.Dir(remotePath) != "." {
		_ = conn.MakeDir(filepath.Dir(remotePath))
	}

	data, err := os.Open(localPath)
	if err != nil {
		return err
	}

	defer func(data *os.File) { _ = data.Close() }(data)

	if err = conn.Stor(remotePath, data); err != nil {
		return err
	}

	helper.Success(nil)

	return nil
}

func (s *CliService) Del(ctx context.Context, cmd *cli.Command) error {
	conn, err := s.connect()
	if err != nil {
		return err
	}

	defer func(conn *ftp.ServerConn) { _ = conn.Quit() }(conn)

	for _, file := range cmd.Args().Slice() {
		if file == "" {
			return errors.New(s.t.Get("file path is required"))
		}
		if err = conn.Delete(file); err != nil {
			return err
		}
	}

	helper.Success(nil)

	return nil
}

func (s *CliService) Get(ctx context.Context, cmd *cli.Command) error {
	helper.Error(s.t.Get("not supported for performance reasons"))

	return nil
}

func (s *CliService) Ls(ctx context.Context, cmd *cli.Command) error {
	conn, err := s.connect()
	if err != nil {
		return err
	}

	defer func(conn *ftp.ServerConn) { _ = conn.Quit() }(conn)

	path := cmd.Args().First()
	if path == "" {
		path = "."
	}

	entries, err := conn.List(path)
	if err != nil {
		return err
	}

	if len(entries) > 1 {
		var files types.ListDir
		for _, entry := range entries {
			files = append(files, types.ListFile{
				Name: entry.Name,
				Size: int64(entry.Size),
				Time: entry.Time.Unix(),
			})
		}
		helper.Success(files)
		return nil
	} else if len(entries) == 1 {
		entry := entries[0]
		file := types.ListFile{
			Name: entry.Name,
			Size: int64(entry.Size),
			Time: entry.Time.Unix(),
		}
		helper.Success(file)
		return nil
	}

	helper.Error(s.t.Get("ftp server returned no entries"))

	return nil
}

func (s *CliService) connect() (*ftp.ServerConn, error) {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	conn, err := ftp.Dial(addr)
	if err != nil {
		return nil, err
	}

	err = conn.Login(s.username, s.password)
	if err != nil {
		_ = conn.Quit()
		return nil, err
	}

	return conn, nil
}
