package model

import "backend/global"

type Comment struct {
	global.GORM_MODEL
	Text    string `json:"text"`
	TweetID uint   `json:"tweet_id"`
	UserID  uint   `json:"user_id"`
}
