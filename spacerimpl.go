//go:build impl

package faithtop

import "github.com/therecipe/qt/widgets"

type SpacerImpl struct {
	*LayoutItemImpl
	spacer *widgets.QSpacerItem
}

func init() {
	newHSpacerImpl = func() ISpacer {
		v := &SpacerImpl{
			spacer: widgets.NewQSpacerItem(0, 0, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Fixed),
		}
		v.LayoutItemImpl = layoutItemImplFrom(v.spacer.QLayoutItem_PTR())
		return v
	}
	newVSpacerImpl = func() ISpacer {
		v := &SpacerImpl{
			spacer: widgets.NewQSpacerItem(0, 0, widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Expanding),
		}
		v.LayoutItemImpl = layoutItemImplFrom(v.spacer.QLayoutItem_PTR())
		return v
	}
}
