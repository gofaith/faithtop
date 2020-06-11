package faithtop

import (
	"github.com/therecipe/qt/widgets"
)

type FWindow struct {
	v         *widgets.QMainWindow
	showAfter bool
}

func Win() *FWindow {
	v := widgets.NewQMainWindow(nil, 0)
	f := &FWindow{v: v}
	f.Size(230, 130)
	return f
}

func (f *FWindow) Size(w, h int) *FWindow {
	f.v.SetMinimumSize2(w, h)
	return f
}

func (f *FWindow) Show() *FWindow {
	f.v.Show()
	return f
}

func (f *FWindow) Close() *FWindow {
	f.v.Close()
	return f
}

func (f *FWindow) Hide() *FWindow {
	f.v.Hide()
	return f
}

func (f *FWindow) IsVisible() bool {
	return f.v.IsVisible()
}

func (f *FWindow) DeferShow() *FWindow {
	f.showAfter = true
	return f
}

func (f *FWindow) Add(i IView) *FWindow {
	v := i.baseView().view
	f.v.SetCentralWidget(v)
	if f.showAfter {
		f.Show()
	}
	return f
}
