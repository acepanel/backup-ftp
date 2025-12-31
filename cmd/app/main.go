package main

import "os"

func main() {
	if os.Geteuid() != 0 {
		panic("must run as root")
	}

	cli, err := initCli()
	if err != nil {
		panic(err)
	}

	cli.Run()
}
