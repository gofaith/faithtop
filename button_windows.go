package faithtop

import (
	"github.com/gofaith/walk"
	"github.com/gofaith/walk/declarative"
)

type FButton struct {
	FBaseView
	v     *walk.PushButton
	dec   declarative.PushButton
	click func()
}

func Button() *FButton {
	f := &FButton{}
	f.dec = declarative.PushButton{
		AssignTo: &f.v,
		OnClicked: func() {
			if f.click != nil {
				f.click()
			}
		},
		AlwaysConsumeSpace: true,
	}
	return f
}

func (f *FButton) baseView() *FBaseView {
	return &f.FBaseView
}
func (f *FButton) widget(builder *declarative.Builder) walk.Widget {
	if f.v == nil {
		f.dec.Create(builder)
	}
	return f.v
}
func (f *FButton) declarative() declarative.Widget {
	return f.dec
}

func (f *FButton) Assign(v **FButton) *FButton {
	*v = f
	return f
}

func (f *FButton) Expand() *FButton {
	f.expand = true
	return f
}

func (f *FButton) Text(text string) *FButton {
	if f.v != nil {
		f.v.SetText(text)
		return f
	}
	f.dec.Text = text
	return f
}

func (f *FButton) OnClick(fn func()) *FButton {
	if f.v != nil {
		f.v.Clicked().Once(fn)
		return f
	}
	f.click = fn
	return f
}
