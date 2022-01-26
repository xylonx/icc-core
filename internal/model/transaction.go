package model

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/xylonx/icc-core/internal/config"
	"github.com/xylonx/icc-core/internal/core"
)

// map MIME type to file extension
var supportedMimeType = map[string]string{
	"image/png":  "png",
	"image/jpeg": "jpeg",
}

func AddAuthToken(ctx context.Context, token string) error {
	db := core.DB.WithContext(ctx)
	t := &AuthToken{Token: token}
	return t.insertToken(db)
}

func CheckTokenExists(ctx context.Context, token string) error {
	db := core.DB.WithContext(ctx)
	t := &AuthToken{Token: token}
	return t.tokenExists(db)
}

func AddUploadingBytes(ctx context.Context, token string, bytes int64) error {
	db := core.DB.WithContext(ctx)
	t := &AuthToken{Token: token, UploadingBytes: bytes}
	return t.addUploadingBytes(db)
}

func CheckImageExists(ctx context.Context, md5 string) error {
	db := core.DB.WithContext(ctx)
	i := &Image{MD5Sum: md5}
	return i.checkMD5(db)
}

func GetRichImages(ctx context.Context, before time.Time, tags []string, limit int) ([]RichImage, error) {
	db := core.DB.WithContext(ctx)
	i := &RichImage{UpdatedAt: before, Limit: limit, Tags: tags}
	return i.getRichImages(db)
}

func GetRandomImages(ctx context.Context, tags []string) (RichImage, error) {
	db := core.DB.WithContext(ctx)
	i := &RichImage{Tags: tags}
	return i.getRandom(db)
}

func GetALlTags(ctx context.Context) ([]Tag, error) {
	db := core.DB.WithContext(ctx)
	return getAllTags(db)
}

func GeneratePresignedUploadURL(ctx context.Context, mime, md5 string) (imageID, uri string, err error) {
	imageExt, ok := supportedMimeType[mime]
	if !ok {
		return "", "", ErrUnsupportedImageType
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return "", "", err
	}
	imageID = id.String() + "." + imageExt

	req, err := core.S3Client.PreSignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:   &core.S3Client.Bucket,
		Key:      &imageID,
		Metadata: map[string]string{"x-icc-md5": md5},
	})
	if err != nil {
		return "", "", err
	}

	return imageID, req.URL, nil
}

func DeleteTagsForImage(ctx context.Context, imageID string, tags []string) error {
	db := core.DB.WithContext(ctx)
	return deleteTagsForImage(db, imageID, tags)
}

/*
	below are complex sql needing trancation supports.
*/

func InsertRichImage(ctx context.Context, token, imageID, md5 string, tags []string, imgBytes int64) error {
	db := core.DB.WithContext(ctx)
	i := &Image{MD5Sum: md5, ImageID: imageID, Owner: token}
	t := &AuthToken{Token: token, UploadingBytes: imgBytes}
	if tags == nil {
		if err := i.insertImage(db); err != nil {
			return err
		}
		return nil
	}

	tx := db.Begin()
	if err := i.insertImage(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := insertTags(tx, tags); err != nil {
		tx.Rollback()
		return err
	}
	if err := addTagsForImage(tx, i.ImageID, tags); err != nil {
		tx.Rollback()
		return err
	}
	if err := t.addUploadingBytes(tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func AddTags(ctx context.Context, imageID string, tags []string) error {
	tx := core.DB.WithContext(ctx).Begin()
	if err := insertTags(tx, tags); err != nil {
		tx.Rollback()
		return err
	}
	if err := addTagsForImage(tx, imageID, tags); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func DeleteRichImage(ctx context.Context, imageID, token string) error {
	db := core.DB.WithContext(ctx)

	resp, err := core.S3Client.S3Client.HeadObject(ctx, &s3.HeadObjectInput{
		Bucket: &core.S3Client.Bucket,
		Key:    &imageID,
	})
	if err != nil {
		return err
	}

	i := &Image{ImageID: imageID}
	if token != config.Config.Application.AdminToken {
		i.Owner = token
	}

	t := &AuthToken{Token: token, UploadingBytes: resp.ContentLength}

	tx := db.Begin()
	if err := i.deleteImage(tx); err != nil {
		tx.Rollback()
		return err
	}
	if err := deleteImageAllTags(tx, imageID); err != nil {
		tx.Rollback()
		return err
	}
	if err := t.shrinkUploadingBytes(tx); err != nil {
		tx.Rollback()
		return err
	}

	_, err = core.S3Client.S3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &core.S3Client.Bucket,
		Key:    &imageID,
	})
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}
