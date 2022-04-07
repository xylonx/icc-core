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
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	ImageID    string        `gorm:"column:image_id;primaryKey" json:"image_id"`
	ExternalID string        `gorm:"column:external_id" json:"-"`
	MD5Sum     string        `gorm:"column:md5_sum;unique" json:"md5_sum"` // for shrinking file duplication
	TagIds     pq.Int32Array `gorm:"column:tag_ids;type:integer[]" json:"tag_ids"`
	Limit      int           `gorm:"-"`

	Tags []Tag `gorm:"-"`
}

func (*RichImage) TableName() string {
	return "rich_image"
}

func (i *RichImage) getRichImages(db *gorm.DB, excludeTagIds []int32) (images []RichImage, err error) {
	if i == nil {
		return nil, ErrNilMethodReceiver
	}
	if i.UpdatedAt.Before(_minTimeInt64) || i.UpdatedAt.After(_maxTimeInt64) {
		i.UpdatedAt = time.Now()
	}

	if i.Limit <= 0 || i.Limit > queryMaxLimit {
		i.Limit = queryMaxLimit
	}

	db = db.Table(i.TableName()).Where("updated_at < ?", i.UpdatedAt)
	if i.TagIds != nil {
		db = db.Where("tag_ids @> ?", pq.Array(i.TagIds))
	}
	if excludeTagIds != nil {
		db = db.Not("tag_ids && ?", pq.Array(excludeTagIds))
	}
	db = db.Order("updated_at desc").Limit(i.Limit)
	err = db.Scan(&images).Error

	if err != nil {
		return
	}

	return
}
