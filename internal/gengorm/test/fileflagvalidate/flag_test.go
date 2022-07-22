package fileflagvalidate_test

import (
	"testing"

	"github.com/complex64/protoc-gen-gorm/internal/gengorm/test/fileflagvalidate"
)

func TestGenerateFile(t *testing.T) {
	t.Run("", func(t *testing.T) {
		t.Run("file-level validate=true implies model=true", func(t *testing.T) {
			_ = &fileflagvalidate.MyMessageModel{}
		})
	})
}
