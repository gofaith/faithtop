package faithtop

import (
	"github.com/therecipe/qt/widgets"
)

type FBox struct {
	FBaseView
	layout     *widgets.QBoxLayout
	v          *widgets.QWidget
	isVertical bool
}

func VBox() *FBox {
	f := &FBox{
		v:          widgets.NewQWidget(nil, 0),
		isVertical: true,
	}
	f.layout = widgets.NewQVBoxLayout2(nil).QBoxLayout_PTR()
	f.v.SetLayout(f.layout)

	f.widget = f.v
	return f
}

func HBox() *FBox {
	f := &FBox{
		v: widgets.NewQWidget(nil, 0),
	}
	f.layout = widgets.NewQHBoxLayout2(nil).QBoxLayout_PTR()
	f.v.SetLayout(f.layout)

	f.widget = f.v
	return f
}

//base
func (f *FBox) baseView() *FBaseView {
	return &f.FBaseView
}

// box

func (f *FBox) Append(is ...IView) *FBox {
	if f.isVertical {
		for _, i := range is {
			if i.baseView().expand {
				i.baseView().widget.QWidget_PTR().SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
			}
			f.layout.InsertWidget(-1, i.baseView().widget, 0, 0)
		}
		return f
	}
	for _, i := range is {
		if i.baseView().expand {
			i.baseView().widget.QWidget_PTR().SetSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding)
		}
		f.layout.InsertWidget(-1, i.baseView().widget, 0, 0)
	}
	return f
}

func (f *FBox) OnDragDrop(fn func([]string)) *FBox {
	f.FBaseView.OnDragDrop(fn)
	return f
}
