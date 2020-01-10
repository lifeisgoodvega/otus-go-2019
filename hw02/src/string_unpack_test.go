package stringunpack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	res, err := Unpack("a4")
	assert.NoError(t, err)
	assert.Equal(t, "aaaa", res)
}

func TestEmptyString(t *testing.T) {
	res, err := Unpack("")
	assert.NoError(t, err)
	assert.Equal(t, "", res)
}

func TestLongNumber(t *testing.T) {
	res, err := Unpack("a256")
	assert.NoError(t, err)
	var expectedString string
	for i := 0; i < 256; i++ {
		expectedString += string('a')
	}
	assert.Equal(t, expectedString, res)
}

func TestComplexString(t *testing.T) {
	res, err := Unpack("a4b3c1")
	assert.NoError(t, err)
	assert.Equal(t, "aaaabbbc", res)
}

func TestEscape(t *testing.T) {
	res, err := Unpack("\\44")
	assert.NoError(t, err)
	assert.Equal(t, "4444", res)
}

func TestComplexStringWithEscape(t *testing.T) {
	res, err := Unpack("abc2\\\\2d3a0")
	assert.NoError(t, err)
	assert.Equal(t, "abcc\\\\ddd", res)
}

func TestNoUnpack(t *testing.T) {
	res, err := Unpack("abc\\4\\5")
	assert.NoError(t, err)
	assert.Equal(t, "abc45", res)
}

func TestInvalidString(t *testing.T) {
	_, err := Unpack("256")
	assert.Error(t, err)
}
