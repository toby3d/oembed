package oembed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSync(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		t.Run("source url", func(t *testing.T) {
			assert.NoError(t, Sync("wtf"))
			assert.NotZero(t, len(Providers))
		})
		t.Run("resource body", func(t *testing.T) {
			assert.Error(t, Sync("https://ddg.gg/"))
			assert.NotZero(t, len(Providers))
		})
	})
	t.Run("valid url", func(t *testing.T) {
		assert.NoError(t, Sync(SourceURL))
		assert.NotZero(t, len(Providers))
	})
}
