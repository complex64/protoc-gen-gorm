package require

import (
	"reflect"

	"github.com/stretchr/testify/require"
)

func StructFieldTags(t require.TestingT, i interface{}, field string, tags map[string]string) {
	f, ok := reflect.TypeOf(i).Elem().FieldByName(field)
	require.True(t, ok, "expected struct '%v' to have field '%s'", i, field)

	for key, value := range tags {
		have, ok := f.Tag.Lookup(key)
		require.True(t, ok, "expected field '%s' to have tag '%s'", f.Name, key)
		require.Equal(t, value, have)
	}
}
