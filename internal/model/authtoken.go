package model

import (
	"context"
	"errors"
	"time"

	"github.com/xylonx/icc-core/internal/core"
	"gorm.io/gorm"
)

var (
	ErrNilToken = errors.New("token is nil")
)

type AuthToken struct {
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Token string `gorm:"column:token;primaryKey"`
}

func (*AuthToken) TableName() string {
	return "auth_token"
}

func (t *AuthToken) InsertToken(ctx context.Context) error {
	if t == nil {
		return ErrNilMethodReceiver
	}

	if t.Token == "" {
		return ErrNilToken
	}

	db := core.DB.WithContext(ctx)

	if err := db.Create(t).Error; err != nil {
		return err
	}

	return nil
}

func (t *AuthToken) TokenExists(ctx context.Context) error {
	if t == nil {
		return ErrNilMethodReceiver
	}

	if t.Token == "" {
		return ErrNilToken
	}

	db := core.DB.WithContext(ctx)

	if err := db.First(t).Error; err != nil {
		return err
	}

	return nil
}
