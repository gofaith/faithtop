package faithtop

import (
	"github.com/gofaith/walk"
	"github.com/gofaith/walk/declarative"
)

type FVSpace struct {
	FBaseView
	v   *walk.Spacer
	dec declarative.VSpacer
}

func VSpace() *FVSpace {
	f := &FVSpace{}
	f.dec = declarative.VSpacer{
		AssignTo: &f.v,
	}
	return f
}

func (f *FVSpace) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FVSpace) widget(*declarative.Builder) walk.Widget {
	return f.v
}

func (f *FVSpace) declarative() declarative.Widget {
	return f.dec
}

// hspace
type FHSpace struct {
	FBaseView
	v   *walk.Spacer
	dec declarative.HSpacer
}

func HSpace() *FHSpace {
	f := &FHSpace{}
	f.dec = declarative.HSpacer{
		AssignTo: &f.v,
	}
	return f
}

func (f *FHSpace) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FHSpace) widget(builder *declarative.Builder) walk.Widget {
	return f.v
}
func (f *FHSpace) declarative() declarative.Widget {
	return f.dec
}
