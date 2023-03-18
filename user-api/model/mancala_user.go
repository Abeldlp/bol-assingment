package model

import (
	"github.com/Abeldlp/bol-assignment/user-api/config"
	"gorm.io/gorm"
)

type MancalaUser struct {
	Username string `json:"username"`
	Wins     int    `json:"wins"`
}

func GetUserById(user *MancalaUser, id string) *gorm.DB {
	return config.DB.Where("username = ?", id).First(&user)
}

func CreateUser(user *MancalaUser) *gorm.DB {
	return config.DB.Create(&user)
}
