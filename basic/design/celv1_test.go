package design

import (
	"context"
	"io"
	"testing"
)

type ObjectStoreService interface {
	UploadFile(context.Context, string, io.Reader) error
	GetSignURL(context.Context, string, int64) error
}

func OssFactory(name string) ObjectStoreService {
	switch name {
	case "aliOss":
		return aliOSs{}
	case "txOss":
		return txOSs{}
	default:
		return nil
	}
}

type aliOSs struct {
}

func (ali aliOSs) UploadFile(ctx context.Context, path string, reader io.Reader) error {
	return nil
}

func (ali aliOSs) GetSignURL(ctx context.Context, path string, time int64) error {
	return nil
}

type txOSs struct {
}

func (tx txOSs) UploadFile(ctx context.Context, path string, reader io.Reader) error {
	return nil
}

func (tx txOSs) GetSignURL(ctx context.Context, path string, time int64) error {
	return nil
}

func TestQuestion2Method(t *testing.T) {
	NewMaiMai("你好").SendMsg()
}
