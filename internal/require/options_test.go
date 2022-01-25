package require_test

import (
	"testing"

	"github.com/complex64/protoc-gen-gorm/internal/require"
	"github.com/complex64/protoc-gen-gorm/internal/require/testdata"
)

// MockT allows us to test assertion functions.
type MockT struct {
	Failed bool
}

func (t *MockT) FailNow() {
	t.Failed = true
}

func (t *MockT) Errorf(format string, args ...interface{}) {
	_, _ = format, args
}

func TestMessageOption(t *testing.T) {
	var (
		msg  = &testdata.MyMessage{MessageField: "field value"}
		opts = &testdata.MessageOptions{MessageOption: "message option value"}
	)

	t.Run("matches proto file", func(t *testing.T) {
		require.MessageOption(t, opts, msg)
	})

	t.Run("compares fields", func(t *testing.T) {
		other := &testdata.MessageOptions{MessageOption: "unexpected value"}
		mockT := new(MockT)

		require.MessageOption(mockT, other, msg)
		if !mockT.Failed {
			t.Error("Check should fail")
		}
	})

	t.Run("no such option", func(t *testing.T) {
		mockT := new(MockT)

		require.MessageOption(mockT, msg, msg)
		if !mockT.Failed {
			t.Error("Check should fail")
		}
	})
}

func TestFileOptions(t *testing.T) {
	var (
		msg  = &testdata.MyMessage{MessageField: "field value"}
		opts = &testdata.FileOptions{FileOption: "file option value"}
	)

	t.Run("matches proto file", func(t *testing.T) {
		require.FileOptions(t, opts, msg)
	})

	t.Run("compares fields", func(t *testing.T) {
		other := &testdata.FileOptions{FileOption: "unexpected value"}
		mockT := new(MockT)

		require.FileOptions(mockT, other, msg)
		if !mockT.Failed {
			t.Error("Check should fail")
		}
	})

	t.Run("no such option", func(t *testing.T) {
		mockT := new(MockT)

		require.FileOptions(mockT, msg, msg)
		if !mockT.Failed {
			t.Error("Check should fail")
		}
	})
}

func ExampleMessageOption() {
	var t *testing.T // We're inside a test case.

	// message MyMessage {
	//  option (testdata.message).message_option = "option value";
	//  ...
	// }

	// A message, annotated as shown above.
	msg := &testdata.MyMessage{
		MessageField: "field value",
	}

	// The expected option the annotation should map to.
	opts := &testdata.MessageOptions{
		MessageOption: "message option value",
	}

	require.MessageOption(t, opts, msg) // Test passes.
}

func ExampleFileOptions() {
	var t *testing.T // We're inside a test case.

	// In options.proto:
	// option (testdata.file).file_option = "file option value";
	// message MyMessage {}

	// A message from options.proto which links to its parent file descriptor.
	msg := &testdata.MyMessage{}
	opts := &testdata.FileOptions{
		// Assertions...
	}

	require.FileOptions(t, opts, msg)
}
