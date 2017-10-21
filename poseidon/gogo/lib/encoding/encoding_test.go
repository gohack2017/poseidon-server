package encoding

import (
	"io/ioutil"
	"path"
	"testing"

	"github.com/golib/assert"
	"github.com/qbox/facecloud/app/concerns/kodo"
)

var (
	fixPath = path.Clean("../../tmp/fixtures")
)

func Test_Encoding(t *testing.T) {
	sn := "010f0ff8-cf2e-4cb6-bae1-2c6858569d65"
	face, _ := ioutil.ReadFile(path.Join(fixPath, "gopher.png"))
	encface, _ := ioutil.ReadFile(path.Join(fixPath, "enc-gopher.png"))

	user := kodo.NewUser(sn)
	encoding := New(sn, user.KodoKey())

	encdata := encoding.Encrypt(face)
	assert.Equal(t, string(encface), encdata)

	decdata, err := encoding.Decrypt(string(encface))
	assert.Nil(t, err)
	assert.Equal(t, face, decdata)
}
