package fileflagmodel_test

import (
	"testing"

	"github.com/complex64/protoc-gen-gorm/internal/gengorm/test/fileflagmodel"
)

func TestGenerateFile(t *testing.T) {
	t.Run("", func(t *testing.T) {
		t.Run("file-level model=true implies model=true", func(t *testing.T) {
			_ = &fileflagmodel.MyMessageModel{}
		})
	})
}
