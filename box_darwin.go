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
	f.v.SetSpacing(0)
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
				if i.baseView().expand {
					i.baseView().widget.QWidget_PTR().SetSizePolicy2(widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
				}
				f.v.AddWidget(i.baseView().widget, 0, i.baseView().align)
			}
		}
		return f
	}
	return f
}
