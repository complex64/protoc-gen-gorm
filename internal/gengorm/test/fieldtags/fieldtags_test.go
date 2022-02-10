package fieldtags_test

import (
	"testing"

	"github.com/complex64/protoc-gen-gorm/internal/gengorm/test/fieldtags"
	"github.com/complex64/protoc-gen-gorm/internal/require"
)

// Assert generation of expected GORM field tags.

func TestOptions(t *testing.T) {
	msg := &fieldtags.OptionsModel{}
	require.StructFieldTags(t, msg, "Column", map[string]string{"gorm": "column:my_column"})
	require.StructFieldTags(t, msg, "NotNull", map[string]string{"gorm": "not null"})
	require.StructFieldTags(t, msg, "Default", map[string]string{"gorm": "default:value"})
	require.StructFieldTags(t, msg, "Unique", map[string]string{"gorm": "unique"})
	require.StructFieldTags(t, msg, "PrimaryKey", map[string]string{"gorm": "primaryKey"})
	require.StructFieldTags(t, msg, "DefaultIndex", map[string]string{"gorm": "index"})
	require.StructFieldTags(t, msg, "NamedIndex", map[string]string{"gorm": "index:my_index"})
	require.StructFieldTags(t, msg, "UniqueDefaultIndex", map[string]string{"gorm": "uniqueIndex"})
	require.StructFieldTags(t, msg, "UniqueNamedIndex", map[string]string{"gorm": "uniqueIndex:my_unique_index"})
	require.StructFieldTags(t, msg, "AutoCreateTime", map[string]string{"gorm": "autoCreateTime"})
	require.StructFieldTags(t, msg, "AutoUpdateTime", map[string]string{"gorm": "autoUpdateTime"})
	require.StructFieldTags(t, msg, "Ignore", map[string]string{"gorm": "-"})
}

func TestDenyOptions(t *testing.T) {
	msg := &fieldtags.DenyOptionsModel{}
	require.StructFieldTags(t, msg, "Ignore", map[string]string{"gorm": "-"})
	require.StructFieldTags(t, msg, "UpdateOnly", map[string]string{"gorm": "->:false;<-:update"})
	require.StructFieldTags(t, msg, "CreateOnly", map[string]string{"gorm": "->:false;<-:create"})
	require.StructFieldTags(t, msg, "WriteOnly", map[string]string{"gorm": "->:false;<-"})
	require.StructFieldTags(t, msg, "ReadOnly", map[string]string{"gorm": "->"})
	require.StructFieldTags(t, msg, "ReadUpdate", map[string]string{"gorm": "<-:update"})
	require.StructFieldTags(t, msg, "ReadCreate", map[string]string{"gorm": "<-:create"})
}
