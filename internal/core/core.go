package core

import (
	"time"

	"github.com/xylonx/icc-core/internal/config"
	"github.com/xylonx/icc-core/pkg/db"
	"github.com/xylonx/icc-core/pkg/s3"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	S3Client *s3.Client
)

func Setup() (err error) {
	DB, err = db.NewPostgreConn(&db.Option{
		DSN:         config.Config.Database.Postgres.Dsn,
		MaxOpenConn: int(config.Config.Database.Postgres.MaxOpenConn),
		MaxIdleConn: int(config.Config.Database.Postgres.MaxIdleConn),
		MaxLifetime: time.Duration(config.Config.Database.Postgres.MaxLifeSeconds) * time.Second,
	})
	if err != nil {
		zapx.Error("new postgres connection failed", zap.Error(err))
		return err
	}

	S3Client, err = s3.NewS3Client(&s3.Option{
		Endpoint:       config.Config.Storage.S3.Endpoint,
		AccessID:       config.Config.Storage.S3.AccessId,
		AccessSecret:   config.Config.Storage.S3.AccessSecret,
		BucketName:     config.Config.Storage.S3.Bucket,
		Region:         config.Config.Storage.S3.Region,
		CDNHost:        config.Config.Storage.CdnHost,
		PreSignExpires: time.Duration(config.Config.Storage.SignUploadSeconds) * time.Second,
	})
	if err != nil {
		zapx.Error("new s3 client failed", zap.Error(err))
		return err
	}

	return
}
