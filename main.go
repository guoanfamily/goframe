package main

import ("goframe/router"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"goframe/common"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"goframe/cache"
)
func init() {
	var err error

	common.Db,err = gorm.Open("mysql", "root:@tcp(172.16.4.12:3306)/test?charset=utf8&parseTime=True&loc=Local")
	common.Db.SingularTable(true)
	common.Db.LogMode(true)
	if err != nil {
		panic("failed to connect database")
	}
	common.Db.DB().SetMaxIdleConns(10)
	common.Db.DB().SetMaxOpenConns(100)
	common.Db.DB().Ping()
	//redis插件注册
	cache.Regist()
	//redis init
	common.Rds, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

}
func main(){
	//redis测试代码
	_, err := common.Rds.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(common.Rds.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
	router.Router()

}