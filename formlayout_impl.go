//go:build impl

package faithtop

import (
	"fmt"

	"github.com/therecipe/qt/widgets"
)

type (
	FormLayoutImpl struct {
		*LayoutImpl
		layout *widgets.QFormLayout
	}
	FormRowImpl struct {
		label  IWidget
		str    string
		field  IWidget
		layout ILayout
	}
)

func init() {
	newFormLayoutImpl = func() IFormLayout {
		v := &FormLayoutImpl{
			layout: widgets.NewQFormLayout(nil),
		}
		v.LayoutImpl = layoutImplFrom(v.layout.QLayout_PTR())
		return v
	}
	newFormRow = func(label IWidget, str string, field IWidget, layout ILayout) IFormRow {
		return &FormRowImpl{
			label:  label,
			str:    str,
			field:  field,
			layout: layout,
		}
	}
}

func (f *FormRowImpl) appendTo(formLayout IFormLayout) {
	parent := formLayout.(*FormLayoutImpl).layout

	if f.label != nil && f.field != nil {
		parent.AddRow(f.label.getWidget().(*WidgetImpl).widget, f.field.getWidget().(*WidgetImpl).widget)
		return
	}
	if f.label != nil && f.layout != nil {
		parent.AddRow2(f.label.getWidget().(*WidgetImpl).widget, f.layout.getLayout().(*LayoutImpl).layout)
		return
	}
	if f.str != "" && f.field != nil {
		parent.AddRow3(f.str, f.field.getWidget().(*WidgetImpl).widget)
		return
	}
	if f.str != "" && f.layout != nil {
		parent.AddRow4(f.str, f.layout.getLayout().(*LayoutImpl).layout)
		return
	}
	if f.field != nil {
		parent.AddRow5(f.field.getWidget().(*WidgetImpl).widget)
		return
	}
	if f.layout != nil {
		parent.AddRow6(f.layout.getLayout().(*LayoutImpl).layout)
		return
	}
	panic(fmt.Sprint("invalid formLayout row parameter:", f.label != nil, f.str, f.field != nil, f.layout != nil))
}

func (f *FormLayoutImpl) Append(rows ...IFormRow) IFormLayout {
	for _, row := range rows {
		row.appendTo(f)
	}
	return f
}
