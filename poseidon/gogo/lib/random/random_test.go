package random

import (
	"testing"

	"github.com/golib/assert"
)

func Test_Random_Number(t *testing.T) {
	l, _ := Number.GenerateString(6)
	r, _ := Number.GenerateString(6)

	assert.NotEqual(t, l, r)
}
