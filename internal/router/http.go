package router

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/xylonx/icc-core/internal/core"
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

	adminrouter := subrouter.Group("/admin")
	adminrouter.Use(sessions.Sessions("icc", core.RedisSessionStore))
	adminrouter.Use(sessionMiddleware)
	adminrouter.POST("/image", handler.AddImageHandler)
	adminrouter.POST("/image/upload", handler.GeneratePreSignUpload)
	adminrouter.PUT("/image/tag", handler.UpsertImageTag)

	return &http.Server{
		Addr:         o.Addr,
		Handler:      r,
		ReadTimeout:  o.ReadTimeout,
		WriteTimeout: o.WriteTimeout,
	}
}

func sessionMiddleware(*gin.Context) {
	// sess := sessions.Default(ctx)
	// if sess.Get("user") == nil {
	// 	ctx.Redirect(http.StatusSeeOther, config.Config.Application.HttpSessionRedirectPage)
	// 	ctx.Abort()
	// 	return
	// }
}
