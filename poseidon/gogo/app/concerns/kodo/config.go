package kodo

import (
	"net/url"

	"github.com/qiniu/api.v7/storage"
)

type Config struct {
	AccessKey    string          `json:"access_key"`
	SecretKey    string          `json:"secret_key"`
	Bucket       string          `json:"bucket"`
	BucketDomain string          `json:"bucket_domain"`
	Kodo         *storage.Config `json:"kodo"`
}

func (config *Config) CreatePublicUrl(key string) string {
	return config.BucketDomain + "/" + url.PathEscape(key)
}
