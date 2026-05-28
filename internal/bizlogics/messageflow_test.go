package bizlogics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageFlow_1(t *testing.T) {
	t.Run("返却値がTrueであること", func(t *testing.T) {
		assert.True(t, messageFlow("Hello, World!"))
	})
}

func TestMessageFlow_2(t *testing.T) {
	t.Run("返却値がFalseであること", func(t *testing.T) {
		assert.False(t, messageFlow("Hello, World!"))
	})
}
