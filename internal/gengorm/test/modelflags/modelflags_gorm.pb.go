// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm 2.0.0
// 	protoc          (unknown)
// source: modelflags/modelflags.proto

package modelflags

import (
	context "context"
	fmt "fmt"
	_ "github.com/complex64/protoc-gen-gorm/gormpb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	gorm "gorm.io/gorm"
)

// ModelOptionModel is the GORM model for modelflags.ModelOption.
type ModelOptionModel struct {
}

// ToProto converts a ModelOptionModel to its protobuf representation.
func (m *ModelOptionModel) ToProto() (*ModelOption, error) {
	p := new(ModelOption)
	return p, nil
}

// ToModel converts a ModelOption to its GORM model.
func (p *ModelOption) ToModel() (*ModelOptionModel, error) {
	m := new(ModelOptionModel)
	return m, nil
}

// ValidateImpliesModelModel is the GORM model for modelflags.ValidateImpliesModel.
type ValidateImpliesModelModel struct {
}

// ToProto converts a ValidateImpliesModelModel to its protobuf representation.
func (m *ValidateImpliesModelModel) ToProto() (*ValidateImpliesModel, error) {
	p := new(ValidateImpliesModel)
	return p, nil
}

// ToModel converts a ValidateImpliesModel to its GORM model.
func (p *ValidateImpliesModel) ToModel() (*ValidateImpliesModelModel, error) {
	m := new(ValidateImpliesModelModel)
	return m, nil
}

// CRUDImpliesModelModel is the GORM model for modelflags.CRUDImpliesModel.
type CRUDImpliesModelModel struct {
	Uuid string `gorm:"primaryKey"`
}

// ToProto converts a CRUDImpliesModelModel to its protobuf representation.
func (m *CRUDImpliesModelModel) ToProto() (*CRUDImpliesModel, error) {
	p := new(CRUDImpliesModel)
	p.Uuid = m.Uuid
	return p, nil
}

// ToModel converts a CRUDImpliesModel to its GORM model.
func (p *CRUDImpliesModel) ToModel() (*CRUDImpliesModelModel, error) {
	m := new(CRUDImpliesModelModel)
	m.Uuid = p.Uuid
	return m, nil
}

type CRUDImpliesModelGetOption func(tx *gorm.DB) *gorm.DB
type CRUDImpliesModelListOption func(tx *gorm.DB) *gorm.DB

type CRUDImpliesModelWithDB struct {
	p  *CRUDImpliesModel
	db *gorm.DB
}

func (p *CRUDImpliesModel) WithDB(db *gorm.DB) CRUDImpliesModelWithDB {
	return CRUDImpliesModelWithDB{p: p, db: db}
}

func (c CRUDImpliesModelWithDB) Create(ctx context.Context) (*CRUDImpliesModel, error) {
	if c.p == nil {
		return nil, nil
	}
	m, err := c.p.ToModel()
	if err != nil {
		return nil, err
	}
	db := c.db.WithContext(ctx)
	if err := db.Create(m).Error; err != nil {
		return nil, err
	}
	if y, err := m.ToProto(); err != nil {
		return nil, err
	} else {
		return y, nil
	}
}

