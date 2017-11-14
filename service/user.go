package service

import (
	"time"
	"goframe/common"
	//"goframe/cache"
	//"fmt"
)

type SubObject struct {
	isCache bool
	Id string
}

type Usertable struct {
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
type UserRole struct{
	Username string
	Rolename string

}
func Frist() Usertable{
	var user Usertable
	//user.isCache = true
	common.Db.Where("true=true").First(&user,2)
	return user
}

func Find() []Usertable{
	var users []Usertable
	//user.isCache = true
	common.Db.Where("true=true").Find(&users)
	return users
}

func MQuery() []UserRole{
	var userrole []UserRole
	sql := "select u.`name` username, r.`name` rolename from usertable u join user_role ur on u.id=ur.user_id join role r on ur.role_id = r.id WHERE true=true & u.sex =? "
	common.Db.Raw(sql,2).Scan(&userrole)
	return userrole
}

