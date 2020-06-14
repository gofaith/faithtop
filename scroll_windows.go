package faithtop

import (
	"github.com/gofaith/walk"
	"github.com/gofaith/walk/declarative"
)

type FScroll struct {
	FBaseView
	v   *walk.ScrollView
	dec declarative.ScrollView
}

func newScroll(v, h bool) *FScroll {
	f := &FScroll{}
	f.dec = declarative.ScrollView{
		AssignTo:        &f.v,
		VerticalFixed:   v,
		HorizontalFixed: h,
	}
	if v {
		f.dec.Layout = declarative.VBox{SpacingZero: true, MarginsZero: true}
	} else {
		f.dec.Layout = declarative.HBox{SpacingZero: true, MarginsZero: true}
	}
	return f
}

func Scroll() *FScroll {
	return newScroll(true, true)
}

func VScroll() *FScroll {
	return newScroll(true, false)
}

func HScroll() *FScroll {
	return newScroll(false, true)
}

func (f *FScroll) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FScroll) widget(builder *declarative.Builder) walk.Widget {
	if f.v == nil {
		f.dec.Create(builder)
	}
	return f.v
}

func (f *FScroll) declarative() declarative.Widget {
	return f.dec
}

func (f *FScroll) Assign(v **FScroll) *FScroll {
	*v = f
	return f
}

func (f *FScroll) Append(is ...IView) *FScroll {
	if f.v != nil {
		builder := declarative.NewBuilder(f.v)
		for _, i := range is {
			f.v.Children().Add(i.widget(builder))
		}
		return f
	}

	for _, i := range is {
		f.dec.Children = append(f.dec.Children, i.declarative())
	}
	return f
}
