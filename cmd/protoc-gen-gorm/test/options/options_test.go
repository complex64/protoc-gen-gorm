package options_test

import (
	// Assert compilation.

	"testing"

	"github.com/complex64/protoc-gen-gorm/cmd/protoc-gen-gorm/test/options"
	"github.com/complex64/protoc-gen-gorm/gormpb/v2"
	"github.com/complex64/protoc-gen-gorm/internal/require"
)

// Test that all options in the .proto file are present as expected.
func TestMyMessage(t *testing.T) {
	t.Run("message options", func(t *testing.T) {
		msg := &options.MyMessage{}
		opt := &gormpb.MessageOptions{Enabled: true}
		require.MessageOption(t, msg, opt)
	})
}
