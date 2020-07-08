package models

import (
	"app/dao"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"gopkg.in/mgo.v2/bson"
)

type Logger struct {
	Id         bson.ObjectId `bson:"_id,omitempty"`
	User       bson.ObjectId `bson:"user,omitempty"`
	StartTime  time.Time     `bson:"startTime"`
	EndTime    time.Time     `bson:"endTime"`
	UseTime    string        `bson:"useTime"`
	IP         string        `bson:"ip"`
	Method     string        `bson:"method"`
	Url        string        `bson:"url"`
	StatusCode int           `bson:"statusCode"`
	Type       string        `bson:"type"`
	ErrMsg     string        `bson:"errMsg"`
	StackInfo  string        `bson:"stackInfo"`
	CreatedAt  time.Time     `bson:"createdAt"` // 创建日期
	UpdatedAt  time.Time     `bson:"updatedAt"` // 更新日期
}

func (logger *Logger) Start() {
	logger.StartTime = time.Now()
}

func (logger *Logger) End(c *gin.Context) {

	// 结束时间
	logger.EndTime = time.Now()

	// 执行时间
	logger.UseTime = logger.EndTime.Sub(logger.StartTime).String()

	// 请求方式
	logger.Method = c.Request.Method

	// 请求路由
	logger.Url = c.Request.RequestURI

	// 状态码
	logger.StatusCode = c.Writer.Status()

	// 请求IP
	logger.IP = c.ClientIP()

	if logger.ErrMsg != "" {
		logger.Type = "error"
	} else {
		logger.Type = "normal"
	}

	if err := dao.Create("requestLogs", logger); err != nil {
		fmt.Println("日志记录失败")
	}
}
