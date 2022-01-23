package model

import (
	"errors"
	"math"
	"math/rand"
	"time"

	"github.com/lib/pq"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
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

// TODO: select a random row by using an auto-increment key now.
// it is not efficient but suit for small dataset.
// for big dataset, using tablesample like below for efficiency
//
// 1) enable tsr extension: `CREATE EXTENSION tsm_system_rows`
// 2) make rich_image as a meterialed view
// 3) replace below random select to this: `select * from rich_image tablesample system_rows(1);`
func (i *RichImage) getRandom(db *gorm.DB) (image RichImage, err error) {
	if i == nil {
		return image, ErrNilMethodReceiver
	}

	var maxSeq int64
	err = db.Raw("select last_value from image_seq_id_seq").Scan(&maxSeq).Error
	if err != nil {
		return
	}

	seq := rand.Int63n(maxSeq)
	zapx.Info("current seq", zap.Int64("seq", seq))

	if i.Tags == nil {
		err = db.Table(i.TableName()).Where("seq_id = ?", seq).Scan(&image).Error
	} else {
		err = db.Table(i.TableName()).Where("seq_id = ? AND tags @> ?", seq, pq.Array(i.Tags)).Scan(&image).Error
	}
	if err != nil {
		return
	}

	return
}
