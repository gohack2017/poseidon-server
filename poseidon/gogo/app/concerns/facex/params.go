package facex

import (
	"bytes"
	"encoding/json"
	"io"
)

type FacexInput struct {
	Data []*FacexInputItem `json:"data"`
}

type FacexInputItem struct {
	URI       string            `json:"uri"`
	Attribute map[string]string `json:"attribute"`
}

func NewFacexInput(uri, id string) FacexInput {
	return FacexInput{
		Data: []*FacexInputItem{
			&FacexInputItem{
				URI: uri,
				Attribute: map[string]string{
					"name": id,
				},
			},
		},
	}
}

type SearchInput struct {
	Data map[string]string `json:"data"`
}

type SearchResult struct {
	Message string            `json:"message"`
	Result  *ResultDetections `json:"result"`
}

type ResultDetections struct {
	Detections []*ResultValue `json:"detections"`
}

type ResultValue struct {
	Value *SearchResultValue `json:"value"`
}

type SearchResultValue struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

func NewSearchResult(data []byte) (*SearchResult, error) {
	var ret SearchResult

	err := json.Unmarshal(data, &ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

func (this *SearchResult) IsOK() bool {
	return this.Result != nil &&
		this.Result.Detections != nil &&
		len(this.Result.Detections) > 0 &&
		this.Result.Detections[0].Value.Score > 0.5
}

func (this *SearchResult) Name() string {
	return this.Result.Detections[0].Value.Name
}

func NewSearchInput(uri string) *SearchInput {
	return &SearchInput{
		Data: map[string]string{
			"uri": uri,
		},
	}
}

func toPayload(in interface{}) io.Reader {
	data, _ := json.Marshal(in)

	return bytes.NewBuffer(data)
}
