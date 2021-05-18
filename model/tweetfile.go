package model

import "backend/global"

type TweetFile struct {
	global.GORM_MODEL
	Url     string `json:"url"`
	TweetID uint   `json:"tweet_id"`
}
