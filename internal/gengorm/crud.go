package gengorm

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
)

func (m *Message) genCRUD() {
	if !m.opts.Crud {
		return
	}
	m.genOptionTypes()
	m.genWithDBType()
	m.genCreate()
	m.genGet()
	m.genList()
	m.genUpdate()
	m.genPatch()
	m.genDelete()
	m.genGetOptions()
	m.genListOptions()
	m.genColForPath()
	m.genColsForPaths()
}

func (m *Message) genWithDBType() {
	m.P("type ", m.typeNameWithDB(), " struct {")
	m.P("x *", m.ProtoName())
	m.P("db *", m.identGormDB())
	m.P("}")
	m.P()
	m.P("func (x *", m.ProtoName(), ") WithDB(db *", m.identGormDB(), ") ", m.typeNameWithDB(), " {")
	m.P("return ", m.typeNameWithDB(), "{x: x, db: db}")
	m.P("}") // func
	m.P()
}

func (m *Message) genOptionTypes() {
	m.P("type ", m.typeNameGetOption(), " func(tx *gorm.DB) *gorm.DB")
	m.P("type ", m.typeNameListOption(), " func(tx *gorm.DB) *gorm.DB")
	m.P()
}

func (m *Message) genCreate() {
	m.P("func (c ", m.typeNameWithDB(), ") Create(ctx ", m.identCtx(), ") (*", m.ProtoName(), ", error) {")
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
	m.P("func (c ", m.typeNameWithDB(), ") Get(ctx ", m.identCtx(), ", opts ...", m.typeNameGetOption(), ") (*", m.ProtoName(), ", error) {")
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

	m.P("db := c.db.WithContext(ctx)")
	m.P("for _, opt := range opts {")
	m.P("db = opt(db)")
	m.P("}")

	// SELECT ... WHERE ...
	m.P("out := ", m.ModelName(), "{}")
	m.P("if err := db.Where(m).First(&out).Error; err != nil {")
	m.P("return nil, err")
	m.P("}")

	// GORM -> proto
	m.P("if y, err := out.AsProto(); err != nil {")
	m.P("return nil, err")
	m.P("} else {")
	m.P("return y, nil")
	m.P("}")

	m.P("}") // func
	m.P()
}

func (m *Message) genList() {
	m.P("func (c ", m.typeNameWithDB(), ") List(ctx ", m.identCtx(), ", opts ...", m.typeNameListOption(), ") ([]*", m.ProtoName(), ", error) {")
	m.P("if c.x == nil {")
	m.P("return nil, nil")
	m.P("}")

	m.P("db := c.db.WithContext(ctx)")
	m.P("for _, opt := range opts {")
	m.P("db = opt(db)")
	m.P("}")

	// SELECT * FROM ...
	m.P("var ms []", m.ModelName())
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
	m.P("func (c ", m.typeNameWithDB(), ") Update(ctx ", m.identCtx(), ") (*", m.ProtoName(), ", error) {")
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

	m.P("return c.Get(ctx)")
	m.P("}") // func
	m.P()
}

func (m *Message) genPatch() {
	m.P("func (c ", m.typeNameWithDB(), ") "+
		"Patch(ctx ", m.identCtx(), ", mask *", m.identFieldMask(), ") error {")
	m.P("if c.x == nil {")
	m.P("return nil")
	m.P("}")

	m.P("if mask == nil {")
	m.P("_, err := c.Update(ctx)")
	m.P("return err")
	m.P("}")

	m.P("if !mask.IsValid(c.x) {")
	m.P("return ", m.identErrorf(), "(\"invalid field mask\")")
	m.P("}")

	m.P("paths := mask.Paths")
	m.P("if len(paths) == 0 {")
	m.P("_, err := c.Update(ctx)")
	m.P("return err")
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

	m.P("target := ", m.ModelName(), "{", pk.Name(), ": m.", pk.Name(), "}")

	// UPDATE ... SET ...
	m.P("cols := ", m.funcLookupCols(), "(paths)")
	m.P("db := c.db.WithContext(ctx)")
	m.P("if err := db.Model(&target).Select(cols).Updates(m).Error; err != nil {")
	m.P("return err")
	m.P("}") // if

	m.P("return nil")
	m.P("}") // func
	m.P()
}

