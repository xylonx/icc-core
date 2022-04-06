package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"github.com/xylonx/icc-core/internal/config"
	"github.com/xylonx/icc-core/internal/core"
	"github.com/xylonx/icc-core/internal/model"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
)

var (
	ErrAuthHeaderNotFound = errors.New("auth header not found")
	ErrPermissionDenied   = errors.New("permission denied")
)

type Tag struct {
	TagID     int32  `json:"tag_id"`
	TagNameEN string `json:"tag_name_en"`
	TagNameCN string `json:"tag_name_cn"`
	TagNameJP string `json:"tag_name_jp"`
}

func modelTag2Tag(ts []model.Tag) []Tag {
	tags := make([]Tag, 0, len(ts))
	for i := range ts {
		tags = append(tags, Tag{TagID: ts[i].ID, TagNameEN: ts[i].TagNameEN, TagNameCN: ts[i].TagNameCN, TagNameJP: ts[i].TagNameJP})
	}
	return tags
}

func tag2ModelTag(ts []Tag) []model.Tag {
	tags := make([]model.Tag, 0, len(ts))
	for i := range ts {
		tags = append(tags, model.Tag{ID: ts[i].TagID, TagNameEN: ts[i].TagNameEN, TagNameCN: ts[i].TagNameCN, TagNameJP: ts[i].TagNameJP})
	}
	return tags
}

type GetImagesRequest struct {
	Before        time.Time `json:"before" form:"before" time_format:"unix"`
	TagIds        []int32   `json:"tag_ids" form:"tag_ids"`
	ExcludeTagIds []int32   `json:"exclude_tag_ids" form:"exclude_tag_ids"`
	Limit         uint      `json:"limit" form:"limit"`
}

type GetImageResponse struct {
	ImageURL  string  `json:"image_url"`
	ImageID   string  `json:"image_id"`
	TagIds    []int32 `json:"tag_ids"`
	UpdatedAt int64   `json:"updated_at"`
}

type AddImageRequest struct {
	ImageID    string `json:"image_id"`
	ExternalID string `json:"external_id"`
	Tags       []Tag  `json:"tags"`
}

type GeneratePreSignRequest struct {
	ExternalID string `json:"external_id,omitempty"`
	ImageType  string `json:"image_type"`
	MD5Sum     string `json:"md5_sum"`
}

type ImageTagRequest struct {
	ImageID string `json:"image_id"`
	Tags    []Tag  `json:"tags"`
}

func GetImagesHandler(ctx *gin.Context) {
	var req GetImagesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		respParamBindingError(ctx, err)
		return
	}

	images, err := model.GetRichImages(ctx.Request.Context(), req.Before, req.TagIds, req.ExcludeTagIds, int(req.Limit))
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
		resp[i].ImageURL = core.S3Client.ConstructDownloadURL(ctx.Request.Context(), images[i].ImageID)
		resp[i].ImageID = images[i].ImageID
		resp[i].TagIds = images[i].TagIds
		resp[i].UpdatedAt = images[i].UpdatedAt.Unix()
	}

	respOK(ctx, resp)
}

func GetRandomImageHandler(ctx *gin.Context) {
	var req GetImagesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		respParamBindingError(ctx, err)
		return
	}

	image, err := model.GetRandomImages(ctx.Request.Context(), req.TagIds, 1)
	if err != nil {
		zapx.Error("get random images failed", zap.Error(err))
		respDBError(ctx, err)
		return
	}

	uri := core.S3Client.ConstructDownloadURL(ctx.Request.Context(), image[0].ImageID)
	ctx.Redirect(http.StatusFound, uri)
}

func AddImageHandler(ctx *gin.Context) {
	var req AddImageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respParamBindingError(ctx, err)
		return
	}

	resp, err := core.S3Client.S3Client.HeadObject(ctx.Request.Context(), &s3.HeadObjectInput{
		Bucket: &core.S3Client.Bucket,
		Key:    &req.ImageID,
	})
	if err != nil {
		zapx.Error("head image info failed", zap.Error(err), zap.String("imageID", req.ImageID))
		respS3Error(ctx, err)
		return
	}

	token := mustGetTokenFromCtx(ctx.Request.Context())
	md5 := resp.Metadata["x-icc-md5"]

	err = model.InsertRichImage(ctx.Request.Context(), token, req.ImageID, md5, tag2ModelTag(req.Tags), resp.ContentLength)
	if err != nil {
		respDBError(ctx, err)
		return
	}

	respOK(ctx, struct {
		ImageID  string `json:"image_id"`
		ImageURL string `json:"image_url"`
		Tags     []Tag  `json:"tags"`
	}{
		ImageID:  req.ImageID,
		ImageURL: core.S3Client.ConstructDownloadURL(ctx.Request.Context(), req.ImageID),
		Tags:     req.Tags,
	})
}

