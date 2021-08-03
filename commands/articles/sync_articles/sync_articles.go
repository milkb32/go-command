package sync_articles

import (
	"fmt"
	"github.com/milkb32/go-command/cmds"
	"github.com/urfave/cli"
)

// Init init cmds
func Init() {
	cmds.Register(
		cli.Command{
			Name: "article-sync",
			Usage: "博客文章从mysql同步到es",
			Action: func(c *cli.Context) error {
				return Run()
			},
		})
}

// Run run a cmd
func Run() error {
	fmt.Println("这是第一个command")

	return nil
}
