package require_test

import (
	"testing"

	"github.com/complex64/protoc-gen-gorm/internal/require"
)

func TestStructFieldTags(t *testing.T) {
	type S struct {
		tagged   string `foo:"bar"`
		untagged string
	}

	t.Run("tagged field", func(t *testing.T) {
		mockT := new(MockT)
		s := &S{tagged: "value", untagged: "value"}
		require.StructFieldTags(mockT, s, "tagged", map[string]string{"foo": "bar"})
		if mockT.Failed {
			t.Error("Check should pass")
		}
	})

	t.Run("untagged field", func(t *testing.T) {
		mockT := new(MockT)
		require.StructFieldTags(mockT, new(S), "untagged", nil)
		if mockT.Failed {
			t.Error("Check should pass")
		}
	})

	t.Run("missing field", func(t *testing.T) {
		mockT := new(MockT)
		require.StructFieldTags(mockT, new(S), "missing", nil)
		if !mockT.Failed {
			t.Error("Check should fail")
		}
	})

	t.Run("missing tag", func(t *testing.T) {
		mockT := new(MockT)
		require.StructFieldTags(mockT, new(S), "tagged", map[string]string{"missing": "missing"})
		if !mockT.Failed {
			t.Error("Check should fail")
		}
	})
}
