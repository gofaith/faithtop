package faithtop

import (
	"github.com/therecipe/qt/widgets"
)

type FButton struct {
	FBaseView
	v     *widgets.QPushButton
	click func()
}

func Button() *FButton {
	f := &FButton{
		v: widgets.NewQPushButton(nil),
	}
	
	f.v.ConnectClicked(func(bool) {
		if f.click != nil {
			f.click()
		}
	})
	return f
}

// base
func (f *FButton) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FButton) Assign(v **FButton) *FButton {
	*v = f
	return f
}

func (f *FButton) Expand() *FButton {
	f.FbaseView().expand(true)
	return f
}

func (f *FButton) Size(w, h int) *FButton {
	f.FBaseView.Size(w, h)
	return f
}

func (f *FButton) Invisible() *FButton {
	f.FBaseView.Invisible()
	return f
}
func (f *FButton) Visible() *FButton {
	f.FBaseView.Visible()
	return f
}
// button
func (f *FButton) Text(s string) *FButton {
	f.v.SetText(s)
	return f
}

func (f *FButton) OnClick(fn func()) *FButton {
	f.click = fn
	return f
}

func (f *FButton) Disable() *FButton {
	f.v.SetEnabled(false)
	return f
}

func (f *FButton) Enable() *FButton {
	f.v.SetEnabled(true)
	return f
}
