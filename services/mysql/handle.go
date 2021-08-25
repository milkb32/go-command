package mysql

import (
	"database/sql"
	"errors"
)

var (
	conns map[string]*sql.DB
)

func GetMysqlDB(name string) (*sql.DB, error) {
	if name == "" {
		return nil, errors.New("请输入mysql连接名")
	}
	db, err := sql.Open("mysql", "root:1234567890@/test?charset=utf8")
	if err != nil {
		return nil, err
	}


}
