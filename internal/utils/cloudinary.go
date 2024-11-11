package utils

import (
	"context"

	"github.com/adeyemialameen04/unwind-be/internal/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func NewCloudinaryInstance(cfg *config.Config) (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromURL(cfg.CoudinaryURL)
	if err != nil {
		return nil, err
	}

	return cld, err
}

func UploadImage(cld *cloudinary.Cloudinary, base64Str string, imageName string) (string, error) {
	ctx := context.Background()
	resp, err := cld.Upload.Upload(ctx, base64Str, uploader.UploadParams{
		PublicID:       imageName,
		UniqueFilename: api.Bool(true),
		Overwrite:      api.Bool(true),
	})
	if err != nil {
		return "", nil
	}

	return resp.SecureURL, nil
}