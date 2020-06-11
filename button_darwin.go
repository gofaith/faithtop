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
	f.view = f.v
	f.v.ConnectClick(func() {
		if f.click != nil {
			f.click()
		}
	})
	return f
}

func (f *FButton) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FButton) Text(s string) *FButton {
	f.v.SetText(s)
	return f
}
