package common

import (
	"github.com/jinzhu/gorm"
	"github.com/garyburd/redigo/redis"
)

var Db *gorm.DB
var Rds redis.Conn
