package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrNilTagName = errors.New("no tag name specified")
)

type Tag struct {
	ID        int32          `gorm:"column:id;primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Aliases pq.StringArray `gorm:"column:aliases;type:text[]" json:"aliases"`
	TagNameEN string `gorm:"column:tag_name_en"`
	TagNameCN string `gorm:"column:tag_name_cn"`
	TagNameJP string `gorm:"column:tag_name_jp"`
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

func insertTags(db *gorm.DB, tags []Tag) error {
	return db.Clauses(clause.OnConflict{DoNothing: true}).Table("tag").Create(&tags).Error
}

func updateTags(db *gorm.DB, tags []Tag) error {
	return db.Save(tags).Error
}
