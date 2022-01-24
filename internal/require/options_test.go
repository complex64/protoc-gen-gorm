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
		msg = &testdata.MyMessage{MessageField_1: "message value 1"}
		opt = &testdata.MessageOptions{OptionField_1: "option value 1"}
	)

	t.Run("matches proto file", func(t *testing.T) {
		require.MessageOption(t, msg, opt)
	})

	t.Run("compares fields", func(t *testing.T) {
		other := &testdata.MessageOptions{OptionField_1: "option value 2"}
		mockT := new(MockT)

		require.MessageOption(mockT, msg, other)
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

func ExampleMessageOption() {
	var t *testing.T // We're inside a test case.

	// message MyMessage {
	//  option (testdata.message).option_field_1 = "option value 1";
	//  ...
	// }

	// A message, annotated as shown above.
	msg := &testdata.MyMessage{
		MessageField_1: "message value 1",
	}

	// The expected option the annotation should map to.
	opt := &testdata.MessageOptions{
		OptionField_1: "option value 1",
	}

	require.MessageOption(t, msg, opt) // Test passes.
}
