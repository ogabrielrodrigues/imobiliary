package types

type Environment struct {
	SERVER_ADDR      string
	CORS_ORIGIN      string
	JWT_SECRET       string
	S3_PUBLIC_URL    string
	S3_AVATAR_BUCKET string
	S3_TOKEN         string
	S3_ACCESS_KEY    string
	S3_SECRET_KEY    string
	S3_ACCOUNT_ID    string
	DATABASE_HOST    string
	DATABASE_PORT    string
	DATABASE_NAME    string
	DATABASE_USER    string
	DATABASE_PWD     string
}
