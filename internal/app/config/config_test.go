package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCurrentHost(t *testing.T) {
	t.Run("should return localhost base URL", func(t *testing.T) {
		result := getCurrentHost()
		assert.Equal(t, "http://localhost", result)

	})
}

func TestGetCurrentPort(t *testing.T) {
	t.Run("should return default 8888 port", func(t *testing.T) {
		result := getCurrentPort()
		assert.Equal(t, "8888", result)

	})
}
