package gengorm

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

func (m *Message) genCRUD() {
	if !m.opts.Crud {
		return
	}
	m.genWithDBType()
	m.genCreate()
	m.genGet()
	m.genList()
	m.genUpdate()
	m.genPatch()
	m.genDelete()
}

func (m *Message) genWithDBType() {
	m.P("type ", m.withDBTypeName(), " struct {")
	m.P("x *", m.ProtoName())
	m.P("db *", m.identGormDB())
	m.P("}")
	m.P()
	m.P("func (x *", m.ProtoName(), ") WithDB(db *", m.identGormDB(), ") ", m.withDBTypeName(), " {")
	m.P("return ", m.withDBTypeName(), "{x: x, db: db}")
	m.P("}") // func
	m.P()
}

func (m *Message) genCreate() {
	m.P("func (c ", m.withDBTypeName(), ") Create(ctx ", m.identCtx(), ", opts ...", m.identCreateOption(), ") (*", m.ProtoName(), ", error) {")
	m.P("if c.x == nil {")
	m.P("return nil, nil")
	m.P("}")

	// proto -> GORM
	m.P("m, err := c.x.AsModel()")
	m.P("if err != nil {")
	m.P("return nil, err")
	m.P("}") // if

	// INSERT INTO ...
	m.P("db := c.db.WithContext(ctx)")
	m.P("if err := db.Create(m).Error; err != nil {")
	m.P("return nil, err")
	m.P("}") // if

	// GORM -> proto
	m.P("if y, err := m.AsProto(); err != nil {")
	m.P("return nil, err")
	m.P("} else {")
	m.P("return y, nil")
	m.P("}")

	m.P("}") // func
	m.P()
}

func (m *Message) genGet() {
	m.P("func (c ", m.withDBTypeName(), ") Get(ctx ", m.identCtx(), ", opts ...", m.identGetOption(), ") (*", m.ProtoName(), ", error) {")
	m.P("if c.x == nil {")
	m.P("return nil, nil")
	m.P("}")

	pk := m.primaryKey()
	if pk == nil {
		err := fmt.Errorf("no primary key on message %s", m.ProtoName())
		panic(err)
	}

	if pk.types.Pointer {
		m.P("if c.x.", pk.Name(), " == nil {")
		m.P("return nil, ", m.identErrorf(), "(\"nil primary key\")")
		m.P("}")
	} else {
		m.P("var zero ", pk.types.String())
		m.P("if c.x.", pk.Name(), " == zero {")
		m.P("return nil, ", m.identErrorf(), "(\"empty primary key\")")
		m.P("}")
	}

	// proto -> GORM
	m.P("m, err := c.x.AsModel()")
	m.P("if err != nil {")
	m.P("return nil, err")
	m.P("}")

	// SELECT ... WHERE ...
	m.P("n := ", m.ModelName(), "{}")
	m.P("db := c.db.WithContext(ctx)")
	m.P("if err := db.Where(m).First(&n).Error; err != nil {")
	m.P("return nil, err")
	m.P("}")

	// GORM -> proto
	m.P("if y, err := n.AsProto(); err != nil {")
	m.P("return nil, err")
	m.P("} else {")
	m.P("return y, nil")
	m.P("}")

	m.P("}") // func
	m.P()
}

func (m *Message) genList() {
	m.P("func (c ", m.withDBTypeName(), ") List(ctx ", m.identCtx(), ", opts ...", m.identListOption(), ") ([]*", m.ProtoName(), ", error) {")
	m.P("if c.x == nil {")
	m.P("return nil, nil")
	m.P("}")

	// SELECT * FROM ...
	m.P("var ms []", m.ModelName())
	m.P("db := c.db.WithContext(ctx)")
	m.P("if err := db.Find(&ms).Error; err != nil {")
	m.P("return nil, err")
	m.P("}")

	// []GORM -> []proto
	m.P("xs := make([]*", m.ProtoName(), ", 0, len(ms))")
	m.P("for _, m := range ms {")
	m.P("if x, err := m.AsProto(); err != nil {")
	m.P("return nil, err")
	m.P("} else {")
	m.P("xs = append(xs, x)")
	m.P("}") // if
	m.P("}") // for

	m.P("return xs, nil")
	m.P("}") // func
	m.P()
}

func (m *Message) genUpdate() {
	m.P("func (c ", m.withDBTypeName(), ") Update(ctx ", m.identCtx(), ", opts ...", m.identUpdateOption(), ") (*", m.ProtoName(), ", error) {")
	m.P("if c.x == nil {")
	m.P("return nil, nil")
	m.P("}")

	// proto -> GORM
	m.P("m, err := c.x.AsModel()")
	m.P("if err != nil {")
	m.P("return nil, err")
	m.P("}")

	// UPDATE ... SET ...
	m.P("db := c.db.WithContext(ctx)")
	m.P("if err := db.Save(m).Error; err != nil {")
	m.P("return nil, err")
	m.P("}") // if

	// GORM -> proto
	m.P("if y, err := m.AsProto(); err != nil {")
	m.P("return nil, err")
	m.P("} else {")
	m.P("return y, nil")
	m.P("}")

	m.P("}") // func
	m.P()
}

