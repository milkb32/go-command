package es64

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic"
	"github.com/spf13/viper"
)

type EsHandle struct {
	Client *elastic.Client

	ESUrl   string
	ESIndex string
	ESType  string
	Body map[string]interface{}
	Id string

	OperateType string // insert/update/delete

	// for update
	FieldUpdateKey string

	ctx context.Context
}

// NewArticleEsHandle 初始化article es
func NewArticleEsHandle() (*EsHandle, error) {
	esHandle := &EsHandle{
		Client:         nil,
		ESUrl:          viper.GetString("es.articles.host"),
		ESIndex:        viper.GetString("es.articles.index"),
		ESType:         "_doc",
		Body:           make(map[string]interface{}),
		Id:             "",
		OperateType:    "",
		FieldUpdateKey: "",
		ctx:            context.Background(),
	}

	// 初始化Client
	err := esHandle.NewClient()
	if err != nil {
		return esHandle, err
	}

	return esHandle, nil
}

func (es *EsHandle) NewClient() error {
	if es.ESUrl == "" || es.ESIndex == "" {
		return errors.New("es连接信息为空")
	}

	client, err := elastic.NewClient(
		elastic.SetURL(es.ESUrl),
		elastic.SetSniff(false),
	)
	if err != nil {
		return err
	}

	es.Client = client
	return nil
}

// Insert
func (es *EsHandle) Insert(data map[string]interface{}) error {
	bodyJson, err := json.Marshal(data)
	if err != nil {
		return errors.New("")
	}
	_, err = es.Client.Index().Index(es.ESIndex).Type("_doc").Id(es.Id).
		BodyJson(string(bodyJson)).Do(es.ctx)

	return err
}

// Query
func (es *EsHandle) Query(boolQuery *elastic.BoolQuery, source []string, size int, sortBy string, ascending bool) (*elastic.SearchResult, error) {
	searchService := es.Client.Search(es.ESIndex).Query(boolQuery)
	if sortBy != "" {
		searchService.Sort(sortBy, ascending)
	}

	if size > 0 {
		searchService.Size(size)
	}
	searchFieldsContext := elastic.NewFetchSourceContext(true).Include(source...)
	searchResult, err := searchService.FetchSourceContext(searchFieldsContext).Do(es.ctx)
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}

// Update
func (es *EsHandle) Update() error {
	return nil
}
