package db

import (
	"errors"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/hints"
)

var (
	ErrorNilOption = errors.New("option is nil")
)

type Option struct {
	DSN         string
	MaxOpenConn int
	MaxIdleConn int
	MaxLifetime time.Duration
}

func NewPostgreConn(opt *Option) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(opt.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	// limit connection limit
	sqlDb.SetMaxOpenConns(opt.MaxOpenConn)
	sqlDb.SetMaxIdleConns(opt.MaxIdleConn)
	sqlDb.SetConnMaxIdleTime(opt.MaxLifetime)

	db = db.Clauses(hints.New(" MAX_EXECUTION_TIME(1000) "))

	return db, nil
}
