package model

import (
	"errors"
	"time"

	"github.com/lib/pq"
	"github.com/xylonx/icc-core/internal/core"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrNilMethodReceiver = errors.New(" method receiver is nil")
	ErrNilImageID        = errors.New("image is nil")
	ErrNilTag            = errors.New("tag is nil")
)

const queryMaxLimit = 100

type RichImage struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ImageID    string         `gorm:"column:image_id;primaryKey"`
	ExternalID string         `gorm:"column:external_id;index"`
	Tags       pq.StringArray `gorm:"column:tags;type:text[]"`
}

func (*RichImage) TableName() string {
	return "image"
}

func (i *RichImage) UpsertTags() error {
	if i == nil {
		return ErrNilMethodReceiver
	}
	if i.ImageID == "" {
		return ErrNilImageID
	}

	if err := core.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(i).Error; err != nil {
		return err
	}

	return nil
}

func (i *RichImage) GetRichImagesBefore(limit int) (imgs []RichImage, err error) {
	if i == nil {
		return nil, ErrNilMethodReceiver
	}
	if limit <= 0 || limit > queryMaxLimit {
		limit = queryMaxLimit
	}

	if err = core.DB.Where("updated_at < ?", i.CreatedAt).Limit(limit).Find(&imgs).Error; err != nil {
		return
	}

	return
}

func (i *RichImage) GetRichImagesByTags(limit int) (imgs []RichImage, err error) {
	if i == nil {
		return nil, ErrNilMethodReceiver
	}
	if i.Tags == nil {
		return nil, ErrNilTag
	}

	if err = core.DB.Where("tags @> ?", i.Tags).Limit(limit).Find(&imgs).Error; err != nil {
		return
	}

	return
}