func (c CRUDImpliesModelWithDB) Get(ctx context.Context, opts ...CRUDImpliesModelGetOption) (*CRUDImpliesModel, error) {
	if c.p == nil {
		return nil, nil
	}
	var zero string
	if c.p.Uuid == zero {
		return nil, fmt.Errorf("empty primary key")
	}
	m, err := c.p.ToModel()
	if err != nil {
		return nil, err
	}
	db := c.db.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	out := CRUDImpliesModelModel{}
	if err := db.Where(m).First(&out).Error; err != nil {
		return nil, err
	}
	if p, err := out.ToProto(); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

func (c CRUDImpliesModelWithDB) List(ctx context.Context, opts ...CRUDImpliesModelListOption) ([]*CRUDImpliesModel, error) {
	if c.p == nil {
		return nil, nil
	}
	db := c.db.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	var ms []CRUDImpliesModelModel
	if err := db.Find(&ms).Error; err != nil {
		return nil, err
	}
	protos := make([]*CRUDImpliesModel, 0, len(ms))
	for _, m := range ms {
		if p, err := m.ToProto(); err != nil {
			return nil, err
		} else {
			protos = append(protos, p)
		}
	}
	return protos, nil
}

func (c CRUDImpliesModelWithDB) Update(ctx context.Context) (*CRUDImpliesModel, error) {
	if c.p == nil {
		return nil, nil
	}
	m, err := c.p.ToModel()
	if err != nil {
		return nil, err
	}
	db := c.db.WithContext(ctx)
	if err := db.Save(m).Error; err != nil {
		return nil, err
	}
	return c.Get(ctx)
}

func (c CRUDImpliesModelWithDB) Patch(ctx context.Context, mask *fieldmaskpb.FieldMask) error {
	if c.p == nil {
		return nil
	}
	if mask == nil {
		_, err := c.Update(ctx)
		return err
	}
	if !mask.IsValid(c.p) {
		return fmt.Errorf("invalid field mask")
	}
	paths := mask.Paths
	if len(paths) == 0 {
		_, err := c.Update(ctx)
		return err
	}
	var zero string
	if c.p.Uuid == zero {
		return fmt.Errorf("empty primary key")
	}
	m, err := c.p.ToModel()
	if err != nil {
		return err
	}
	target := CRUDImpliesModelModel{Uuid: m.Uuid}
	cols := LookupCRUDImpliesModelModelColumns(paths)
	db := c.db.WithContext(ctx)
	if err := db.Model(&target).Select(cols).Updates(m).Error; err != nil {
		return err
	}
	return nil
}

func (c CRUDImpliesModelWithDB) Delete(ctx context.Context) error {
	if c.p == nil {
		return nil
	}
	var zero string
	if c.p.Uuid == zero {
		return fmt.Errorf("empty primary key")
	}
	m, err := c.p.ToModel()
	if err != nil {
		return err
	}
	db := c.db.WithContext(ctx)
	if err := db.Where(m).Delete(&CRUDImpliesModelModel{}).Error; err != nil {
		return err
	}
	return nil
}

func WithCRUDImpliesModelGetFieldMask(mask *fieldmaskpb.FieldMask) CRUDImpliesModelGetOption {
	return func(tx *gorm.DB) *gorm.DB {
		cols := LookupCRUDImpliesModelModelColumns(mask.Paths)
		tx = tx.Select(cols)
		return tx
	}
}

func WithCRUDImpliesModelListFilter(filter string) CRUDImpliesModelListOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx
	}
}

func WithCRUDImpliesModelListLimit(n int) CRUDImpliesModelListOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Limit(n)
	}
}

func WithCRUDImpliesModelListFieldMask(mask *fieldmaskpb.FieldMask) CRUDImpliesModelListOption {
	return func(tx *gorm.DB) *gorm.DB {
		cols := LookupCRUDImpliesModelModelColumns(mask.Paths)
		tx = tx.Select(cols)
		return tx
	}
}

func WithCRUDImpliesModelListOffset(n int) CRUDImpliesModelListOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(n)
	}
}

func WithCRUDImpliesModelListOrder(order string) CRUDImpliesModelListOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Order(order)
	}
}

var fieldColumnsCRUDImpliesModelModel = map[string]string{
	"uuid": "Uuid",
}

func LookupCRUDImpliesModelModelColumn(field string) string {
	if col, ok := fieldColumnsCRUDImpliesModelModel[field]; ok {
		return col
	} else {
		panic(field)
	}
}

func LookupCRUDImpliesModelModelColumns(paths []string) (cols []string) {
	for _, p := range paths {
		cols = append(cols, LookupCRUDImpliesModelModelColumn(p))
	}
	return
}
