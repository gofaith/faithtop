package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type FScroll struct {
	FBaseView
	v   *widgets.QScrollArea
	box *FBox
}

func newScroll() *FScroll {
	f := &FScroll{
		v: widgets.NewQScrollArea(nil),
	}
	f.widget = f.v
	return f
}

func Scroll(child IView) *FScroll {
	f := newScroll()
	f.v.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAsNeeded)
	f.v.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAsNeeded)
	f.v.SetWidget(child.baseView().widget)
	return f
}

func VScroll() *FScroll {
	f := newScroll()
	f.v.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAsNeeded)
	f.v.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	f.box = VBox()
	return f
}

func HScroll() *FScroll {
	f := newScroll()
	f.v.SetHorizontalScrollBarPolicy(core.Qt__ScrollBarAsNeeded)
	f.v.SetVerticalScrollBarPolicy(core.Qt__ScrollBarAlwaysOff)
	f.box = HBox()
	return f
}

//base
func (f *FScroll) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FScroll) Assign(v **FScroll) *FScroll {
	*v = f
	return f
}

// scroll

func (f *FScroll) Append(is ...IView) *FScroll {
	f.box.Append(is...)
	central := widgets.NewQWidget(nil, 0)
	central.SetLayout(f.box.layout)
	f.v.SetWidget(central)
	return f
}
