//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type BoxLayoutImpl struct {
	*LayoutImpl
	boxLayout        *widgets.QBoxLayout
	defaultAlignment AlignmentFlag
}

func boxLayoutImplFrom(parent *widgets.QBoxLayout) *BoxLayoutImpl {
	v := &BoxLayoutImpl{
		boxLayout:        parent,
		defaultAlignment: AlignLeading,
	}
	v.LayoutImpl = layoutImplFrom(v.boxLayout.QLayout_PTR())
	return v
}

func (l *BoxLayoutImpl) Align(align AlignmentFlag) IBoxLayout {
	l.defaultAlignment = align
	return l
}

func (l *BoxLayoutImpl) AppendWidgets(children ...IWidget) IBoxLayout {
	for _, child := range children {
		v := child.Widget().(*WidgetImpl)
		align := v.alignmentFlag
		if align == 0 {
			align = l.defaultAlignment
		}
		l.boxLayout.AddWidget(v.widget, 0, core.Qt__AlignmentFlag(align))
	}
	return l
}

func (l *BoxLayoutImpl) AppendItems(items ...ILayoutItem) IBoxLayout {
	for _, item := range items {
		l.boxLayout.AddItem(item.LayoutItem().(*LayoutItemImpl).layoutItem)
	}
	return l
}
