package require

import (
	"github.com/stretchr/testify/require"
)

func StructFieldTags(t require.TestingT, value interface{}, name string, tags map[string]string) {
	field := Field(t, value, name)
	for key, value := range tags {
		have, ok := field.Tag.Lookup(key)
		require.True(t, ok, "expected field '%s' to have tag '%s'", field.Name, key)
		require.Equal(t, value, have)
	}
}
