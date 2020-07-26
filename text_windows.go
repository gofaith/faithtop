package faithtop

import (
	"github.com/StevenZack/livedata"
	"github.com/gofaith/walk"
	"github.com/gofaith/walk/declarative"
)

type FText struct {
	FBaseView
	v   *walk.Label
	dec declarative.Label
}

func Text(s string) *FText {
	f := &FText{}
	f.dec = declarative.Label{
		AssignTo: &f.v,
		Text:     s,
	}
	return f
}

func (f *FText) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FText) widget(builder *declarative.Builder) walk.Widget {
	if f.v == nil {
		f.dec.Create(builder)
	}
	return f.v
}

func (f *FText) declarative() declarative.Widget {
	return f.dec
}

func (f *FText) Assign(v **FText) *FText {
	*v = f
	return f
}

func (f *FText) Expand() *FText {
	f.expand = true
	return f
}

func (f *FText) Text(text string) *FText {
	if f.v != nil {
		f.v.SetText(text)
		return f
	}
	f.dec.Text = text
	return f
}

func (f *FText) BindText(l *livedata.String) *FText {
	l.ObserveForever(func(str string) {
		f.Text(str)
	})
	return f
}
