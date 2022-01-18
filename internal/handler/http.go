package handler

import (
	"context"
	"errors"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"github.com/xylonx/icc-core/internal/config"
	"github.com/xylonx/icc-core/internal/model"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
)

var (
	ErrAuthHeaderNotFound = errors.New("auth header not found")
	ErrPermissionDenied   = errors.New("permission denied")
)

var kuid *ksuid.KSUID

type GetImagesRequest struct {
	Before time.Time `json:"before" form:"before" time_format:"unix"`
	Tag    string    `json:"tag" form:"tag"`
	Tags   []string  `json:"-" form:"tags"` // equal strings.Split(tag, ',')
	Limit  uint      `json:"limit" form:"limit"`
}

type GetImageResponse struct {
	ImageURL string   `json:"image_url"`
	Tags     []string `json:"tags"`
}

type AddImageRequest struct {
	ImageID    string `json:"image_id"`
	ExternalID string `json:"external_id"`
}

type GeneratePreSignRequest struct {
	ExternalID string `json:"external_id,omitempty"`
	ImageType  string `json:"image_type"`
}

type UpsertImageTagRequest struct {
	ImageID string   `json:"image_id"`
	Tags    []string `json:"tags"`
}

func HttpAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			zapx.Error("auth header missing", zap.String("Authorization", authHeader))
			respAbortWithForbiddenError(c, ErrAuthHeaderNotFound)
			return
		}

		token := authHeader[7:]

		t := model.AuthToken{
			Token: token,
		}

		if err := t.TokenExists(c.Request.Context()); err != nil {
			zapx.Error("check token exists failed", zap.Error(err), zap.String("token", token))
			respAbortWithForbiddenError(c, err)
			return
		}

		ctx := context.WithValue(c.Request.Context(), "ICCAuthToken", token) // nolint:staticcheck
		c.Request = c.Request.Clone(ctx)
	}
}

func GetImagesHandler(ctx *gin.Context) {
	var req GetImagesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		respParamBindingError(ctx, err)
		return
	}

	if req.Tag == "" {
		req.Tags = nil
	} else {
		req.Tags = strings.Split(req.Tag, ",")
	}

	zapx.Info("get request param", zap.Any("req", req))

	ri := model.RichImage{
		UpdatedAt: req.Before,
		Tags:      req.Tags,
	}

	images, err := ri.GetRichImages(ctx.Request.Context(), int(req.Limit))
	if err != nil {
		respDBError(ctx, err)
		return
	}

	if len(images) == 0 {
		respNotFoundError(ctx)
		return
	}

	resp := make([]GetImageResponse, len(images))
	for i := range images {
		resp[i].ImageURL = path.Join(config.Config.Storage.CdnHost, config.Config.Storage.S3.Bucket, images[i].ImageID)
		resp[i].Tags = images[i].Tags
	}

	respOK(ctx, resp)
}

func AddImageHandler(ctx *gin.Context) {
	var req AddImageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respParamBindingError(ctx, err)
		return
	}

	ri := model.RichImage{
		ImageID:    req.ImageID,
		ExternalID: req.ExternalID,
	}

	if err := ri.InsertImages(ctx.Request.Context()); err != nil {
		respDBError(ctx, err)
		return
	}

	respOK(ctx, ri)
}

func GeneratePreSignUpload(ctx *gin.Context) {
	var req GeneratePreSignRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respParamBindingError(ctx, err)
		return
	}

	ri := model.RichImage{
		ExternalID: req.ExternalID,
	}

	uri, err := ri.GeneratePreSignUpload(ctx.Request.Context(), req.ImageType)
	if err != nil {
		respDBError(ctx, err)
		return
	}

	respOK(ctx, struct {
		PresignedURI string `json:"presigned_uri"`
		ImageID      string `json:"image_id"`
	}{
		PresignedURI: uri,
		ImageID:      ri.ImageID,
	})
}

func UpsertImageTag(ctx *gin.Context) {
	var req UpsertImageTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respParamBindingError(ctx, err)
		return
	}

	ri := model.RichImage{
		ImageID: req.ImageID,
		Tags:    req.Tags,
	}

	if err := ri.UpsertTags(ctx.Request.Context()); err != nil {
		respDBError(ctx, err)
		return
	}

	respOK(ctx, ri)
}

func GenereateToken(ctx *gin.Context) {
	authToken := ctx.Request.Context().Value("ICCAuthToken").(string) // nolint:forcetypeassert
	if authToken != config.Config.Application.AdminToken {
		respAbortWithForbiddenError(ctx, ErrPermissionDenied)
		return
	}

	if kuid == nil {
		uid, err := ksuid.NewRandom()
		if err != nil {
			zapx.Error("generate ksuid failed", zap.Error(err))
			respUnknownError(ctx, err)
			return
		}

		kuid = &uid
	}

	id := kuid.String()
	nxtId := kuid.Next()
	kuid = &nxtId

	t := model.AuthToken{Token: id}
	if err := t.InsertToken(ctx.Request.Context()); err != nil {
		zapx.Error("insert token failed", zap.Error(err))
		respDBError(ctx, err)
		return
	}

	respOK(ctx, id)
}
