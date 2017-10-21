package kodo

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/poseidon/lib/encoding"
)

// NOTE: DO NOT change codes here!!!
type User struct {
	sn  string
	key []byte
}

func NewUser(sn string) *User {
	return &User{
		sn:  sn,
		key: []byte("nX#WC$#BBe!u%49MhG"), // NOTE: NEVER change this!!!
	}
}

func (user *User) KodoKey() string {
	hash := hmac.New(md5.New, user.key)
	hash.Write([]byte(user.sn))

	return hex.EncodeToString(hash.Sum(nil))
}

func (user *User) KodoFaceKey() string {
	return user.KodoSubKey(time.Now().Format(DatetimeLayout))
}

func (user *User) KodoSubKey(key string) string {
	return user.KodoKey() + "/" + key
}

func (user *User) Encode(data []byte) string {
	encoding := encoding.New(user.sn, user.KodoKey())

	return encoding.Encrypt(data)
}

func (user *User) Decode(s string) ([]byte, error) {
	encoding := encoding.New(user.sn, user.KodoKey())

	return encoding.Decrypt(s)
}
