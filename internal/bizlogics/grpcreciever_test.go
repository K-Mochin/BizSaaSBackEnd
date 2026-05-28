package bizlogics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrpcReciever(t *testing.T) {
	t.Run("返却値がTrueであること", func(t *testing.T) {
		assert.True(t, grpcreciever("Hello, World!"))
	})
}
