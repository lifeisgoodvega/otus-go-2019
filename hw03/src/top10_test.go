package top10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	res := Top10("abc cde")
	assert.Equal(t, []string{"abc", "cde"}, res)
}

func TestSimple2(t *testing.T) {
	res := Top10("abc cde cde")
	assert.Equal(t, []string{"cde", "abc"}, res)
}

func TestSimple3(t *testing.T) {
	res := Top10("aaa")
	assert.Equal(t, []string{"aaa"}, res)
}

func TestTenUnique(t *testing.T) {
	res := Top10("a b c d e f g h i k")
	assert.Equal(t, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "k"}, res)
}

func TestEmpty(t *testing.T) {
	res := Top10("")
	assert.Equal(t, []string(nil), res)
}

func TestComplex(t *testing.T) {
	res := Top10("a b a c d a d e e e f f f g")
	assert.Equal(t, []string{"a", "e", "f", "d", "b", "c", "g"}, res)
}

func TestMoreThanTen(t *testing.T) {
	res := Top10("a b b c c c d d d d e e e e e f f f f f f g g g g g g g h h h h h h h h i i i i i i i i i k k k k k k k k k k l l l l l l l l l l l")
	assert.Equal(t, []string{"l", "k", "i", "h", "g", "f", "e", "d", "c", "b"}, res)
}

func TestRealText(t *testing.T) {
	res := Top10("Пол-литра - это 500 миллилитров, а литр - это 1000 миллилитров")
	assert.Equal(t, []string{"миллилитров", "это", "1000", "500", "а", "литр", "пол-литра"}, res)
}
