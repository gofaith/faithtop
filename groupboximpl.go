//go:build impl

package faithtop

import "github.com/therecipe/qt/widgets"

type GroupBoxImpl struct {
	*WidgetImpl
	groupBox *widgets.QGroupBox
}

func init() {
	newGroupBoxImpl = func() IGroupBox {
		v := &GroupBoxImpl{
			groupBox: widgets.NewQGroupBox(nil),
		}
		v.WidgetImpl = widgetImplFrom(v.groupBox.QWidget_PTR())
		return v
	}
}

func (g *GroupBoxImpl) Title(title string) IGroupBox {
	g.groupBox.SetTitle(title)
	return g
}

func (g *GroupBoxImpl) Checkable(b bool) IGroupBox {
	g.groupBox.SetCheckable(b)
	return g
}

func (g *GroupBoxImpl) Checked(b bool) IGroupBox {
	g.groupBox.SetChecked(b)
	return g
}

func (g *GroupBoxImpl) IsChecked() bool {
	return g.groupBox.IsChecked()
}

func (g *GroupBoxImpl) OnToggled(fn func(self IGroupBox, b bool)) IGroupBox {
	g.groupBox.ConnectToggled(func(on bool) {
		fn(g, on)
	})
	return g
}

func (g *GroupBoxImpl) TitleAlign(align AlignmentFlag) IGroupBox {
	g.groupBox.SetAlignment(int(align))
	return g
}

func (g *GroupBoxImpl) Flat(b bool) IGroupBox {
	g.groupBox.SetFlat(b)
	return g
}
