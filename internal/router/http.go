package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpOption struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func InitHttpServer(o *HttpOption) *http.Server {
	r := gin.New()

	r.Use()

	return &http.Server{
		Addr:         o.Addr,
		Handler:      r,
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,
	}
}
