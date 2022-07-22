package fileflags_test

import (
	"testing"

	"github.com/complex64/protoc-gen-gorm/internal/gengorm/test/fileflags"
)

func TestGenerateFile(t *testing.T) {
	t.Run("", func(t *testing.T) {
		// Assert *Model types are generated:
		t.Run("file-level model=true implies message-level model=true", func(t *testing.T) {
			_ = &fileflags.ModelFlagMessageModel{}
		})
		t.Run("file-level validate=true implies message-level model=true", func(t *testing.T) {
			_ = &fileflags.ValidateFlagMessageModel{}
		})
		t.Run("file-level crud=true implies message-level model=true", func(t *testing.T) {
			_ = &fileflags.CrudFlagMessageModel{}
		})
	})
}
