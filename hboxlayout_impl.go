//go:build impl

package faithtop

import "github.com/therecipe/qt/widgets"

type HBoxLayoutImpl struct {
	*BoxLayoutImpl
	layout *widgets.QHBoxLayout
}

func init() {
	newHBoxLayoutImpl = func() IHBoxLayout {
		v := &HBoxLayoutImpl{
			layout: widgets.NewQHBoxLayout(),
		}
		v.BoxLayoutImpl = boxLayoutImplFrom(v.layout.QBoxLayout_PTR())
		return v
	}
}
