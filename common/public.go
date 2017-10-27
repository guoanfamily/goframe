package common

import (
	"github.com/guoanfamily/r-gorm"
	"github.com/go-redis/redis"
)

var Db *gorm.DB
var Rds *redis.Client
