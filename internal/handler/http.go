package handler

import (
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xylonx/icc-core/internal/config"
	"github.com/xylonx/icc-core/internal/model"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
)

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

func GetImagesHandler(ctx *gin.Context) {
	var req GetImagesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		respParamBindingError(ctx, err)
		return
	}

	req.Tags = strings.Split(req.Tag, ",")

	zapx.Info("get request param", zap.Any("req", req))

	ri := model.RichImage{
		UpdatedAt: req.Before,
		Tags:      req.Tags,
	}

	images, err := ri.GetRichImages(int(req.Limit))
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

	respOK(ctx, images)
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

	if err := ri.InsertImages(); err != nil {
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

	respOK(ctx, uri)
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

	if err := ri.UpsertTags(); err != nil {
		respDBError(ctx, err)
		return
	}

	respOK(ctx, ri)
}
