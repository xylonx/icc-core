package s3

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	ErrorNilOption = errors.New("the option is not set")
)

type Client struct {
	S3Client      *s3.Client
	PreSignClient *s3.PresignClient
	Bucket        string
}

type Option struct {
	Endpoint       string
	AccessID       string
	AccessSecret   string
	BucketName     string
	Region         string
	PreSignExpires time.Duration
}

func NewS3Client(opt *Option) (*Client, error) {
	if opt == nil {
		return nil, ErrorNilOption
	}

	ep := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{URL: opt.Endpoint, SigningRegion: region}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(opt.AccessID, opt.AccessSecret, "")),
		config.WithRegion(opt.Region),
		config.WithEndpointResolverWithOptions(ep),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	presc := s3.NewPresignClient(client, func(po *s3.PresignOptions) {
		po.Expires = opt.PreSignExpires
	})

	return &Client{S3Client: client, PreSignClient: presc, Bucket: opt.BucketName}, nil
}
