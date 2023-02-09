package models

import (
	"github.com/qiniu/qmgo/field"
)

type User struct {
	field.DefaultField `bson:",inline"`
	Name               string `bson:"name" json:"name" binding:"required"`
	Password           string `bson:"password" json:"password" binding:"required"`
	Role               int    `bson:"role"`
}
