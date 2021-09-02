package faithtop

type IFormLayout interface {
	ILayout
	Append(rows ...IFormRow) IFormLayout
}

type IFormRow interface {
	appendTo(formLayout IFormLayout)
}

var (
	newFormLayoutImpl func() IFormLayout
	newFormRow        func(label IWidget, str string, field IWidget, layout ILayout) IFormRow
)

func FormLayout() IFormLayout {
	return newFormLayoutImpl()
}

func FormLayout2(rows ...IFormRow) IFormLayout {
	return newFormLayoutImpl().Append(rows...)
}

func FormRow(label, field IWidget) IFormRow {
	return newFormRow(label, "", field, nil)
}

func FormRow2(label IWidget, layout ILayout) IFormRow {
	return newFormRow(label, "", nil, layout)
}

func FormRow3(str string, field IWidget) IFormRow {
	return newFormRow(nil, str, field, nil)
}

func FormRow4(str string, field ILayout) IFormRow {
	return newFormRow(nil, str, nil, field)
}

func FormRow5(widget IWidget) IFormRow {
	return newFormRow(nil, "", widget, nil)
}

func FormRow6(layout ILayout) IFormRow {
	return newFormRow(nil, "", nil, layout)
}
