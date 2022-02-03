package tabler_test

import (
	"testing"

	"github.com/complex64/protoc-gen-gorm/internal/gengorm/test/tabler"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm/schema"
)

func TestImplementsTabler(t *testing.T) {
	// Expect generation of `TableName() string`.
	// https://gorm.io/docs/conventions.html#TableName
	var _ schema.Tabler = &tabler.ImplementsTablerModel{}
	require.Equal(t, "name", (&tabler.ImplementsTablerModel{}).TableName())
}
