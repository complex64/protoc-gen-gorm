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

			p := &crud.Crud{Uuid: "abc"}
			y, err := p.WithDB(db).Create(ctx)
			require.NoError(t, err)
			ireq.EqualProtos(t, p, y)

			var z crud.CrudModel
			require.NoError(t, db.First(&z).Error)
		})
	})
}

func TestCrudWithDB_Get(t *testing.T) {
	p := &crud.Crud{
		Uuid:        "abc",
		StringField: "a string",
		Int32Field:  123,
		BoolField:   true,
	}

	t.Run("returns an error when not found", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))

			out, err := p.WithDB(db).Get(ctx)
			require.Equal(t, gorm.ErrRecordNotFound, err)
			require.Nil(t, out)
		})
	})

	t.Run("returns existing records", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))
			{
				_, err := p.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				out, err := p.WithDB(db).Get(ctx)
				require.NoError(t, err)
				ireq.EqualProtos(t, p, out)
			}
		})
	})

	t.Run("respects field mask", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))
			{
				_, err := p.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				model := p.WithDB(db)
				mask := &fieldmaskpb.FieldMask{Paths: []string{
					"string_field",
					"bool_field",
				}}
				out, err := model.Get(ctx, crud.WithCrudGetFieldMask(mask))
				require.NoError(t, err)
				expected := &crud.Crud{
					StringField: p.StringField,
					BoolField:   p.BoolField,
				}
				ireq.EqualProtos(t, expected, out)
			}
		})
	})
}

func TestCrudWithDB_List(t *testing.T) {
	p := &crud.Crud{Uuid: "abc", BoolField: false}
	y := &crud.Crud{Uuid: "def", BoolField: true}

	t.Run("returns empty list", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))

			out, err := p.WithDB(db).List(ctx)
			require.NoError(t, err)
			require.Empty(t, out)
		})
	})

	t.Run("returns all records", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))
			{
				_, err := p.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				out, err := p.WithDB(db).List(ctx)
				require.NoError(t, err)
				require.Len(t, out, 1)
				ireq.EqualProtos(t, p, out[0])
			}
		})
	})

	t.Run("respects field mask", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))
			{
				_, err := p.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				_, err := y.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				mask := &fieldmaskpb.FieldMask{Paths: []string{"bool_field"}}
				out, err := p.WithDB(db).List(ctx, crud.WithCrudListFieldMask(mask))
				require.NoError(t, err)
				require.Len(t, out, 2)
				ireq.EqualProtos(t, &crud.Crud{BoolField: false}, out[0])
				ireq.EqualProtos(t, &crud.Crud{BoolField: true}, out[1])
			}
		})
	})

	t.Run("paginates", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))
			{
				_, err := p.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				_, err := y.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				out, err := p.WithDB(db).List(ctx,
					crud.WithCrudListOffset(0),
					crud.WithCrudListLimit(1),
				)
				require.NoError(t, err)
				require.Len(t, out, 1)
				ireq.EqualProtos(t, p, out[0])
			}
			{
				out, err := p.WithDB(db).List(ctx,
					crud.WithCrudListOffset(1),
					crud.WithCrudListLimit(1),
				)
				require.NoError(t, err)
				require.Len(t, out, 1)
				ireq.EqualProtos(t, y, out[0])
			}
		})
	})

	t.Run("paginates", func(t *testing.T) {
		sortCol := crud.LookupCrudModelColumn("uuid")

		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))
			{
				_, err := p.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				_, err := y.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				out, err := p.WithDB(db).List(ctx,
					crud.WithCrudListOrder(sortCol),
				)
				require.NoError(t, err)
				require.Len(t, out, 2)
				ireq.EqualProtos(t, p, out[0])
				ireq.EqualProtos(t, y, out[1])
			}
			{
				out, err := p.WithDB(db).List(ctx,
					crud.WithCrudListOrder(sortCol+" desc"),
				)
				require.NoError(t, err)
				require.Len(t, out, 2)
				ireq.EqualProtos(t, y, out[0])
				ireq.EqualProtos(t, p, out[1])
			}
		})
	})
}

func TestCrudWithDB_Update(t *testing.T) {
	p := &crud.Crud{Uuid: "abc"}

	t.Run("inserts missing records", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))

			out, err := p.WithDB(db).Update(ctx)
			require.NoError(t, err)
			require.NotNil(t, out)
			ireq.EqualProtos(t, p, out)
		})
	})

	t.Run("updates existing records", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))
			{
				_, err := p.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				p.Uuid = "new-uuid"
				out, err := p.WithDB(db).Update(ctx)
				require.NoError(t, err)
				require.NotNil(t, out)
				ireq.EqualProtos(t, p, out)
			}
			{
				out, err := p.WithDB(db).Get(ctx)
				require.NoError(t, err)
				ireq.EqualProtos(t, p, out)
			}
		})
	})
}

func TestCrudWithDB_Patch(t *testing.T) {
	var (
		record   = &crud.Crud{Uuid: "abc", StringField: "foo", Int32Field: 10, BoolField: true}
		update   = &crud.Crud{Uuid: "abc", StringField: "bar", BoolField: false}
		expected = &crud.Crud{Uuid: "abc", StringField: "bar", Int32Field: 10, BoolField: false}
		target   = &crud.Crud{Uuid: "abc"}
	)

	t.Run("updates selected fields for existing records", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))
			{
				_, err := record.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				mask := &fieldmaskpb.FieldMask{Paths: []string{
					"string_field",
					"bool_field",
				}}
				err := update.WithDB(db).Patch(ctx, mask)
				require.NoError(t, err)
			}
			{
				out, err := target.WithDB(db).Get(ctx)
				require.NoError(t, err)
				ireq.EqualProtos(t, expected, out)
			}
		})
	})
}

func TestCrudWithDB_Delete(t *testing.T) {
	p := &crud.Crud{Uuid: "abc"}

	t.Run("no effect on missing records", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))

			err := p.WithDB(db).Delete(ctx)
			require.NoError(t, err)
		})
	})

	t.Run("removes existing records", func(t *testing.T) {
		withDB(t, func(db *gorm.DB) {
			require.NoError(t, db.AutoMigrate(&crud.CrudModel{}))
			{
				_, err := p.WithDB(db).Create(ctx)
				require.NoError(t, err)
			}
			{
				err := p.WithDB(db).Delete(ctx)
				require.NoError(t, err)
			}
			{
				out, err := p.WithDB(db).Get(ctx)
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
