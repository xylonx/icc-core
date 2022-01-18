package router

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xylonx/icc-core/internal/handler"
)

type HttpOption struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	AllowOrigins []string
}

func InitHttpServer(o *HttpOption) *http.Server {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// config cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = o.AllowOrigins
	r.Use(cors.New(corsConfig))

	// register routers
	subrouter := r.Group("/api/v1")
	subrouter.GET("/images", handler.GetImagesHandler)

	auth := subrouter.Group("/auth")
	auth.Use(handler.HttpAuthMiddleware())

	auth.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	auth.POST("/image/complete", handler.AddImageHandler)
	auth.POST("/image/upload", handler.GeneratePreSignUpload)
	auth.PUT("/image/tag", handler.UpsertImageTag)

	auth.POST("/token", handler.GenereateToken)

	return &http.Server{
		Addr:         o.Addr,
		Handler:      r,
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,
	}
}
