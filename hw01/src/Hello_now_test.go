package hellonow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	err := Now()
	assert.NoError(t, err)
}
