package oembed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsValidURL(t *testing.T) {
	assert := assert.New(t)
	t.Run("invalid", func(t *testing.T) {
		assert.False(isValidURL("str"))
	})
	t.Run("valid", func(t *testing.T) {
		assert.True(isValidURL("http://www.kickstarter.com"))
	})
}
