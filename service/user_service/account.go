package user_service

import (
	"app/dao"
	"app/models"
	"app/pkg/app"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountSvc struct {
}

func (*AccountSvc) Register(user *models.User) error {
	if user.Name == "" || user.Password == "" {
		return app.NewError("信息不完整!")
	}
	return dao.Insert("users", user)
}

func (*AccountSvc) Login(user *models.User) error {
	err := dao.FindOne("users", bson.M{"name": user.Name, "password": user.Password}, user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return app.NewError("用户不存在")
		}
		return err
	}
	if user.Id.IsZero() {
		return app.NewError("用户不存在")
	}
	return nil
}
