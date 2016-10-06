package hammingweight

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOfMethod(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(0, Of(0))
	assert.Equal(1, Of(1))
	assert.Equal(1, Of(2))
	assert.Equal(16, Of(65535))
	assert.Equal(32, Of(0xFFFFFFFF))

}
