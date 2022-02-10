package gengorm

func (m *Message) genCRUD() {
	m.genCreate()
	m.genGet()
	m.genList()
	m.genUpdate()
	m.genDelete()
}

func (m *Message) genCreate() {
	ctx := m.file.out.QualifiedGoIdent(pkgCtx)
	m.P("func Create", m.ModelName(), "(ctx ", ctx, ") {}")
	m.P()
}

func (m *Message) genGet() {
	ctx := m.file.out.QualifiedGoIdent(pkgCtx)
	m.P("func Get", m.ModelName(), "(ctx ", ctx, ") {}")
	m.P()
}

func (m *Message) genList() {
	ctx := m.file.out.QualifiedGoIdent(pkgCtx)
	m.P("func List", m.ModelName(), "(ctx ", ctx, ") {}")
	m.P()
}

func (m *Message) genUpdate() {
	ctx := m.file.out.QualifiedGoIdent(pkgCtx)
	m.P("func Update", m.ModelName(), "(ctx ", ctx, ") {}")
	m.P()
}

func (m *Message) genDelete() {
	ctx := m.file.out.QualifiedGoIdent(pkgCtx)
	m.P("func Delete", m.ModelName(), "(ctx ", ctx, ") {}")
	m.P()
}
