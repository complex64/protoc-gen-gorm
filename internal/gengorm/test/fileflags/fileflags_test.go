package fileflags_test

import (
	"testing"

	"github.com/complex64/protoc-gen-gorm/internal/gengorm/test/fileflags"
)

func TestGenerateFile(t *testing.T) {
	t.Run("", func(t *testing.T) {
		t.Run("file-level model=true implies model=true", func(t *testing.T) {
			_ = &fileflags.ImplicitModel{}
		})
		t.Run("file-level validate=true implies model=true", func(t *testing.T) {
			_ = &fileflags.ImplicitValidate{}
		})
		t.Run("file-level crud=true implies model=true", func(t *testing.T) {
			_ = &fileflags.ImplicitCRUD{}
		})
	})
}
