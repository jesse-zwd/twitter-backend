package model

import "backend/global"

type Tweet struct {
	global.GORM_MODEL
	Text       string      `json:"text"`
	Tags       string      `json:"tags"`
	TweetFiles []TweetFile `gorm:"foreignKey:TweetID;references:ID" json:"files"`
	Comments   []Comment   `gorm:"foreignKey:TweetID;references:ID" json:"comments"`
	UserID     uint        `json:"user_id"`
}
