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
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Aliases pq.StringArray `gorm:"column:aliases"`
	Name    string         `gorm:"column:name;primaryKey"`
}

func (*Tag) TableName() string {
	return "tag"
}

func getAllTags(db *gorm.DB) (tags []Tag, err error) {
	if err = db.Find(&tags).Error; err != nil {
		return
	}
	return
}

func insertTags(db *gorm.DB, tagNames []string) error {
	tags := make([]Tag, len(tagNames))
	for i := range tags {
		tags[i].Name = tagNames[i]
	}
	return db.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&tags).Error
}
