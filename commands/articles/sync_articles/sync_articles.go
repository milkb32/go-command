package sync_articles

import (
	"fmt"
	"github.com/milkb32/go-command/cmds"
	"github.com/milkb32/go-command/commands/articles/articles_mysql"
	"github.com/milkb32/go-command/commands/articles/articles_es"
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
	fmt.Println("start....")
	// 获取es中最小的文章id
	maxId, err := articles_es.GetMaxIdFromEs()
	if err != nil {
		return err
	}

	// 从mysql一次获取100篇文章
	articles, err := articles_mysql.GetArticlesFromMysqlById(maxId, 100)
	if err != nil {
		return err
	}
	fmt.Println(articles)
	// 将文章写入es
	articles_es.SaveToEs(articles)


	return nil
}
