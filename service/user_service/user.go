package user_service

import (
	"app/common/app"
	"app/dao"
	"app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserSvc struct {
}

func (*UserSvc) Register(user *models.User) error {
	return dao.Insert("users", user)
}

func (*UserSvc) Login(user *models.User) error {
	err := dao.FindOne("users", bson.M{"name": user.Name, "password": user.Password}, user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return app.NewError("用户不存在")
		}
		return err
	}
	return nil
}
