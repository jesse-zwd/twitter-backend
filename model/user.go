package model

import (
	"backend/global"
	"time"
)

// User model
type User struct {
	global.GORM_MODEL
	UserName   string    `gorm:"not null;unique" json:"username"`
	Password   string    `json:"-"`
	Nickname   string    `gorm:"not null;unique" json:"nickname"`
	Status     string    `json:"status"`
	Avatar     string    `gorm:"size:1000" json:"avatar"`
	Email      string    `gorm:"not null;unique" json:"email"`
	Bio        string    `gorm:"size:2000" json:"bio"`
	Website    string    `json:"website"`
	CoverPhoto string    `json:"coverPhoto"`
	Dob        time.Time `json:"dob"`
	Location   string    `json:"location"`
	Tweets     []Tweet   `gorm:"foreignKey:UserID;references:ID" json:"tweets"`
}
