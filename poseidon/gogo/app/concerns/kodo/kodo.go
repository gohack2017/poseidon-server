package kodo

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

const (
	MinKeys = 10
	MaxKeys = 20

	DatetimeLayout = "2006-01-02T15-04-05"
)

type Kodo struct {
	accessKey string
	secretKey string
	bucket    string
	config    *storage.Config
}

func New(config *Config) *Kodo {
	return &Kodo{
		accessKey: config.AccessKey,
		secretKey: config.SecretKey,
		bucket:    config.Bucket,
		config:    config.Kodo,
	}
}

func (kodo *Kodo) ResourceURL(key string) string {
	qmac := qbox.NewMac(kodo.accessKey, kodo.secretKey)

	rsurl := kodo.config.Zone.RsHost
	if !strings.HasPrefix(rsurl, "http://") && !strings.HasPrefix(rsurl, "https://") {
		rsurl = "http://" + rsurl
	}

	rsurl = strings.TrimSuffix(rsurl, "/")
	rsurl += "/get/" + storage.EncodedEntry(kodo.bucket, key)
	rsurl += "/expires/" + strconv.FormatInt(time.Now().Add(60*time.Second).Unix(), 10)

	client := storage.NewClient(qmac, nil)

	var result map[string]interface{}

	err := client.Call(context.Background(), &result, "POST", rsurl)
	if err != nil {
		return ""
	}

	return result["url"].(string)
}

func (kodo *Kodo) PutWithUser(user *User, data []byte) (err error) {
	key := user.KodoKey()
	encdata := user.Encode(data)

	return kodo.PutWithKey(key, []byte(encdata))
}

func (kodo *Kodo) PutWithKey(key string, data []byte) (err error) {
	qmac := qbox.NewMac(kodo.accessKey, kodo.secretKey)
	client := storage.NewFormUploader(kodo.config)
	policy := storage.PutPolicy{
		Scope:      kodo.bucket + ":" + key,
		InsertOnly: 0,
	}

	var result storage.ListItem

	return client.Put(
		context.Background(),
		&result,
		policy.UploadToken(qmac),
		key,
		bytes.NewBuffer(data),
		int64(len(data)),
		&storage.PutExtra{
			MimeType: "application/vnd.qiniu-facecloud.photo",
		},
	)
}

func (kodo *Kodo) ReadWithUser(user *User) (data []byte, err error) {
	data, err = kodo.ReadWithKey(user.KodoKey())
	if err != nil {
		return
	}

	return user.Decode(string(data))
}

func (kodo *Kodo) ReadWithKey(key string) (data []byte, err error) {
	request, err := http.NewRequest("GET", kodo.ResourceURL(key), nil)
	if err != nil {
		return
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	// spec for kodo none 2xx response
	if response.StatusCode/100 != 2 {
		err = errors.New(string(data))
		data = nil
		return
	}

	return
}

func (kodo *Kodo) HeadWithUser(user *User) (err error) {
	qmac := qbox.NewMac(kodo.accessKey, kodo.secretKey)
	client := storage.NewBucketManager(qmac, kodo.config)

	_, err = client.Stat(kodo.bucket, user.KodoKey())
	return
}

func (kodo *Kodo) ListWithUser(user *User, marker string, maxKeys int) (result *ListObjectsOutput, err error) {
	qmac := qbox.NewMac(kodo.accessKey, kodo.secretKey)
	client := storage.NewBucketManager(qmac, kodo.config)

	// adjust max keys
	if maxKeys <= 0 || maxKeys > MaxKeys {
		maxKeys = MaxKeys
	}

	entries, _, nextMarker, hasMore, err := client.ListFiles(kodo.bucket, user.KodoKey()+"/", "/", marker, maxKeys)
	if err != nil {
		return
	}

	result = NewListObjectsOutput(entries, marker, nextMarker, maxKeys, hasMore)
	return
}

func (kodo *Kodo) ListWithBucket(bucket, marker string, maxKeys int) (result *ListObjectsOutput, err error) {
	qmac := qbox.NewMac(kodo.accessKey, kodo.secretKey)
	client := storage.NewBucketManager(qmac, kodo.config)

	// adjust max keys
	if maxKeys <= 0 || maxKeys > MaxKeys {
		maxKeys = MaxKeys
	}

	entries, _, nextMarker, hasMore, err := client.ListFiles(bucket, "", "", marker, maxKeys)
	if err != nil {
		return
	}

	result = NewListObjectsOutput(entries, marker, nextMarker, maxKeys, hasMore)
	return
}

func (kodo *Kodo) Put(key string, data []byte) (err error) {
	qmac := qbox.NewMac(kodo.accessKey, kodo.secretKey)
	client := storage.NewFormUploader(kodo.config)
	policy := storage.PutPolicy{
		Scope:      kodo.bucket + ":" + key,
		InsertOnly: 0,
	}

	var result storage.ListItem

	return client.Put(
		context.Background(),
		&result,
		policy.UploadToken(qmac),
		key,
		bytes.NewBuffer(data),
		int64(len(data)),
		nil,
	)
}
