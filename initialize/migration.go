package initialize

import (
	"backend/global"
	"backend/model"
)

//database migration

func migration() {
	global.GORM_DB.AutoMigrate(&model.User{})
	global.GORM_DB.AutoMigrate(&model.Comment{})
	global.GORM_DB.AutoMigrate(&model.Follow{})
	global.GORM_DB.AutoMigrate(&model.Like{})
	global.GORM_DB.AutoMigrate(&model.Tweet{})
	global.GORM_DB.AutoMigrate(&model.TweetFile{})
	global.GORM_DB.AutoMigrate(&model.Retweet{})
}
