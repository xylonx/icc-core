package handler

import (
	"github.com/xylonx/icc-core/internal/proto/icc"
)

type ICCServer struct {
	icc.UnimplementedICCServer
}

var _ icc.ICCServer = &ICCServer{}

func NewICCServer() icc.ICCServer {
	return &ICCServer{}
}

// func (s *ICCServer) IssuePreSignUpload(context.Context, *icc.PreSignObjectRequest) (*icc.PreSignObjectResponse, error) {
// 	return nil, errors.New("not implemented")
// }

// func (s *ICCServer) IssuePreSignDownload(context.Context, *icc.PreSignObjectRequest) (*icc.PreSignObjectResponse, error) {
// 	return nil, errors.New("not implemented")
// }

// func (s *ICCServer) UpsertImageWithTags(context.Context, *icc.UpsertImageRequest) (*icc.UpsertImageResponse, error) {
// 	return nil, errors.New("not implemented")
// }
