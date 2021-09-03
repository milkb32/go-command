package sync_articles

import (
	"fmt"
	"github.com/milkb32/go-command/cmds"
	"github.com/milkb32/go-command/commands/articles/articles_mysql"
	"github.com/milkb32/go-command/commands/articles/articles_es"
	"github.com/urfave/cli"
	"time"
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
	i := 0
	for ;i <= 100; {
		err := RunOnce()
		if err != nil {
			fmt.Println("error:", err.Error())
		}

		time.Sleep(time.Duration(5) * time.Second)
		i += 1
	}

	return nil
}

func RunOnce() error {
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

	if len(articles) == 0 {
		fmt.Println("暂无新的文章，当前最大id:", maxId)
		return nil
	}

	fmt.Println("同步文章数量：", len(articles))
	// 将文章写入es
	articles_es.SaveToEs(articles)

	return nil
}


