package core

import (
	"strconv"
	"time"

	sess_redis "github.com/gin-contrib/sessions/redis"
	"github.com/go-redis/redis/v8"
	"github.com/xylonx/icc-core/internal/config"
	"github.com/xylonx/icc-core/pkg/db"
	"github.com/xylonx/icc-core/pkg/s3"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DB                *gorm.DB
	S3Client          *s3.Client
	RedisSessionStore sess_redis.Store
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
		PreSignExpires: time.Duration(config.Config.Storage.SignUploadSeconds) * time.Second,
	})
	if err != nil {
		zapx.Error("new s3 client failed", zap.Error(err))
		return err
	}

	var redisOpt *redis.Options
	redisOpt, err = redis.ParseURL(config.Config.Database.Redis.Dsn)
	if err != nil {
		zapx.Error("parse redis dsn failed", zap.Error(err))
		return err
	}
	RedisSessionStore, err = sess_redis.NewStoreWithDB(
		10, redisOpt.Network, redisOpt.Addr, redisOpt.Password,
		strconv.FormatInt(int64(redisOpt.DB), 10), []byte(config.Config.Application.HttpSessionSecret),
	)
	if err != nil {
		zapx.Error("create redis session store failed", zap.Error(err))
		return err
	}

	return
}