func DeleteRichImageHandler(ctx *gin.Context) {
	imageID := ctx.Param("id")

	token := mustGetTokenFromCtx(ctx.Request.Context())
	if err := model.DeleteRichImage(ctx.Request.Context(), imageID, token); err != nil {
		zapx.Error("delete rich image failed", zap.Error(err))
		respDBError(ctx, err)
		return
	}

	respOK(ctx, "ok")
}

func GeneratePreSignUpload(ctx *gin.Context) {
	var req GeneratePreSignRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		zapx.Error("bind json failed", zap.Error(err))
		respParamBindingError(ctx, err)
		return
	}

	if err := model.CheckImageExists(ctx.Request.Context(), req.MD5Sum); err != nil {
		zapx.Error("check image exists failed", zap.Error(err))
		respDBError(ctx, err)
		return
	}

	imageID, uri, err := model.GeneratePresignedUploadURL(ctx.Request.Context(), req.ImageType, req.MD5Sum)
	if err != nil {
		zapx.Error("generate pre signed uri failed", zap.Error(err))
		respDBError(ctx, err)
		return
	}

	respOK(ctx, struct {
		PresignedURI string `json:"presigned_uri"`
		ImageID      string `json:"image_id"`
	}{
		PresignedURI: uri,
		ImageID:      imageID,
	})
}

func GetAllTags(ctx *gin.Context) {
	tags, err := model.GetAllTags(ctx.Request.Context())
	if err != nil {
		zapx.Error("get all tags failed", zap.Error(err))
		respDBError(ctx, err)
		return
	}

	respOK(ctx, modelTag2Tag(tags))
}

func AddTagToImage(ctx *gin.Context) {
	var req ImageTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respParamBindingError(ctx, err)
		return
	}

	ri := model.RichImage{ImageID: req.ImageID}
	if err := model.AddTags(ctx.Request.Context(), req.ImageID, tag2ModelTag(req.Tags)); err != nil {
		zapx.Error("add tags to image failed", zap.Error(err))
		respDBError(ctx, err)
		return
	}

	respOK(ctx, ri)
}

func DeleteTagToImage(ctx *gin.Context) {
	var req ImageTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respParamBindingError(ctx, err)
		return
	}

	ri := model.RichImage{ImageID: req.ImageID}
	if err := model.DeleteTagsForImage(ctx.Request.Context(), req.ImageID, tag2ModelTag(req.Tags)); err != nil {
		zapx.Error("delete tags to image failed", zap.Error(err))
		respDBError(ctx, err)
		return
	}
	respOK(ctx, ri)
}

func GenereateToken(ctx *gin.Context) {
	if mustGetTokenFromCtx(ctx.Request.Context()) != config.Config.Application.AdminToken {
		respAbortWithUnauthError(ctx, ErrPermissionDenied)
		return
	}

	uid, err := ksuid.NewRandom()
	if err != nil {
		zapx.Error("generate ksuid failed", zap.Error(err))
		respUnknownError(ctx, err)
		return
	}
	id := uid.String()

	if err := model.AddAuthToken(ctx.Request.Context(), id); err != nil {
		zapx.Error("insert token failed", zap.Error(err))
		respDBError(ctx, err)
		return
	}

	respOK(ctx, id)
}

func I18nTagsHandler(ctx *gin.Context) {
	req := new(struct {
		Tags []Tag `json:"tags"`
	})
	if err := ctx.ShouldBindJSON(req); err != nil {
		zapx.Error("bind i18n tag failed", zap.Error(err))
		respParamBindingError(ctx, err)
		return
	}

	if err := model.UpdateTagI18n(ctx.Request.Context(), tag2ModelTag(req.Tags)); err != nil {
		zapx.Error("update tag i18n failed", zap.Error(err))
		respDBError(ctx, err)
		return
	}
	respOK(ctx, req.Tags)
}
