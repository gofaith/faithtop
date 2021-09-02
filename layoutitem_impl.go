//go:build impl

package faithtop

import "github.com/therecipe/qt/widgets"

type LayoutItemImpl struct {
	layoutItem *widgets.QLayoutItem
}

func layoutItemImplFrom(parent *widgets.QLayoutItem) *LayoutItemImpl {
	return &LayoutItemImpl{
		layoutItem: parent,
	}
}

func (l *LayoutItemImpl) LayoutItem() ILayoutItem {
	return l
}
