package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
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

func (i *RichImage) InsertImages() (err error) {
	if i == nil {
		return ErrNilMethodReceiver
	}

	if i.ImageID == "" {
		// i.ImageID, err = uuid.NewUUID()
		var id uuid.UUID
		id, err = uuid.NewUUID()
		if err != nil {
			return
		}

		i.ImageID = id.String()
	}

	if err = core.DB.Create(i).Error; err != nil {
		return err
	}

	return
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

func (i *RichImage) GetRichImages(limit int) (images []RichImage, err error) {
	if i == nil {
		return nil, ErrNilMethodReceiver
	}

	if limit <= 0 || limit > queryMaxLimit {
		limit = queryMaxLimit
	}

	if err = core.DB.Where("updated_at < ? AND tags @> ?", i.UpdatedAt, i.Tags).Limit(limit).Find(&images).Error; err != nil {
		return
	}

	return
}
