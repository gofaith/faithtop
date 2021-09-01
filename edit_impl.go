//go:build impl

package faithtop

import "github.com/therecipe/qt/widgets"

type EditImpl struct {
	*WidgetImpl
	edit *widgets.QLineEdit
}

func init() {
	newEditImpl = func() IEdit {
		v := &EditImpl{
			edit: widgets.NewQLineEdit(nil),
		}
		v.WidgetImpl = widgetImplFrom(v.edit.QWidget_PTR())
		return v
	}
}

func (l *EditImpl) Assign(v *IEdit) IEdit {
	*v = l
	return l
}

func (l *EditImpl) Text(s string) IEdit {
	l.edit.SetText(s)
	return l
}

func (l *EditImpl) GetText() string {
	return l.edit.Text()
}
