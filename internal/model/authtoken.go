package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrNilToken         = errors.New("token is nil")
	ErrNilUploadingSize = errors.New("uploading size is 0")
)

type AuthToken struct {
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Token          string `gorm:"column:token;primaryKey"`
	UploadingBytes int64  `gorm:"column:uploading_bytes"`
}

func (*AuthToken) TableName() string {
	return "auth_token"
}

func (t *AuthToken) insertToken(db *gorm.DB) error {
	if t == nil {
		return ErrNilMethodReceiver
	}

	if t.Token == "" {
		return ErrNilToken
	}

	if err := db.Create(t).Error; err != nil {
		return err
	}

	return nil
}

func (t *AuthToken) tokenExists(db *gorm.DB) error {
	if t == nil {
		return ErrNilMethodReceiver
	}

	if t.Token == "" {
		return ErrNilToken
	}

	if err := db.First(t).Error; err != nil {
		return err
	}

	return nil
}

func (t *AuthToken) addUploadingBytes(db *gorm.DB) error {
	if t == nil {
		return ErrNilMethodReceiver
	}

	if t.Token == "" {
		return ErrNilToken
	}

	if t.UploadingBytes == 0 {
		return ErrNilUploadingSize
	}

	if err := db.Model(t).Where("token = ?", t.Token).
		Update("uploading_bytes", gorm.Expr("uploading_bytes + ?", t.UploadingBytes)).Error; err != nil {
		return err
	}

	return nil
}

func (t *AuthToken) shrinkUploadingBytes(db *gorm.DB) error {
	if t == nil {
		return ErrNilMethodReceiver
	}

	if t.Token == "" {
		return ErrNilToken
	}

	if t.UploadingBytes == 0 {
		return ErrNilUploadingSize
	}

	if err := db.Model(t).Where("token = ?", t.Token).
		Update("uploading_bytes", gorm.Expr("uploading_bytes - ?", t.UploadingBytes)).Error; err != nil {
		return err
	}

	return nil
}
