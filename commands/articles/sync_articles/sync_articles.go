package sync_articles

import (
	"github.com/milkb32/go-command/cmds"
	"github.com/urfave/cli"
)

func Init() {
	cmds.Register(
		cli.Command{
			Name:                   "article-sync",
			Usage:              "博客文章从mysql同步到es",
			Action: func(c cli.Context) error {
				return Run()
			},
		})
}

func Run() error {
	return nil
}
