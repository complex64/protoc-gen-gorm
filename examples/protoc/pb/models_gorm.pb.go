// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm 2.0.0
// 	protoc          v3.19.3
// source: models.proto

package pb

import (
	_ "github.com/complex64/protoc-gen-gorm/gormpb"
)

// UserModel is the GORM model for pb.User.
type UserModel struct {
	Name string `gorm:"not null;unique;primaryKey"`
}

// ToProto converts a UserModel to its protobuf representation.
func (m *UserModel) ToProto() (*User, error) {
	p := new(User)
	p.Name = m.Name
	return p, nil
}

// ToModel converts a User to its GORM model.
func (p *User) ToModel() (*UserModel, error) {
	m := new(UserModel)
	m.Name = p.Name
	return m, nil
}
