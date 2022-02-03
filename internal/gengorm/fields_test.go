package gengorm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_permissionTag(t *testing.T) {
	require.Equal(t, "", permissionTag(false, false, false))
	require.Equal(t, "<-:create", permissionTag(false, false, true))
	require.Equal(t, "<-:update", permissionTag(false, true, false))
	require.Equal(t, "->", permissionTag(false, true, true))
	require.Equal(t, "->:false;<-", permissionTag(true, false, false))
	require.Equal(t, "->:false;<-:create", permissionTag(true, false, true))
	require.Equal(t, "->:false;<-:update", permissionTag(true, true, false))
	require.Equal(t, "-", permissionTag(true, true, true))
}