// TODO: Soft delete, expiration?
func (m *Message) genDelete() {
	m.P("func (c ", m.typeNameWithDB(), ") Delete(ctx ", m.identCtx(), ") error {")
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

func (m *Message) genGetOptions() {
	m.genWithGetFieldMask()
}

func (m *Message) genWithGetFieldMask() {
	m.P("func With", m.proto.GoIdent.GoName, "GetFieldMask(mask *", m.identFieldMask(), ") ", m.typeNameGetOption(), " {")

	m.P("return func(tx *gorm.DB) *gorm.DB {")
	m.P("cols := ", m.funcLookupCols(), "(mask.Paths)")
	m.P("tx = tx.Select(cols)")
	m.P("return tx")
	m.P("}") // inner func

	m.P("}") // func
	m.P()
}

func (m *Message) genListOptions() {
	m.genWithFilter()
	m.genWithListFieldMask()
	m.genWithOrder()
	m.genWithPagination()
}

func (m *Message) genWithFilter() {
	m.P("func With", m.proto.GoIdent.GoName, "ListFilter(filter string) ", m.typeNameListOption(), " {")

	m.P("return func(tx *gorm.DB) *gorm.DB {")
	// TODO
	m.P("return tx")
	m.P("}") // inner func

	m.P("}") // func
	m.P()
}

func (m *Message) genWithListFieldMask() {
	m.P("func With", m.proto.GoIdent.GoName, "ListFieldMask(mask *", m.identFieldMask(), ") ", m.typeNameListOption(), " {")

	m.P("return func(tx *gorm.DB) *gorm.DB {")
	m.P("cols := ", m.funcLookupCols(), "(mask.Paths)")
	m.P("tx = tx.Select(cols)")
	m.P("return tx")
	m.P("}") // inner func

	m.P("}") // func
	m.P()
}

func (m *Message) genWithOrder() {
	m.P("func With", m.proto.GoIdent.GoName, "ListOrder(orderBy string) ", m.typeNameListOption(), " {")

	m.P("return func(tx *gorm.DB) *gorm.DB {")
	// TODO
	m.P("return tx")
	m.P("}") // inner func

	m.P("}") // func
	m.P()
}

func (m *Message) genWithPagination() {
	m.P("func With", m.proto.GoIdent.GoName, "ListPagination(page, size int) ", m.typeNameListOption(), " {")
	m.P("return func(tx *gorm.DB) *gorm.DB {")
	m.P("tx = tx.Limit(size)")
	m.P("tx = tx.Offset((page-1)*size)")
	m.P("return tx")
	m.P("}") // inner func
	m.P("}") // func
	m.P()
}

func (m *Message) genColForPath() {
	m.P("func ", m.funcLookupCol(), "(path string) string {")
	m.P("switch path {")
	for _, field := range m.fields {
		m.P("case \"", field.proto.Desc.Name(), "\":")
		if col := field.opts.Column; col != "" {
			m.P("return \"", col, "\"")
		} else {
			m.P("return \"", field.Name(), "\"")
		}
	}
	m.P("}") // switch
	m.P("panic(path)")
	m.P("}") // func
	m.P()
}

func (m *Message) genColsForPaths() {
	m.P("func ", m.funcLookupCols(), "(paths []string) (cols []string) {")
	m.P("for _, p := range paths {")
	m.P("cols = append(cols, ", m.funcLookupCol(), "(p))")
	m.P("}") // for
	m.P("return")
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

func (m Message) funcLookupCol() string {
	return "Lookup" + m.ModelName() + "Column"
}

func (m Message) funcLookupCols() string {
	return "Lookup" + m.ModelName() + "Columns"
}

func (m Message) typeNameGetOption() string {
	return m.proto.GoIdent.GoName + "GetOption"
}

func (m Message) typeNameListOption() string {
	return m.proto.GoIdent.GoName + "ListOption"
}

func (m Message) typeNameWithDB() string {
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

func (m *Message) identAipGo(pkg, goName string) string {
	return m.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       goName,
		GoImportPath: protogen.GoImportPath("go.einride.tech/aip/" + pkg),
	})
}
