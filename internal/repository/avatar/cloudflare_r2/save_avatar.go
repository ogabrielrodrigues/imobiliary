package repository

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/ogabrielrodrigues/imobiliary/internal/entity/user"
	"github.com/ogabrielrodrigues/imobiliary/internal/middleware"
	"github.com/ogabrielrodrigues/imobiliary/internal/types/response"
)

func (r2 *CloudflareR2AvatarRepository) SaveAvatar(ctx context.Context, avatar multipart.File) (string, *response.Err) {
	id := ctx.Value(middleware.UserIDKey).(string)
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
