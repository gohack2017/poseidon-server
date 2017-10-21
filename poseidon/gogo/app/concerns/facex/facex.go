package facex

import (
	"fmt"
	"strings"
	"time"

	"github.com/poseidon/lib/request"
)

type Facex struct {
	endpoint  string
	accessKey string
	secretKey string
	groupId   string
	timeout   time.Duration
}

func NewFacex(config *Config) *Facex {
	return &Facex{
		endpoint:  config.Endpoint,
		accessKey: config.AccessKey,
		secretKey: config.SecretKey,
		groupId:   config.GroupId,
		timeout:   time.Duration(config.Timeout) * time.Second,
	}
}

const (
	GroupNewAPI    = "/v1/face/group/%s/new"
	GroupAddAPI    = "/v1/face/group/%s/add"
	GroupSearchAPI = "/v1/face/group/%s/search"
)

func (facex *Facex) API(api string) string {
	return fmt.Sprintf(strings.TrimSuffix(facex.endpoint, "/")+api, facex.groupId)
}

func (facex *Facex) NewGroup(input FacexInput) (err error) {
	client := request.New(request.NewAccessKey(facex.accessKey, facex.secretKey))

	_, err = client.Send("POST", facex.API(GroupNewAPI), "application/json", facex.timeout, toPayload(input))
	return
}

func (facex *Facex) AddFace(uri, id string) (err error) {
	client := request.New(request.NewAccessKey(facex.accessKey, facex.secretKey))

	input := NewFacexInput(uri, id)
	_, err = client.Send("POST", facex.API(GroupAddAPI), "application/json", facex.timeout, toPayload(input))

	return
}

func (facex *Facex) Search(uri string) (result *SearchResult, err error) {
	client := request.New(request.NewAccessKey(facex.accessKey, facex.secretKey))
	input := NewSearchInput(uri)

	data, err := client.Send("POST", facex.API(GroupSearchAPI), "application/json", facex.timeout, toPayload(input))
	if err != nil {
		return
	}

	result, err = NewSearchResult(data)
	return
}
