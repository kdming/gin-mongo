package v1

import "app/service/user_service"

var userSvc *user_service.UserSvc
var tokenSvc *user_service.TokenSvc

func init() {
	userSvc = &user_service.UserSvc{}
	tokenSvc = &user_service.TokenSvc{}
}
