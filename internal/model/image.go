package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrNilMD5Sum = errors.New("md5 sum is nil. not valid")
)

type Image struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ImageID    string `gorm:"column:image_id;primaryKey"`
	ExternalID string `gorm:"column:external_id"`
	MD5Sum     string `gorm:"column:md5_sum;unique"` // for shrinking file duplication

	Owner string `gorm:"column:owner"`
}

func (*Image) TableName() string {
	return "image"
}

func (i *Image) insertImage(db *gorm.DB) (err error) {
	if i == nil {
		return ErrNilMethodReceiver
	}

	if i.ImageID == "" {
		var id uuid.UUID
		id, err = uuid.NewUUID()
		if err != nil {
			return
		}
		i.ImageID = id.String()
	}

	return db.Create(i).Error
}

// if no duplicated image find, return nil
func (i *Image) checkMD5(db *gorm.DB) (err error) {
	if i == nil {
		return ErrNilMethodReceiver
	}

	if i.MD5Sum == "" {
		return ErrNilMD5Sum
	}

	if err = db.Model(i).Where("md5_sum = ?", i.MD5Sum).Take(i).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return ErrDuplicateImage
}

func (i *Image) deleteImage(db *gorm.DB) error {
	if i == nil {
		return ErrNilMethodReceiver
	}

	if i.ImageID == "" {
		return ErrNilImageID
	}

	if err := db.Model(i).Where("image_id = ?", i.ImageID).Delete(&i).Error; err != nil {
		return err
	}

	return nil
}
