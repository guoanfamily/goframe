package main

import ("goframe/router"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goframe/common"
)
func init() {
	var err error

	common.Db,err = gorm.Open("mysql", "root:@tcp(172.16.4.12:3306)/test?charset=utf8&parseTime=True&loc=Local")
	common.Db.SingularTable(true)
	if err != nil {
		panic("failed to connect database")
	}

}
func main(){
	router.Router()
}