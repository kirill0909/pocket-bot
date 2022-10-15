package repository

const (
	AccessTokens  Bucket = "accessTokens"
	RequestTokens Bucket = "requestTokens"
)

type Bucket string

type TokenRepository interface {
	Save(chatID int64, token string, bucket Bucket) error
	Get(chatID int64, bucket Bucket) (string, error)
}
