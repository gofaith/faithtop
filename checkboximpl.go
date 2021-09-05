//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type CheckBoxImpl struct {
	*WidgetImpl
	checkbox *widgets.QCheckBox
}

func init() {
	newCheckBoxImpl = func() ICheckBox {
		v := &CheckBoxImpl{
			checkbox: widgets.NewQCheckBox(nil),
		}
		v.WidgetImpl = widgetImplFrom(v.checkbox.QWidget_PTR())
		return v
	}
}

func (c *CheckBoxImpl) Checked(b bool) ICheckBox {
	c.checkbox.SetChecked(b)
	return c
}

func (c *CheckBoxImpl) Text(s string) ICheckBox {
	c.checkbox.SetText(s)
	return c
}

func (c *CheckBoxImpl) Assign(v *ICheckBox) ICheckBox {
	*v = c
	return c
}

func (c *CheckBoxImpl) IsChecked() bool {
	return c.checkbox.IsChecked()
}

func (c *CheckBoxImpl) OnChanged(fn func(b bool)) ICheckBox {
	c.checkbox.ConnectStateChanged(func(state int) {
		if fn != nil {
			fn(state == int(core.Qt__Checked))
		}
	})
	return c
}
