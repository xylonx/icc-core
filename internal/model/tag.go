package model

import (
	"errors"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrNilTagName = errors.New("no tag name specified")
)

type Tag struct {
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Aliases pq.StringArray `gorm:"column:aliases;type:text[]" json:"aliases"`
	TagName string         `gorm:"column:tag_name;primaryKey" json:"tag_name"`
}

func (*Tag) TableName() string {
	return "tag"
}

func getAllTags(db *gorm.DB) (tags []Tag, err error) {
	if err = db.Table("tag").Find(&tags).Error; err != nil {
		return
	}
	return
}

func insertTags(db *gorm.DB, tagNames []string) error {
	tags := make([]Tag, len(tagNames))
	for i := range tags {
		tags[i].TagName = tagNames[i]
	}
	return db.Clauses(clause.OnConflict{DoNothing: true}).Table("tag").Create(&tags).Error
}
