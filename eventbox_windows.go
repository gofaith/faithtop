package faithtop

import (
	"github.com/gofaith/walk"
	"github.com/gofaith/walk/declarative"
)

type FEventBox struct {
	FBaseView
	v   *walk.Composite
	dec declarative.Composite
}

func EventBox(i IView) *FEventBox {
	f := &FEventBox{}
	f.dec = declarative.Composite{
		AssignTo: &f.v,
		Layout:   declarative.VBox{SpacingZero: true, MarginsZero: true},
		Children: []declarative.Widget{
			i.declarative(),
		},
	}
	return f
}

func (f *FEventBox) OnClick(fn func()) *FEventBox {
	f.dec.OnMouseUp = func(x, y int, button walk.MouseButton) {
		if fn != nil {
			fn()
		}
	}
	return f
}

func (f *FEventBox) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FEventBox) widget(builder *declarative.Builder) walk.Widget {
	if f.v == nil {
		f.dec.Create(builder)
	}
	return f.v
}

func (f *FEventBox) declarative() declarative.Widget {
	return f.dec
}

func (f *FEventBox) Assign(v **FEventBox) *FEventBox {
	*v = f
	return f
}
