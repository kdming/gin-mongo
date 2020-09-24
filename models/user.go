package models

import (
	"app/dao"
	"app/pkg/e"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/qiniu/qmgo/field"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	field.DefaultField `bson:",inline"`
	Name               string `bson:"name" json:"name" binding:"required"`
	Password           string `bson:"password" json:"password" binding:"required"`
}

func (user *User) Save() error {
	return dao.Insert("users", user)
}

func (user *User) Login() error {
	err := dao.FindOne("users", bson.M{"name": user.Name, "password": user.Password}, user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return e.NewError("用户不存在")
		}
		return err
	}
	if user.Id.IsZero() {
		return e.NewError("用户不存在")
	}
	return nil
}
