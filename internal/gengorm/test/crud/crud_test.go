package crud_test

import (
	"context"
	"testing"

	"github.com/complex64/protoc-gen-gorm/internal/gengorm/test/crud"
	ireq "github.com/complex64/protoc-gen-gorm/internal/require"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	ctx = context.Background()
)

func TestCrudWithGorm_Create(t *testing.T) {
	t.Run("creates and returns", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))

			x := &crud.Crud{Uuid: "abc"}
			y, err := x.WithDB(db).Create(ctx)
			require.NoError(t, err)
			ireq.EqualProtos(t, x, y)

			var z crud.CrudModel
			require.NoError(t, db.First(&z).Error)
		})
	})
}

func TestCrudWithDB_Get(t *testing.T) {
	x := &crud.Crud{Uuid: "abc"}

	t.Run("returns an error when not found", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))

			out, err := x.WithDB(db).Get(ctx)
			require.Equal(t, gorm.ErrRecordNotFound, err)
			require.Nil(t, out)
		})
	})

	t.Run("returns existing records", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))
			{
				_, err := x.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				out, err := x.WithDB(db).Get(ctx)
				require.NoError(t, err)
				ireq.EqualProtos(t, x, out)
			}
		})
	})
}

func TestCrudWithDB_List(t *testing.T) {
	x := &crud.Crud{Uuid: "abc"}

	t.Run("returns empty list", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))

			out, err := x.WithDB(db).List(ctx)
			require.NoError(t, err)
			require.Empty(t, out)
		})
	})

	t.Run("returns all records", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))
			{
				_, err := x.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				out, err := x.WithDB(db).List(ctx)
				require.NoError(t, err)
				require.Len(t, out, 1)
				ireq.EqualProtos(t, x, out[0])
			}
		})
	})
}

func TestCrudWithDB_Update(t *testing.T) {
	x := &crud.Crud{Uuid: "abc"}

	t.Run("inserts missing records", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))

			out, err := x.WithDB(db).Update(ctx)
			require.NoError(t, err)
			require.NotNil(t, out)
			ireq.EqualProtos(t, x, out)
		})
	})

	t.Run("updates existing records", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))
			{
				_, err := x.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				x.Uuid = "new-uuid"
				out, err := x.WithDB(db).Update(ctx)
				require.NoError(t, err)
				require.NotNil(t, out)
				ireq.EqualProtos(t, x, out)
			}
			{
				out, err := x.WithDB(db).Get(ctx)
				require.NoError(t, err)
				ireq.EqualProtos(t, x, out)
			}
		})
	})
}

func TestCrudWithDB_Patch(t *testing.T) {
	var (
		target = &crud.Crud{Uuid: "abc"}
		record = &crud.Crud{Uuid: "abc", StringField: "foo"}
		update = &crud.Crud{Uuid: "abc", StringField: "bar"}
	)

	t.Run("updates selected fields for existing records", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))
			{
				_, err := record.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				mask := &fieldmaskpb.FieldMask{Paths: []string{"string_field"}}
				out, err := update.WithDB(db).Patch(ctx, mask)
				require.NoError(t, err)
				require.NotNil(t, out)
				ireq.EqualProtos(t, update, out)
			}
			{
				out, err := target.WithDB(db).Get(ctx)
				require.NoError(t, err)
				ireq.EqualProtos(t, update, out)
			}
		})
	})
}

func TestCrudWithDB_Delete(t *testing.T) {
	x := &crud.Crud{Uuid: "abc"}

	t.Run("no effect on missing records", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))

			err := x.WithDB(db).Delete(ctx)
			require.NoError(t, err)
		})
	})

	t.Run("removes existing records", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))
			{
				_, err := x.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				err := x.WithDB(db).Delete(ctx)
				require.NoError(t, err)
			}
			{
				out, err := x.WithDB(db).Get(ctx)
				require.Equal(t, gorm.ErrRecordNotFound, err)
				require.Nil(t, out)
			}
		})
	})
}

func withDB(t *testing.T, f func(db *gorm.DB)) {
	db, err := gorm.Open(sqlite.Open(""), &gorm.Config{})
	require.NoError(t, err)
	// db = db.Debug()
	f(db)
}
