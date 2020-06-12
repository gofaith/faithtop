package faithtop

import (
	"github.com/therecipe/qt/widgets"
)

type FBox struct {
	FBaseView
	v *widgets.QVBoxLayout
	h *widgets.QHBoxLayout
}

func VBox() *FBox {
	f := &FBox{
		v: widgets.NewQVBoxLayout2(nil),
	}
	f.layout = f.v
	return f
}

//base
func (f *FBox) baseView() *FBaseView {
	return &f.FBaseView
}

// box

func (f *FBox) Append(is ...IView) *FBox {
	if f.v != nil {
		for _, i := range is {
			if i.baseView().isLayout() {
				f.v.AddItem(i.baseView().layout)
			} else {
				f.v.AddWidget(i.baseView().widget, 0, 0)
			}
		}
		return f
	}
	return f
}
