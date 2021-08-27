package sync_articles

import (
	"errors"
	"fmt"
	"github.com/milkb32/go-command/cmds"
	"github.com/milkb32/go-command/commands/articles/articles_to_es"
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
	maxId, err := articles_to_es.GetMaxIdFromEs()
	if err != nil {
		return errors.New("GetMaxIdFromEs error")
	}

	fmt.Println(maxId)

	return nil
}
