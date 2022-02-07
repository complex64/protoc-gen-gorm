package require_test

import (
	"testing"

	"github.com/complex64/protoc-gen-gorm/internal/require"
)

func TestFieldType(t *testing.T) {
	type S struct {
		field string
	}
	var (
		value = &S{field: "value"}
	)

	t.Run("existing field", func(t *testing.T) {
		mockT := new(MockT)
		require.FieldType(t, value, "field", "sample")
		if mockT.Failed {
			t.Error("Check should pass")
		}
	})

	t.Run("mismatched type", func(t *testing.T) {
		mockT := new(MockT)
		require.FieldType(mockT, value, "field", int(0))
		if !mockT.Failed {
			t.Error("Check should fail")
		}
	})

	t.Run("missing field", func(t *testing.T) {
		mockT := new(MockT)
		require.FieldType(mockT, value, "missing", "ignored")
		if !mockT.Failed {
			t.Error("Check should fail")
		}
	})

	t.Run("nil sample", func(t *testing.T) {
		mockT := new(MockT)
		require.FieldType(mockT, value, "missing", nil)
		if !mockT.Failed {
			t.Error("Check should fail")
		}
	})
}
