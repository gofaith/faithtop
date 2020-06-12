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
	f.widget = f.v
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
	f.FBaseView.Expand(true)
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
