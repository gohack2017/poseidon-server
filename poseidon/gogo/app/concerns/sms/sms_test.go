package sms

import (
	"fmt"
	"testing"

	"github.com/golib/assert"
)

func Test_ResourceAPI(t *testing.T) {
	assetion := assert.New(t)

	cfg := &Config{
		Endpoint: "http://106.ihuyi.com",
		Account:  "",
		Secret:   "",
		Template: "您的验证码是：%s。请不要把验证码泄露给其他人。",
	}

	smsClient := New(cfg)

	out, err := smsClient.Send("13301631167", cfg.Render("121212"))
	assetion.Nil(err)
	fmt.Println(out)
}
