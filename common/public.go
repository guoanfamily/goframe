package common

import (
	"github.com/jinzhu/gorm"
	"github.com/go-redis/redis"
)

var Db *gorm.DB
var Rds *redis.Client
