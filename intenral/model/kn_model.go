package model

import (
	"time"
)

type Forest struct {
	UID           uint      `gorm:"primaryKey;column:uid"`
	Title         string    `gorm:"column:knowledge_title;size:200;not null"`
	OwnerUID      uint      `gorm:"column:owner_uid;not null"` 
	RegisteredAt  time.Time `gorm:"column:registered_datetime;autoCreateTime"`
}

func (Forest) TableName() string {
	return "kn_forest"
}

type Tree struct {
	UID           uint      `gorm:"primaryKey;column:uid"`
	ForestUID     uint      `gorm:"column:forest_uid;not null"`   // FK to kn_forest.uid
	CommentUID    uint      `gorm:"column:comment_uid;not null"`  // FK to kn_comment.uid
	RegisteredAt  time.Time `gorm:"column:registered_datetime;autoCreateTime"`
}

func (Tree) TableName() string {
	return "kn_tree"
}

type Comment struct {
	UID           uint      `gorm:"primaryKey;column:uid"`
	TreeUID       uint      `gorm:"column:tree_uid;not null"`    
	UserUID       uint      `gorm:"column:user_uid;not null"`    
	Content       string    `gorm:"column:content;type:text"`
	RegisteredAt  time.Time `gorm:"column:registered_datetime;autoCreateTime"`
}

func (Comment) TableName() string {
	return "kn_comment"
}
