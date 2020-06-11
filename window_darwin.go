package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type FWindow struct {
	w         *widgets.QMainWindow
	showAfter bool
	destroy   func()
}

func Win() *FWindow {
	w := widgets.NewQMainWindow(nil, 0)
	f := &FWindow{w: w}
	f.Size(230, 130)
	f.w.ConnectDestroyed(func(*core.QObject) {
		if f.destroy != nil {
			f.destroy()
		}
	})
	return f
}

func TopWin() *FWindow {
	return Win().Top()
}

func (f *FWindow) Top() *FWindow {
	f.w.SetWindowFlag(core.Qt__WindowStaysOnTopHint, true)
	return f
}

func (f *FWindow) Title(s string) *FWindow {
	f.w.SetWindowTitle(s)
	return f
}

func (f *FWindow) Size(w, h int) *FWindow {
	f.w.SetMinimumSize2(w, h)
	return f
}

func (f *FWindow) Show() *FWindow {
	f.w.Show()
	return f
}

func (f *FWindow) Close() *FWindow {
	f.w.DestroyQMainWindow()
	return f
}

func (f *FWindow) Hide() *FWindow {
	f.w.Hide()
	return f
}

func (f *FWindow) IsVisible() bool {
	return f.w.IsVisible()
}

func (f *FWindow) DeferShow() *FWindow {
	f.showAfter = true
	return f
}

func (f *FWindow) Add(i IView) *FWindow {
	if f.w != nil {
		f.w.SetCentralWidget(i.baseView().view)
	}
	if f.showAfter {
		f.Show()
		f.showAfter = false
	}
	return f
}

func (f *FWindow) OnDestroy(fn func()) *FWindow {
	f.destroy = fn
	return f
}

func (f *FWindow) OnCloseClicked(fn func() bool) *FWindow {
	f.w.ConnectCloseEvent(func(e *gui.QCloseEvent) {
		b := fn()
		if b {
			e.Ignore()
		}
	})
	return f
}
