package dao

import (
	"app/pkg/config"
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

// 全局Session
var GlobalMgoSession *mgo.Session

// 数据库名称
var DataBase string

// 连接数据库
func Connect() bool {

	// 读取配置文件，初始化mongo链接配置信息
	conf := config.GetConfig()
	DataBase = conf.DB_NAME

	info := &mgo.DialInfo{}
	if conf.DB_PWD != "" {
		info = &mgo.DialInfo{
			Addrs:    []string{conf.DB_HOST},
			Timeout:  60 * time.Second,
			Database: conf.DB_NAME,
			Username: conf.DB_USER,
			Password: conf.DB_PWD,
		}
	} else {
		info = &mgo.DialInfo{
			Addrs:    []string{conf.DB_HOST},
			Timeout:  60 * time.Second,
			Database: conf.DB_NAME,
		}
	}

	// 链接数据库
	globalSession, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
		return false
	}
	GlobalMgoSession = globalSession
	GlobalMgoSession.SetMode(mgo.Monotonic, true)
	//default is 4096
	GlobalMgoSession.SetPoolLimit(100) // 设置session连接池最大值
	mgo.SetDebug(true)

	fmt.Println("数据库连接成功")

	return true

}

// 获取session
func GetSession(table string) (*mgo.Session, *mgo.Collection) {
	s := GlobalMgoSession.Copy()
	c := s.DB(DataBase).C(table)
	return s, c
}
