//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/quickcontrols2"
	"github.com/therecipe/qt/widgets"
	"os"
)

type AppImpl struct {
	app    *widgets.QApplication
	bridge *QtBridge
}

func init() {
	appImpl = func() IApp {
		v := &AppImpl{
			app: widgets.NewQApplication(len(os.Args), os.Args),
		}
		v.bridge = NewQtBridge(v.app)
		return v
	}
}

func (a *AppImpl) Run() int {
	return a.app.Exec()
}

func (a *AppImpl) Quit() {
	a.app.Quit()
}

func (a *AppImpl) SetClipboardText(s string) IApp {
	a.app.Clipboard().SetText(s, gui.QClipboard__Clipboard)
	return a
}

func (a *AppImpl) GetClipboardText() string {
	return a.app.Clipboard().Text(gui.QClipboard__Clipboard)
}

func (a *AppImpl) OnClipboardTextChanged(fn func(self IApp, s string)) IApp {
	a.app.Clipboard().ConnectChanged(func(mode gui.QClipboard__Mode) {
		if mode == gui.QClipboard__Clipboard {
			fn(a, a.GetClipboardText())
		}
	})
	return a
}

func (a *AppImpl) SetQuickStyle(quickStyle QuickStyle) {
	quickcontrols2.QQuickStyle_SetStyle(string(quickStyle))
}

func (a *AppImpl) SetQuitOnLastWindowClosed(b bool) IApp {
	a.app.SetQuitOnLastWindowClosed(b)
	return a
}

func (a *AppImpl) RunOnUIThread(fn func()) {
	a.bridge.RunOnUIThread(fn)
}

//bridge
type QtBridge struct {
	core.QObject
	_  func() `constructor:"init"`
	_  func() `signal:"runInMainThread,auto"`
	ch chan func()
}

func (b *QtBridge) init() {
	b.ch = make(chan func(), 12)
}

func (b *QtBridge) runInMainThread() {
	select {
	case fn, ok := <-b.ch:
		if !ok {
			return
		}
		fn()
	default:
	}
}

func (b *QtBridge) RunOnUIThread(fn func()) {
	b.ch <- fn
	b.RunInMainThread()
}
