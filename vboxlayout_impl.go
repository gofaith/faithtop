//go:build impl

package faithtop

import "github.com/therecipe/qt/widgets"

type VBoxLayoutImpl struct {
	*BoxLayoutImpl
	layout *widgets.QVBoxLayout
}

func init() {
	newVBoxLayoutImpl = func() IVBoxLayout {
		v := &VBoxLayoutImpl{
			layout: widgets.NewQVBoxLayout(),
		}
		v.BoxLayoutImpl = boxLayoutImplFrom(v.layout.QBoxLayout_PTR())
		return v
	}
}
