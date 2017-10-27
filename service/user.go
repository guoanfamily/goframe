package service

import (
	"time"
	"goframe/common"
	//"goframe/cache"
	//"fmt"
)

type SubObject struct {
	Id string
}

type User struct {
	SubObject
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
	//OrgId        string
	//Introduce    string
	CreateTime   time.Time
	//CityId       string
	//ValidFlag    string
}

//func (s *SubObject) AfterFind() (err error) {
//	fmt.Println("AfterFind")
//	return
//}
func Frist() User{
	var user User
	common.Db.First(&user, 1)
	return user
}

