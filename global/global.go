package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	GORM_DB *gorm.DB
    GO_REDIS *redis.Client
)