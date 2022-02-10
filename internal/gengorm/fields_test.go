package gengorm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_permissionTag(t *testing.T) {
	require.Equal(t, "", permTags(false, false, false))
	require.Equal(t, "<-:create", permTags(false, false, true))
	require.Equal(t, "<-:update", permTags(false, true, false))
	require.Equal(t, "->", permTags(false, true, true))
	require.Equal(t, "->:false;<-", permTags(true, false, false))
	require.Equal(t, "->:false;<-:create", permTags(true, false, true))
	require.Equal(t, "->:false;<-:update", permTags(true, true, false))
	require.Equal(t, "-", permTags(true, true, true))
}
