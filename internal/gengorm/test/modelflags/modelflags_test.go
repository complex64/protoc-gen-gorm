package modelflags_test

import (
	"testing"

	"github.com/complex64/protoc-gen-gorm/internal/gengorm/test/modelflags"
)

func TestGenerateFile(t *testing.T) {
	t.Run("generates model type when model=true", func(t *testing.T) {
		_ = &modelflags.ModelOption{}
	})
	t.Run("validate=true implies model=true", func(t *testing.T) {
		_ = &modelflags.ValidateImpliesModel{}
	})
	t.Run("crud=true implies model=true", func(t *testing.T) {
		_ = &modelflags.CRUDImpliesModel{}
	})
}
