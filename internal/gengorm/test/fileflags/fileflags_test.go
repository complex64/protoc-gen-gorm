package fileflags_test

import (
	"testing"

	"github.com/complex64/protoc-gen-gorm/internal/gengorm/test/fileflags"
)

func TestGenerateFile(t *testing.T) {
	t.Run("", func(t *testing.T) {
		t.Run("file-level model=true implies model=true", func(t *testing.T) {
			_ = &fileflags.MyMessageModel{}
		})
		t.Run("file-level hooks=true implies model=true", func(t *testing.T) {
			_ = &fileflags.MyHooksMessageModel{}
		})
		t.Run("file-level validate=true implies model=true", func(t *testing.T) {
			_ = &fileflags.MyValidateMessageModel{}
		})
		t.Run("file-level crud=true implies model=true", func(t *testing.T) {
			_ = &fileflags.MyCrudMessageModel{}
		})
	})
}
