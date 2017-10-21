package facex

import (
	"testing"

	"github.com/golib/assert"
)

func TestFaceXAPI(t *testing.T) {
	assertion := assert.New(t)

	facex := NewFacex(&Config{
		Endpoint: "http://argus.atlab.ai",
		GroupId:  "groupId",
	})

	assertion.Equal("http://argus.atlab.ai/v1/face/group/groupId/add", facex.API(GroupAddAPI))
}

func TestFaceXNewGroup(t *testing.T) {
	assertion := assert.New(t)

	facex := NewFacex(&Config{
		Endpoint:  "http://argus.atlab.ai",
		AccessKey: "",
		SecretKey: "",
		GroupId:   "gohack2017",
		Timeout:   8,
	})

	input := NewFacexInput("http://oy5ixix1l.bkt.clouddn.com/face1.png", "1")

	err := facex.NewGroup(input)
	assertion.Nil(err)
}

func TestFaceXSearch(t *testing.T) {
	assertion := assert.New(t)

	facex := NewFacex(&Config{
		Endpoint:  "http://argus.atlab.ai",
		AccessKey: "",
		SecretKey: "",
		GroupId:   "gohack2017",
		Timeout:   8,
	})

	result, err := facex.Search("http://oy5ixix1l.bkt.clouddn.com/miss.png")
	assertion.Nil(err)
	assertion.NotNil(result)

	assertion.False(result.IsOK())

	result, err = facex.Search("http://oy5ixix1l.bkt.clouddn.com/face2.png")
	assertion.Nil(err)
	assertion.NotNil(result)

	assertion.True(result.IsOK())
}

func TestFaceXAddGroup(t *testing.T) {
	assertion := assert.New(t)

	facex := NewFacex(&Config{
		Endpoint:  "http://argus.atlab.ai",
		AccessKey: "",
		SecretKey: "",
		GroupId:   "gohack2017",
		Timeout:   8,
	})

	err := facex.AddFace("http://oy5ixix1l.bkt.clouddn.com/face2.png", "2")
	assertion.Nil(err)
}
