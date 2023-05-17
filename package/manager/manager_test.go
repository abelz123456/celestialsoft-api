package manager

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessInit(t *testing.T) {
	_, err := Init(".")
	assert.Nil(t, err)
}
