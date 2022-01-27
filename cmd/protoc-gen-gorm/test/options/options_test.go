package options_test

import (
	// Assert compilation.

	"testing"

	"github.com/complex64/protoc-gen-gorm/cmd/protoc-gen-gorm/test/options"
	"github.com/complex64/protoc-gen-gorm/gormpb/v2"
	"github.com/complex64/protoc-gen-gorm/internal/require"
)

// Test that all options in the .proto file are present as expected.

func TestMessageOptions(t *testing.T) {
	msg := &options.MyMessage{}
	opts := &gormpb.MessageOptions{
		// Expect all to be present and false by default.
	}
	require.MessageOption(t, opts, msg)
}

func TestFileOptions(t *testing.T) {
	msg := &options.MyMessage{}
	opts := &gormpb.FileOptions{
		// Expect all to be present and false by default.
	}
	require.FileOptions(t, opts, msg)
}
