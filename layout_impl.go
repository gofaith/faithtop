//go:build impl

package faithtop

import "github.com/therecipe/qt/widgets"

type LayoutImpl struct {
	layout *widgets.QLayout
}

func (l *LayoutImpl) Layout() ILayout {
	return l
}

func layoutImplFrom(parent *widgets.QLayout) *LayoutImpl {
	return &LayoutImpl{
		layout: parent,
	}
}
