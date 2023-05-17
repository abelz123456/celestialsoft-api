package server

import (
	"testing"

	"github.com/abelz123456/celestial-api/package/config"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	cfg := config.Config{
		DevelopmentPort: "8080",
	}

	s, err := Init(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, s)

	// Check if the engine has been initialized
	assert.NotNil(t, s.Engine)
}
