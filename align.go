package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FAlign struct {
	FBaseView
	v *gtk.Alignment
}

func (v *FAlign) getBaseView() *FBaseView {
	return &v.FBaseView
}
func newAlign(i IView, xa, ya, xs, ys float64) *FAlign {
	f := &FAlign{}
	f.v = gtk.NewAlignment(xa, ya, xs, ys)
	f.view = f.v
	f.widget = &f.v.Widget
	f.v.Add(i.getBaseView().widget)
	setupWidget(f)
	return f
}
func AlignLeft(i IView) *FAlign {
	return newAlign(i, 0, 0, 0, 1)
}
func AlignRight(i IView) *FAlign {
	return newAlign(i, 1, 0, 0, 1)
}
func AlignHCenter(i IView) *FAlign {
	return newAlign(i, 0.5, 0, 0, 1)
}
func AlignTop(i IView) *FAlign {
	return newAlign(i, 0, 0, 1, 0)
}
func AlignBottom(i IView) *FAlign {
	return newAlign(i, 0, 1, 1, 0)
}
func AlignVCenter(i IView) *FAlign {
	return newAlign(i, 0, 0.5, 1, 0)
}
func AlignCenter(i IView) *FAlign {
	return newAlign(i, 0.5, 0.5, 0, 0)
}
func AlignCustom(xAlign, yAlign, xScale, yScale float64, i IView) *FAlign {
	return newAlign(i, xAlign, yAlign, xScale, yScale)
}
