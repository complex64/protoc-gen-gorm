package fileflagcrud_test

import (
	"testing"

	"github.com/complex64/protoc-gen-gorm/internal/gengorm/test/fileflagcrud"
)

func TestGenerateFile(t *testing.T) {
	t.Run("", func(t *testing.T) {
		t.Run("file-level crud=true implies model=true", func(t *testing.T) {
			_ = &fileflagcrud.MyMessageModel{}
		})
	})
}
