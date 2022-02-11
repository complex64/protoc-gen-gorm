package require

import (
	"reflect"

	"github.com/stretchr/testify/require"
)

func Field(t require.TestingT, value interface{}, name string) *reflect.StructField {
	field, ok := reflect.TypeOf(value).Elem().FieldByName(name)
	require.True(t, ok, "expected struct '%v' to have field '%s'", reflect.ValueOf(value), name)
	if !ok {
		return nil
	}
	return &field
}

func FieldType(t require.TestingT, value interface{}, name string, sample interface{}) {
	if reflect.TypeOf(sample) == nil {
		t.Errorf("nil sample")
		t.FailNow()
		return
	}
	field := Field(t, value, name)
	if field == nil {
		t.FailNow()
		return
	}
	want := reflect.TypeOf(sample)
	require.Equal(t, want, field.Type,
		"expected struct '%+v' to have field '%s' of type '%s', not '%v'",
		reflect.ValueOf(value),
		name,
		want.String(),
		field.Type.String())
}
