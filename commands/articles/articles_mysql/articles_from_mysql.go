package articles_mysql

import (
	"github.com/milkb32/go-command/commands/articles"
	"github.com/milkb32/go-command/services/mysql"
)

func GetArticlesFromMysqlById(minId int, size int) ([]articles.Article, error) {
	dbHandle, err := mysql.NewArticleDB()
	if err != nil {
		return nil, err
	}

	defer dbHandle.Close()

	rows, err := dbHandle.Query("select id,title,`desc`,content,`type`,ctime,mtime,is_valid from articles where id >? limit ?", minId, size)
	if err != nil {
		return nil, err
	}
	articleArr := []articles.Article{}
	for rows.Next() {
		if err != nil {
			continue
		}

		thisArticle := articles.Article{
			Aid:     0,
			Title:   "",
			Desc:    "",
			Content: "",
			Type:    "",
			Ctime:   0,
			Mtime:   0,
			IsValid: 0,
		}
		_ = rows.Scan(&thisArticle.Aid, &thisArticle.Title, &thisArticle.Desc, &thisArticle.Content, &thisArticle.Type,
			&thisArticle.Ctime, &thisArticle.Mtime, &thisArticle.IsValid)
		articleArr = append(articleArr, thisArticle)
	}

	return articleArr, nil
}

func Format() {

}
