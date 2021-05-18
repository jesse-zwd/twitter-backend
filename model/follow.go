package model

import "backend/global"

type Follow struct {
	global.GORM_MODEL
	Following uint `json:"following"`
	Follower  uint `json:"follower"`
}
