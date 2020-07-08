package models

import (
	"app/dao"
	"app/pkg/e"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `bson:"_id,omitempty"`
	Name      string        `bson:"name" form:"name" json:"name" binding:"required"`
	Password  string        `bson:"password" form:"password" json:"password" binding:"required"`
	CreatedAt time.Time     `bson:"createdAt"` // 创建日期
	UpdatedAt time.Time     `bson:"updatedAt"` // 更新日期
}

func (user *User) Save() error {
	user.Id = bson.NewObjectId()
	return dao.Create("users", user)
}

func (user *User) Login() error {
	err := dao.FindOne("users", user, bson.M{"name": user.Name, "password": user.Password})
	if err != nil {
		return err
	}
	if user.Id.Hex() == "" {
		return e.NewError("用户不存在")
	}
	return nil
}
