package models

import "time"

type Follow struct {
	Id           uint       `gorm:"primary_key" json:"id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
	Created      int        `json:"created" gorm:"autoCreateTime"`
	Updated      int        `json:"updated" gorm:"autoUpdateTime"`
	FollowUserId uint       `json:"follow_user_id"` // 被关注者 id
	UserId       uint       `json:"user_id"`        // 关注者 id
	FollowType   bool       `json:"follow_type"`    // true 关注 false 取消关注
}

func (Follow) TableName() string {
	return "follow"
}