func (m *Message) genPatch() {
	m.P("func (c ", m.withDBTypeName(), ") "+
		"Patch(ctx ", m.identCtx(), ", mask *", m.identFieldMask(), ", opts ...", m.identPatchOption(), ") (*", m.ProtoName(), ", error) {")
	m.P("if c.x == nil {")
	m.P("return nil, nil")
	m.P("}")

	m.P("if mask == nil {")
	m.P("return c.Update(ctx)")
	m.P("}")

	m.P("if !mask.IsValid(c.x) {")
	m.P("return nil, ", m.identErrorf(), "(\"invalid field mask\")")
	m.P("}")

	m.P("paths := mask.Paths")
	m.P("if len(paths) == 0 {")
	m.P("return c.Update(ctx)")
	m.P("}")

	pk := m.primaryKey()
	if pk == nil {
		err := fmt.Errorf("no primary key on message %s", m.ProtoName())
		panic(err)
	}

	if pk.types.Pointer {
		m.P("if c.x.", pk.Name(), " == nil {")
		m.P("return nil, ", m.identErrorf(), "(\"nil primary key\")")
		m.P("}")
	} else {
		m.P("var zero ", pk.types.String())
		m.P("if c.x.", pk.Name(), " == zero {")
		m.P("return nil, ", m.identErrorf(), "(\"empty primary key\")")
		m.P("}")
	}

	m.P("var cols []string")
	m.P("for _, path := range paths {")
	m.P("switch path {")
	for _, field := range m.fields {
		m.P("case \"", field.proto.Desc.Name(), "\":")
		m.P("cols = append(cols, \"", field.Name(), "\")")
		// if col := field.opts.Column; col != "" {
		// 	m.P("cols = append(cols, \"", col, "\")")
		// } else {
		// }
	}
	m.P("}") // switch
	m.P("}") // for

	// proto -> GORM
	m.P("m, err := c.x.AsModel()")
	m.P("if err != nil {")
	m.P("return nil, err")
	m.P("}")

	m.P("target := ", m.ModelName(), "{", pk.Name(), ": m.", pk.Name(), "}")

	// UPDATE ... SET ...
	m.P("db := c.db.WithContext(ctx)")
	m.P("if err := db.Model(&target).Select(cols).Updates(m).Error; err != nil {")
	m.P("return nil, err")
	m.P("}") // if

	// GORM -> proto
	m.P("if y, err := m.AsProto(); err != nil {")
	m.P("return nil, err")
	m.P("} else {")
	m.P("return y, nil")
	m.P("}")

	m.P("}") // func
	m.P()
}

// TODO: Soft delete, expiration?
func (m *Message) genDelete() {
	m.P("func (c ", m.withDBTypeName(), ") Delete(ctx ", m.identCtx(), ", opts ...", m.identDeleteOption(), ") error {")
	m.P("if c.x == nil {")
	m.P("return nil")
	m.P("}")

	pk := m.primaryKey()
	if pk == nil {
		err := fmt.Errorf("no primary key on message %s", m.ProtoName())
		panic(err)
	}

	if pk.types.Pointer {
		m.P("if c.x.", pk.Name(), " == nil {")
		m.P("return ", m.identErrorf(), "(\"nil primary key\")")
		m.P("}")
	} else {
		m.P("var zero ", pk.types.String())
		m.P("if c.x.", pk.Name(), " == zero {")
		m.P("return ", m.identErrorf(), "(\"empty primary key\")")
		m.P("}")
	}

	// proto -> GORM
	m.P("m, err := c.x.AsModel()")
	m.P("if err != nil {")
	m.P("return err")
	m.P("}")

	// DELETE FROM ...
	m.P("db := c.db.WithContext(ctx)")
	m.P("if err := db.Where(m).Delete(&", m.ModelName(), "{}).Error; err != nil {")
	m.P("return err")
	m.P("}")

	m.P("return nil")
	m.P("}") // func
	m.P()
}

func (m *Message) primaryKey() *Field {
	for _, field := range m.fields {
		if field.opts.PrimaryKey {
			return field
		}
	}
	return nil
}

func (m Message) withDBTypeName() string {
	return m.ProtoName() + "WithDB"
}

func (m *Message) identGormDB() string {
	return m.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "DB",
		GoImportPath: "gorm.io/gorm",
	})
}

func (m *Message) identCtx() string {
	return m.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Context",
		GoImportPath: "context",
	})
}

func (m *Message) identErrorf() string {
	return m.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Errorf",
		GoImportPath: "fmt",
	})
}

func (m *Message) identFieldMask() string {
	return m.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "FieldMask",
		GoImportPath: "google.golang.org/protobuf/types/known/fieldmaskpb",
	})
}

func (m *Message) identCreateOption() string { return m.identGenGorm("CreateOption") }
func (m *Message) identGetOption() string    { return m.identGenGorm("GetOption") }
func (m *Message) identListOption() string   { return m.identGenGorm("ListOption") }
func (m *Message) identUpdateOption() string { return m.identGenGorm("UpdateOption") }
func (m *Message) identPatchOption() string  { return m.identGenGorm("PatchOption") }
func (m *Message) identDeleteOption() string { return m.identGenGorm("DeleteOption") }

func (m *Message) identGenGorm(goName string) string {
	return m.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       goName,
		GoImportPath: "github.com/complex64/protoc-gen-gorm/pkg/gengorm",
	})
}
