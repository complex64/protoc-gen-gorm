// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm 2.0.0
// 	protoc          (unknown)
// source: users.proto

package pb

import (
	context "context"
	sql "database/sql"
	fmt "fmt"
	_ "github.com/complex64/protoc-gen-gorm/gormpb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	gorm "gorm.io/gorm"
	time "time"
)

// UserModel is the GORM model for pb.User.
type UserModel struct {
	Name       string       `gorm:"not null;unique;primaryKey;index;<-:create"`
	CreateTime sql.NullTime `gorm:"not null;autoCreateTime"`
	UpdateTime sql.NullTime `gorm:"autoUpdateTime"`
	GivenName  string
	FamilyName string
	Email      string `gorm:"not null"`
}

// ToProto converts a UserModel to its protobuf representation.
func (m *UserModel) ToProto() (*User, error) {
	p := new(User)
	p.Name = m.Name
	if m.CreateTime.Valid && m.CreateTime.Time != (time.Time{}) {
		p.CreateTime = timestamppb.New(m.CreateTime.Time)
	}
	if m.UpdateTime.Valid && m.UpdateTime.Time != (time.Time{}) {
		p.UpdateTime = timestamppb.New(m.UpdateTime.Time)
	}
	p.GivenName = m.GivenName
	p.FamilyName = m.FamilyName
	p.Email = m.Email
	return p, nil
}

// ToModel converts a User to its GORM model.
func (p *User) ToModel() (*UserModel, error) {
	m := new(UserModel)
	m.Name = p.Name
	if t := p.CreateTime; t != nil {
		m.CreateTime = sql.NullTime{
			Valid: true,
			Time:  t.AsTime(),
		}
	}
	if t := p.UpdateTime; t != nil {
		m.UpdateTime = sql.NullTime{
			Valid: true,
			Time:  t.AsTime(),
		}
	}
	m.GivenName = p.GivenName
	m.FamilyName = p.FamilyName
	m.Email = p.Email
	return m, nil
}

func (m *UserModel) TableName() string {
	return "users"
}

type UserGetOption func(tx *gorm.DB) *gorm.DB
type UserListOption func(tx *gorm.DB) *gorm.DB

type UserWithDB struct {
	p  *User
	db *gorm.DB
}

func (p *User) WithDB(db *gorm.DB) UserWithDB {
	return UserWithDB{p: p, db: db}
}

func (c UserWithDB) Create(ctx context.Context) (*User, error) {
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

func (c UserWithDB) Get(ctx context.Context, opts ...UserGetOption) (*User, error) {
	if c.p == nil {
		return nil, nil
	}
	var zero string
	if c.p.Name == zero {
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
	out := UserModel{}
	if err := db.Where(m).First(&out).Error; err != nil {
		return nil, err
	}
	if p, err := out.ToProto(); err != nil {
		return nil, err
	} else {
		return p, nil
	}
}

func (c UserWithDB) List(ctx context.Context, opts ...UserListOption) ([]*User, error) {
	if c.p == nil {
		return nil, nil
	}
	db := c.db.WithContext(ctx)
	for _, opt := range opts {
		db = opt(db)
	}
	var ms []UserModel
	if err := db.Find(&ms).Error; err != nil {
		return nil, err
	}
	protos := make([]*User, 0, len(ms))
	for _, m := range ms {
		if p, err := m.ToProto(); err != nil {
			return nil, err
		} else {
			protos = append(protos, p)
		}
	}
	return protos, nil
}

func (c UserWithDB) Update(ctx context.Context) (*User, error) {
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

func (c UserWithDB) Patch(ctx context.Context, mask *fieldmaskpb.FieldMask) error {
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
	if c.p.Name == zero {
		return fmt.Errorf("empty primary key")
	}
	m, err := c.p.ToModel()
	if err != nil {
		return err
	}
	target := UserModel{Name: m.Name}
	cols := LookupUserModelColumns(paths)
	db := c.db.WithContext(ctx)
	if err := db.Model(&target).Select(cols).Updates(m).Error; err != nil {
		return err
	}
	return nil
}

func (c UserWithDB) Delete(ctx context.Context) error {
	if c.p == nil {
		return nil
	}
	var zero string
	if c.p.Name == zero {
		return fmt.Errorf("empty primary key")
	}
	m, err := c.p.ToModel()
	if err != nil {
		return err
	}
	db := c.db.WithContext(ctx)
	if err := db.Where(m).Delete(&UserModel{}).Error; err != nil {
		return err
	}
	return nil
}

func WithUserGetFieldMask(mask *fieldmaskpb.FieldMask) UserGetOption {
	return func(tx *gorm.DB) *gorm.DB {
		cols := LookupUserModelColumns(mask.Paths)
		tx = tx.Select(cols)
		return tx
	}
}

func WithUserListFilter(filter string) UserListOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx
	}
}

func WithUserListLimit(n int) UserListOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Limit(n)
	}
}

func WithUserListFieldMask(mask *fieldmaskpb.FieldMask) UserListOption {
	return func(tx *gorm.DB) *gorm.DB {
		cols := LookupUserModelColumns(mask.Paths)
		tx = tx.Select(cols)
		return tx
	}
}

func WithUserListOffset(n int) UserListOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Offset(n)
	}
}

func WithUserListOrder(order string) UserListOption {
	return func(tx *gorm.DB) *gorm.DB {
		return tx.Order(order)
	}
}

var fieldColumnsUserModel = map[string]string{
	"name":        "Name",
	"create_time": "CreateTime",
	"update_time": "UpdateTime",
	"given_name":  "GivenName",
	"family_name": "FamilyName",
	"email":       "Email",
}

func LookupUserModelColumn(field string) string {
	if col, ok := fieldColumnsUserModel[field]; ok {
		return col
	} else {
		panic(field)
	}
}

func LookupUserModelColumns(paths []string) (cols []string) {
	for _, p := range paths {
		cols = append(cols, LookupUserModelColumn(p))
	}
	return
}
