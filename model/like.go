package model

import "backend/global"

type Like struct {
	global.GORM_MODEL
	TweetID uint `json:"tweet_id"`
	UserID  uint `json:"user_id"`
}
