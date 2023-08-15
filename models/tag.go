package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	Id        uint       `json:"id" gorm:"primary_key" `
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Created   int        `json:"created" gorm:"autoCreateTime"`
	Updated   int        `json:"updated" gorm:"autoUpdateTime"`
	Name      string     `json:"name"`     // 标签名称
	BlogId    uint       `json:"blog_id"`  // 博客id
	ImageId   uint       `json:"image_id"` // 图片id
	UserId    uint       `json:"user_id"`  // 该标签作者
}

func (tag *Tag) TableName() string {
	return "tag"
}

func (tag *Tag) BeforeCreate(db *gorm.DB) (err error) {
	tag.Id = uint(uuid.New().ID())
	return
}
