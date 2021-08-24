package dao

import (
	"app/common/config"
	"context"
	"fmt"

	"github.com/qiniu/qmgo"
)

var client *qmgo.Client
var dbName string

func Connect() {
	conf := config.GetConfig()
	option := &qmgo.Config{
		Uri: "mongodb://" + conf.DB_USER + ":" + conf.DB_PWD + "@" + conf.DB_HOST + "/" + conf.DB_NAME,
	}
	dbName = conf.DB_NAME
	var err error
	ctx := context.Background()
	client, err = qmgo.NewClient(ctx, option)
	if err != nil {
		panic("mongodb connect error" + err.Error())
	}
	fmt.Println("mongodb connect success !")
}

func getColl(table string) *qmgo.Collection {
	return client.Database(dbName).Collection(table)
}
