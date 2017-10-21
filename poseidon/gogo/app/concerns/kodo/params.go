package kodo

import (
	"time"

	"github.com/qiniu/api.v7/storage"
)

type ObjectItem struct {
	Key   string    `json:"key"`
	Hash  string    `json:"hash"`
	Size  int64     `json:"size"`
	Mime  string    `json:"mime"`
	Ctime time.Time `json:"ctime"`
}

func NewObjectItem(item *storage.ListItem) *ObjectItem {
	ctime := time.Unix(item.PutTime/int64(time.Second), 0)

	return &ObjectItem{
		Key:   item.Key,
		Hash:  item.Hash,
		Size:  item.Fsize,
		Mime:  item.MimeType,
		Ctime: ctime,
	}
}

type ListObjectsOutput struct {
	Items      []*ObjectItem `json:"items"`
	PrevMarker string        `json:"prev_marker"`
	NextMarker string        `json:"next_marker"`
	MaxKeys    int           `json:"max_keys"`
	HasMore    bool          `json:"has_more"`
}

func NewListObjectsOutput(objects []storage.ListItem, prevMarker, nextMarker string, maxKeys int, hasMore bool) *ListObjectsOutput {
	result := &ListObjectsOutput{
		Items:      make([]*ObjectItem, len(objects)),
		PrevMarker: prevMarker,
		NextMarker: nextMarker,
		MaxKeys:    maxKeys,
		HasMore:    hasMore,
	}
	for i, obj := range objects {
		result.Items[i] = NewObjectItem(&obj)
	}

	return result
}
