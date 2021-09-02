package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)


func NewArticleDB() (*sql.DB, error) {
	return GetMysqlDB("articles")
}

func GetMysqlDB(name string) (*sql.DB, error) {
	if name == "" {
		return nil, errors.New("请输入mysql连接名")
	}

	host := viper.GetString("mysql." + name + ".host")
	port := viper.GetString("mysql." + name + ".port")
	user := viper.GetString("mysql." + name + ".user")
	passwd := viper.GetString("mysql." + name + ".passwd")
	database := viper.GetString("mysql." + name + ".db")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&timeout=5000ms",
		user, passwd, host, port, database))
	if err != nil {
		return nil, err
	}

	return db, err
}
