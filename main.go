package main

import ("goframe/router"
	"github.com/guoanfamily/r-gorm"
	_ "github.com/guoanfamily/r-gorm/dialects/mysql"
	"goframe/common"
	"fmt"
	"github.com/go-redis/redis"
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
	//redis init
	common.Rds = gorm.RedisOpen(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := common.Rds.Ping().Result()
	fmt.Println(pong, err)

}
func main(){
	//redis测试代码
	err := common.Rds.Set("mykey", "superWang",0).Err()
	if err != nil {
		panic(err)
	}

	username, err := common.Rds.Get("mykey").Result()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
	router.Router()

}