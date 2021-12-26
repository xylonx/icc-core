package model

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/xylonx/icc-core/internal/core"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	ErrNilMethodReceiver    = errors.New(" method receiver is nil")
	ErrNilImageID           = errors.New("image is nil")
	ErrNilTag               = errors.New("tag is nil")
	ErrUnsupportedImageType = errors.New("image type is not supported now")

	_minTimeInt64 = time.Unix(0, math.MinInt64)
	_maxTimeInt64 = time.Unix(0, math.MaxInt64)
)

var supportedImageExt = map[string]struct{}{
	"png": {},
	"jpg": {},
}

const queryMaxLimit = 100

type RichImage struct {
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	ImageID    string         `gorm:"column:image_id;primaryKey" json:"image_id"`
	ExternalID string         `gorm:"column:external_id;index" json:"-"`
	Tags       pq.StringArray `gorm:"column:tags;type:text[]" json:"tags"`
}

func (*RichImage) TableName() string {
	return "image"
}

func (i *RichImage) InsertImages() (err error) {
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
	if i.UpdatedAt.Before(_minTimeInt64) || i.UpdatedAt.After(_maxTimeInt64) {
		i.UpdatedAt = time.Now()
	}

	if limit <= 0 || limit > queryMaxLimit {
		limit = queryMaxLimit
	}

	if err = core.DB.Where("updated_at < ? AND tags @> ?", i.UpdatedAt, pq.Array(i.Tags)).
		Limit(limit).Find(&images).Statement.Error; err != nil {
		return
	}

	return
}

func (i *RichImage) GeneratePreSignUpload(ctx context.Context, imageExt string) (string, error) {
	if _, ok := supportedImageExt[imageExt]; !ok {
		return "", ErrUnsupportedImageType
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	i.ImageID = id.String() + "." + imageExt

	req, err := core.S3Client.PreSignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: &core.S3Client.Bucket,
		Key:    &i.ImageID,
	})
	if err != nil {
		return "", err
	}

	return req.URL, nil
}
