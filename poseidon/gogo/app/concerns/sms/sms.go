package sms

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	ServiceAPI = "/webservice/sms.php"
)

type SMS struct {
	endpoint string
	account  string
	secret   string
}

func New(config *Config) *SMS {
	return &SMS{
		endpoint: config.Endpoint,
		account:  config.Account,
		secret:   config.Secret,
	}
}

func (sms *SMS) ResourceAPI(params url.Values) string {
	apiurl := sms.endpoint
	if !strings.HasPrefix(apiurl, "http://") &&
		!strings.HasPrefix(apiurl, "https://") &&
		!strings.HasPrefix(apiurl, "mitm://") {
		apiurl = "https://" + apiurl
	}
	apiurl = strings.TrimSuffix(apiurl, "/")
	apiurl += ServiceAPI
	apiurl += "?" + params.Encode()

	return apiurl
}

func (sms *SMS) Send(phone, content string) (out *SMSOutput, err error) {
	params := url.Values{}
	params.Add("method", "Submit")

	payload := url.Values{}
	payload.Add("account", sms.account)
	payload.Add("password", sms.secret)
	payload.Add("mobile", phone)
	payload.Add("content", content)
	payload.Add("format", "json")

	fmt.Println(payload.Encode())

	request, err := http.NewRequest("POST", sms.ResourceAPI(params), strings.NewReader(payload.Encode()))
	if err != nil {
		return
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if response.StatusCode/100 != 2 {
		err = errors.New(string(data))
		return
	}

	err = json.Unmarshal(data, &out)
	if out.Code != 2 {
		err = errors.New(string(data))
	}

	return
}
