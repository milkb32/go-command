package articles_es

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/milkb32/go-command/commands/articles"
	"github.com/milkb32/go-command/common"
	"github.com/milkb32/go-command/services/es64"
	"github.com/olivere/elastic"
	"strconv"
)

// 获取es最大的文章id
func GetMaxIdFromEs() (int, error) {
	esHandle, err := es64.NewArticleEsHandle()
	if err != nil {
		return 0, errors.New("获取es handle失败:" + err.Error())
	}
	boolQuery := elastic.NewBoolQuery()
	result, err := esHandle.Query(boolQuery, []string{"aid", "title"}, 1, "aid", false)
	if err != nil {
		return 0, errors.New("查询es失败:" + err.Error())
	}

	if result.TotalHits() == 0 {
		return 0, nil
	}

	maxItem := result.Hits.Hits
	article := &articles.Article{}
	err = json.Unmarshal(*maxItem[0].Source, article)
	if err != nil {
		return 0, err
	}

	return article.Aid, nil
}

func SaveToEs(articles []articles.Article) error {
	esHandle, err := es64.NewArticleEsHandle()
	if err != nil {
		return err
	}

	for _, article := range articles {
		esHandle.Id = strconv.Itoa(article.Aid)
		artType, _ := strconv.Atoi(article.Type)
		toEsMap := map[string]interface{}{
			"aid" : article.Aid,
			"author_id": 0,
			"author_name": "",
			"tags": []interface{}{},
			"title" : article.Title,
			"content" : article.Content,
			"desc" : article.Desc,
			"type" : artType,
			"ctime" : common.TimestampToDate(int64(article.Ctime)),
			"utime" : common.TimestampToDate(int64(article.Mtime)),
			"is_valid" : article.IsValid,
		}
		err = esHandle.Insert(toEsMap)
		if err != nil {
			fmt.Print(err.Error())
			continue
		}
	}

	return nil
}

