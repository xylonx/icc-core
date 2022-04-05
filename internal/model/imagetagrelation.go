package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ImageTagRelation struct {
	ImageID string `gorm:"column:image_id;primaryKey"`
	// TagName string `gorm:"column:tag_name;primaryKey"`
	TagID int32 `gorm:"column:tag_id;primaryKey"`
}

func (*ImageTagRelation) TableName() string {
	return "image_tag_relation"
}

func addTagsForImage(db *gorm.DB, imageID string, tags []Tag) error {
	itrs := make([]ImageTagRelation, len(tags))
	for i := range itrs {
		itrs[i].ImageID = imageID
		itrs[i].TagID = tags[i].ID
	}
	return db.Clauses(clause.OnConflict{DoNothing: true}).Table("image_tag_relation").Create(&itrs).Error
}

func deleteTagsForImage(db *gorm.DB, imageID string, tagIds []int32) error {
	return db.Where("image_id = ? AND tag_id IN ?", imageID, tagIds).Delete(&ImageTagRelation{}).Error
}

func deleteImageAllTags(db *gorm.DB, imageID string) error {
	return db.Where("image_id = ?", imageID).Delete(&ImageTagRelation{}).Error
}
