package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Tag struct {
	Id        uint       `json:"id" form:"id" gorm:"primaryKey"`
	Uid       string     `json:"uid" form:"uid"`
	Created   int64      `gorm:"autoCreateTime:milli"`
	Updated   int64      `gorm:"autoUpdateTime:milli"`
	DeletedAt *time.Time `json:"deleted_at"`
	// 字段
	TagName string `json:"tag_name" form:"tag_name"`
	Type    string `json:"type" form:"type" gorm:"type:enum('image', 'article')"`
	UserId  uint   `json:"user_id" form:"user_id"`
	// 关系
	User    User       `json:"user" form:"user" gorm:"reference:UserId"`
	Article []*Article `json:"articles" gorm:"many2many:article_tag"`
	Image   []*Image   `json:"images" gorm:"many2many:image_tag"`
}

func (t *Tag) TableName() string {
	return "tag"
}

func (t *Tag) BeforeCreate(*gorm.DB) (err error) {
	uid := uuid.NewV1()
	t.Uid = strings.Replace(uid.String(), "-", "", -1)
	return nil
}