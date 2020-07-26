package faithtop

import (
	"github.com/StevenZack/livedata"
	"github.com/gofaith/walk"
	"github.com/gofaith/walk/declarative"
)

type FMarkup struct {
	FBaseView
	v   *walk.LinkLabel
	dec declarative.LinkLabel
}

func Markup(s string) *FMarkup {
	f := &FMarkup{}
	f.dec = declarative.LinkLabel{
		AssignTo: &f.v,
		Text:     s,
	}
	return f
}

func (f *FMarkup) Markup(s string) *FMarkup {
	if f.v != nil {
		f.v.SetText(s)
		return f
	}
	f.dec.Text = s
	return f
}

func (f *FMarkup) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FMarkup) widget(builder *declarative.Builder) walk.Widget {
	if f.v == nil {
		f.dec.Create(builder)
	}
	return f.v
}

func (f *FMarkup) declarative() declarative.Widget {
	return f.dec
}

func (f *FMarkup) BindMarkup(l *livedata.String)*FMarkup  {
	l.ObserveForever(func(str string){
		f.Markup(str)
	})
	return f
}