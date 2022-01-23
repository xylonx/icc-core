package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ImageTagRelation struct {
	ImageID string `gorm:"column:image_id;primaryKey"`
	TagName string `gorm:"column:tag_name;primaryKey"`
}

func (*ImageTagRelation) TableName() string {
	return "image_tag_relation"
}

func addTagsForImage(db *gorm.DB, imageID string, tags []string) error {
	itrs := make([]ImageTagRelation, len(tags))
	for i := range itrs {
		itrs[i].ImageID = imageID
		itrs[i].TagName = tags[i]
	}
	return db.Clauses(clause.OnConflict{DoNothing: true}).Table("image_tag_relation").Create(&itrs).Error
}

func deleteTagsForImage(db *gorm.DB, imageID string, tags []string) error {
	return db.Where("image_id = ? AND tag_name IN ?", imageID, tags).Delete(&ImageTagRelation{}).Error
}

func deleteImageAllTags(db *gorm.DB, imageID string) error {
	return db.Where("image_id = ?", imageID).Delete(&ImageTagRelation{}).Error
}
