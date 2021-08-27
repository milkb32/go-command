package articles_to_es

import (
	"errors"
	"fmt"
	"github.com/milkb32/go-command/services/es64"
	"github.com/olivere/elastic"
)

// 获取es最大的文章id
func GetMaxIdFromEs() (int, error) {
	esHandle, err := es64.NewArticleEsHandle()
	if err != nil {
		return 0, errors.New("获取es handle失败")
	}
	boolQuery := &elastic.BoolQuery{}
	result, err := esHandle.Query(boolQuery, 1, "id", false)
	if err != nil {
		return 0, errors.New("查询es失败")
	}

	if len(result.Hits.Hits) == 0 {
		return 0, nil
	}

	maxItem := result.Hits.Hits
	firstItem, err := maxItem[0].Source.MarshalJSON()
	if err != nil {
		return 0, errors.New("es查询结果解析失败")
	}

	fmt.Println(firstItem)

	//firstItemArr := firstItem

	return 0, nil
}