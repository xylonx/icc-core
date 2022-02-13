package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"
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

type GetImagesRequest struct {
	Before time.Time `json:"before" form:"before" time_format:"unix"`
	Tag    string    `json:"tag" form:"tag"`
	Limit  uint      `json:"limit" form:"limit"`
}

type GetImageResponse struct {
	ImageURL  string   `json:"image_url"`
	ImageID   string   `json:"image_id"`
	Tags      []string `json:"tags"`
	Timestamp int64    `json:"timestamp"`
}

type AddImageRequest struct {
	ImageID    string   `json:"image_id"`
	ExternalID string   `json:"external_id"`
	Tags       []string `json:"tags"`
}

type GeneratePreSignRequest struct {
	ExternalID string `json:"external_id,omitempty"`
	ImageType  string `json:"image_type"`
	MD5Sum     string `json:"md5_sum"`
}

type ImageTagRequest struct {
	ImageID string   `json:"image_id"`
	Tags    []string `json:"tags"`
}

func HttpAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			zapx.Error("auth header missing", zap.String("Authorization", authHeader))
			respAbortWithUnauthError(c, ErrAuthHeaderNotFound)
			return
		}

		token := authHeader[7:]

		if err := model.CheckTokenExists(c.Request.Context(), token); err != nil {
			zapx.Error("check token exists failed", zap.Error(err), zap.String("token", token))
			respAbortWithUnauthError(c, err)
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

	var tags []string
	if req.Tag != "" {
		tags = strings.Split(req.Tag, ",")
	}

	images, err := model.GetRichImages(ctx.Request.Context(), req.Before, tags, int(req.Limit))
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
		resp[i].Tags = images[i].Tags
		resp[i].Timestamp = images[i].UpdatedAt.Unix()
	}

	respOK(ctx, resp)
}

func GetRandomImageHandler(ctx *gin.Context) {
	var req GetImagesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		respParamBindingError(ctx, err)
		return
	}

	var tags []string
	if req.Tag != "" {
		tags = strings.Split(req.Tag, ",")
	}

	image, err := model.GetRandomImages(ctx.Request.Context(), tags, 1)
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

	token := ctx.Request.Context().Value("ICCAuthToken").(string) //nolint:forcetypeassert
	md5 := resp.Metadata["x-icc-md5"]

	err = model.InsertRichImage(ctx.Request.Context(), token, req.ImageID, md5, req.Tags, resp.ContentLength)
	if err != nil {
		respDBError(ctx, err)
		return
	}

	respOK(ctx, struct {
		ImageID  string   `json:"image_id"`
		ImageURL string   `json:"image_url"`
		Tags     []string `json:"tags"`
	}{
		ImageID:  req.ImageID,
		ImageURL: core.S3Client.ConstructDownloadURL(ctx.Request.Context(), req.ImageID),
		Tags:     req.Tags,
	})
}

func DeleteRichImageHandler(ctx *gin.Context) {
	imageID := ctx.Param("id")

	token := ctx.Request.Context().Value("ICCAuthToken").(string) //nolint:forcetypeassert
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
	tags, err := model.GetALlTags(ctx.Request.Context())
	if err != nil {
		zapx.Error("get all tags failed", zap.Error(err))
		respDBError(ctx, err)
		return
	}

	respOK(ctx, tags)
}

func AddTagToImage(ctx *gin.Context) {
	var req ImageTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		respParamBindingError(ctx, err)
		return
	}

	ri := model.RichImage{ImageID: req.ImageID, Tags: req.Tags}
	if err := model.AddTags(ctx.Request.Context(), req.ImageID, req.Tags); err != nil {
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

	ri := model.RichImage{ImageID: req.ImageID, Tags: req.Tags}
	if err := model.DeleteTagsForImage(ctx.Request.Context(), req.ImageID, req.Tags); err != nil {
		zapx.Error("delete tags to image failed", zap.Error(err))
		respDBError(ctx, err)
		return
	}
	respOK(ctx, ri)
}

func GenereateToken(ctx *gin.Context) {
	authToken := ctx.Request.Context().Value("ICCAuthToken").(string) // nolint:forcetypeassert
	if authToken != config.Config.Application.AdminToken {
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
