package model

import (
	"time"

	"gorm.io/gorm"
)

// TagAlias - just a map: map 'alias' to 'name'
type TagAlias struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Alias string `gorm:"column:alias;primaryKey"`
	Name  string `gorm:"column:name"`
}

func (*TagAlias) TableName() string {
	return "tag_alias"
}
