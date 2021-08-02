package commands

import "github.com/milkb32/go-command/commands/articles/sync_articles"

func Register() {
	sync_articles.Init()
}
