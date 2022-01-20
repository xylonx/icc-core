package model

import (
	"errors"
	"math"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	ErrNilMethodReceiver    = errors.New("method receiver is nil")
	ErrNilImageID           = errors.New("image is nil")
	ErrNilTag               = errors.New("tag is nil")
	ErrUnsupportedImageType = errors.New("image type is not supported now")
	ErrDuplicateImage       = errors.New("image is duplicated")

	_minTimeInt64 = time.Unix(0, math.MinInt64)
	_maxTimeInt64 = time.Unix(0, math.MaxInt64)
)

const queryMaxLimit = 100

type RichImage struct {
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	ImageID    string         `gorm:"column:image_id;primaryKey" json:"image_id"`
	ExternalID string         `gorm:"column:external_id" json:"-"`
	MD5Sum     string         `gorm:"column:md5_sum;unique" json:"md5_sum"` // for shrinking file duplication
	Tags       pq.StringArray `gorm:"column:tags;type:text[]" json:"tags"`
	Limit      int            `gorm:"-"`
}

func (*RichImage) TableName() string {
	return "rich_image"
}

func (i *RichImage) getRichImages(db *gorm.DB) (images []RichImage, err error) {
	if i == nil {
		return nil, ErrNilMethodReceiver
	}
	if i.UpdatedAt.Before(_minTimeInt64) || i.UpdatedAt.After(_maxTimeInt64) {
		i.UpdatedAt = time.Now()
	}

	if i.Limit <= 0 || i.Limit > queryMaxLimit {
		i.Limit = queryMaxLimit
	}

	if i.Tags == nil {
		err = db.Table(i.TableName()).Where("updated_at < ?", i.UpdatedAt).Order("updated_at desc").Limit(i.Limit).Scan(&images).Error
	} else {
		err = db.Table(i.TableName()).Where("updated_at < ? AND tags @> ?", i.UpdatedAt, pq.Array(i.Tags)).
			Order("updated_at desc").Limit(i.Limit).Scan(&images).Error
	}
	if err != nil {
		return
	}

	return
}
