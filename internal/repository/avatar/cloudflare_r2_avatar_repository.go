package repository

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

const (
	ERR_STORAGE_INTERNAL_ERROR = "storage internal error"
)

type CloudflareR2AvatarRepository struct {
	s3_public_url string
	s3_endpoint   string
	s3_bucket     string
	s3_access_key string
	s3_secret_key string
	s3_account_id string
	client        *s3.Client
	manager       *manager.Uploader
}

func NewCloudflareR2AvatarRepository(s3_public_url, s3_bucket, s3_access_key, s3_secret_key, s3_account_id string) (*CloudflareR2AvatarRepository, *response.Err) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(s3_access_key, s3_secret_key, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil, response.NewErr(http.StatusInternalServerError, ERR_STORAGE_INTERNAL_ERROR)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", s3_account_id))
	})

	return &CloudflareR2AvatarRepository{
		s3_public_url: s3_public_url,
		s3_account_id: s3_account_id,
		s3_endpoint:   fmt.Sprintf("https://%s.r2.cloudflarestorage.com", s3_account_id),
		s3_bucket:     s3_bucket,
		s3_access_key: s3_access_key,
		s3_secret_key: s3_secret_key,
		client:        client,
		manager:       manager.NewUploader(client),
	}, nil
}

func (r2 *CloudflareR2AvatarRepository) SaveAvatar(ctx context.Context, avatar multipart.File) (string, *response.Err) {
	id := ctx.Value("user_id").(string)
	user_id, err := uuid.Parse(id)
	if err != nil {
		return "", response.NewErr(http.StatusBadRequest, user.ERR_UUID_INVALID)
	}

	file_key := fmt.Sprintf("avatars/%s", user_id)

	out, err := r2.manager.Upload(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r2.s3_bucket),
		Key:         aws.String(file_key),
		ContentType: aws.String("image/png"),
		Body:        avatar,
	})
	if err != nil {
		return "", response.NewErr(http.StatusInternalServerError, ERR_STORAGE_INTERNAL_ERROR)
	}

	avatar_url := fmt.Sprintf("%s/%s", r2.s3_public_url, *out.Key)

	return avatar_url, nil
}
