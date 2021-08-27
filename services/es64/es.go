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
		ESUrl:          viper.GetString("es.articles.hosts"),
		ESIndex:        viper.GetString("es.articles.index"),
		ESType:         "",
		Body:           make(map[string]interface{}),
		Id:             "",
		OperateType:    "",
		FieldUpdateKey: "",
		ctx:            *new(context.Context),
	}

	// 初始化Client
	err := esHandle.NewClient()
	if err != nil {
		return esHandle, errors.New("初始化es client失败")
	}

	return esHandle, nil
}

func (es *EsHandle) NewClient() error {
	if es.ESUrl == "" || es.ESIndex == "" {
		return errors.New("es连接为空")
	}

	client, err := elastic.NewClient(
		elastic.SetURL(es.ESUrl))
	if err != nil {
		return errors.New("es连接初始化出错")
	}

	es.Client = client
	return nil
}

// Insert
func (es *EsHandle) Insert() error {
	bodyJson, err := json.Marshal(es.Body)
	if err != nil {
		return errors.New("")
	}
	_, err = es.Client.Index().Index(es.ESIndex).Type(es.ESType).Id(es.Id).
		BodyJson(bodyJson).Refresh("true").Do(es.ctx)

	return err
}

// Query
func (es *EsHandle) Query(boolQuery *elastic.BoolQuery, size int, sortBy string, ascending bool) (*elastic.SearchResult, error) {
	searchService := es.Client.Search(es.ESIndex).Query(boolQuery)
	if sortBy != "" {
		searchService.Sort(sortBy, ascending)
	}

	if size > 0 {
		searchService.Size(size)
	}
	searchResult, err := searchService.Pretty(true).Do(es.ctx)
	if err != nil {
		return nil, errors.New("query出错")
	}

	return searchResult, nil
}

// Update
func (es *EsHandle) Update() error {
	return nil
}
