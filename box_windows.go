package faithtop

import (
	"github.com/gofaith/walk"
	"github.com/gofaith/walk/declarative"
)

type FBox struct {
	FBaseView
	v   *walk.Composite
	dec declarative.Composite
}

func VBox() *FBox {
	f := &FBox{}
	f.dec = declarative.Composite{
		AssignTo: &f.v,
		Layout:   declarative.VBox{SpacingZero: true, MarginsZero: true},
	}
	return f
}

func HBox() *FBox {
	f := &FBox{}
	f.dec = declarative.Composite{
		AssignTo: &f.v,
		Layout:   declarative.HBox{SpacingZero: true, MarginsZero: true},
	}
	return f
}

func (f *FBox) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FBox) widget() walk.Widget {
	return f.v
}
func (f *FBox) declarative() declarative.Widget {
	return f.dec
}

func (f *FBox) Assign(v **FBox) *FBox {
	*v = f
	return f
}

func (f *FBox) Append(is ...IView) *FBox {
	for _, i := range is {
		f.dec.Children = append(f.dec.Children, i.declarative())
	}
	return f
}
