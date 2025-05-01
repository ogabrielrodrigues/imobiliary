package provider

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

func (r2 *CloudflareR2AvatarRepository) ChangeAvatar(ctx context.Context, avatar multipart.File, mime string) (string, *response.Err) {
	id := ctx.Value(middleware.UserIDKey).(string)
	user_id, err := uuid.Parse(id)
	if err != nil {
		return "", response.NewErr(http.StatusBadRequest, user.ERR_UUID_INVALID)
	}

	file_key := fmt.Sprintf("avatars/%s", user_id)

	r2.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(r2.s3_bucket),
		Key:         aws.String(file_key),
		ContentType: aws.String(mime),
		Body:        avatar,
	})

	avatar_url := fmt.Sprintf("%s/%s", r2.s3_public_url, file_key)

	return avatar_url, nil
}
