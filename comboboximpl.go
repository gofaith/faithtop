//go:build impl

package faithtop

import "github.com/therecipe/qt/widgets"

type ComboBoxImpl struct {
	*WidgetImpl
	comboBox *widgets.QComboBox
}

func init() {
	newComboBoxImpl = func() IComboBox {
		v := &ComboBoxImpl{
			comboBox: widgets.NewQComboBox(nil),
		}
		v.WidgetImpl = widgetImplFrom(v.comboBox.QWidget_PTR())
		return v
	}
}

func (c *ComboBoxImpl) Assign(v *IComboBox) IComboBox {
	*v = c
	return c
}

func (c *ComboBoxImpl) Clear() IComboBox {
	c.comboBox.Clear()
	return c
}

func (c *ComboBoxImpl) Data(texts []string) IComboBox {
	c.Clear()
	c.comboBox.AddItems(texts)
	return c
}

func (c *ComboBoxImpl) Editable(b bool) IComboBox {
	c.comboBox.SetEditable(b)
	return c
}

func (c *ComboBoxImpl) OnTextChanged(fn func(text string)) IComboBox {
	c.comboBox.ConnectCurrentTextChanged(fn)
	return c
}

func (c *ComboBoxImpl) OnIndexChanged(fn func(index int)) IComboBox {
	c.comboBox.ConnectCurrentIndexChanged(fn)
	return c
}

func (c *ComboBoxImpl) Current(i int) IComboBox {
	c.comboBox.SetCurrentIndex(i)
	return c
}
