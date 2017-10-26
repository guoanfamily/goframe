package service

import (
	"time"
	"goframe/common"
)

type User struct {
	Id           string
	Name         string
	Username     string
	Password     string
	Phone        string
	Sex          string
	Status       string
	CredentialNo string
	Email        string
	AliAccount   string
	Portrait     string
	OrgId        string
	Introduce    string
	CreateTime   time.Time
	CityId       string
	ValidFlag    string
}

func Frist() User{
	var user User
	common.Db.First(&user, 1)
	return user
}

