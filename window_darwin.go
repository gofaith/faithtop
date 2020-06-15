package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type FWindow struct {
	w         *widgets.QMainWindow
	showAfter bool
}

func Win() *FWindow {
	w := widgets.NewQMainWindow(nil, 0)
	f := &FWindow{w: w}
	f.Size(230, 130)
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
	f.w.SetMaximumSize2(w, h)
	return f
}

func (f *FWindow) Show() *FWindow {
	f.w.Show()
	return f
}

func (f *FWindow) Close() {
	f.w.DestroyQMainWindow()
}

func (f *FWindow) Hide() {
	f.w.Hide()
}

func (f *FWindow) IsVisible() bool {
	return f.w.IsVisible()
}

func (f *FWindow) DeferShow() *FWindow {
	f.showAfter = true
	return f
}

func (f *FWindow) Add(i IView) *FWindow {
	f.w.SetCentralWidget(i.baseView().widget)

	if f.showAfter {
		f.Show()
	}
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

func ShowWin(i IView) *FWindow {
	return TopWin().Size(200, 100).Title("").DeferShow().Add(i)
}

func ShowPrompt(w *FWindow, title, label, defaultText string, onSubmit func(string)) {
	dialog := widgets.NewQInputDialog(nil, core.Qt__WindowStaysOnTopHint)
	ok := false
	str := dialog.GetText(w.w, title, label, widgets.QLineEdit__Normal, "text", &ok, core.Qt__WindowStaysOnTopHint, core.Qt__ImhExclusiveInputMask)
	if ok && onSubmit != nil {
		onSubmit(str)
	}
}

func ShowInfo(w *FWindow, title, info string) {
	box := widgets.NewQMessageBox(nil)
	box.Information(nil, title, info, widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func ShowErr(w *FWindow, title, err string) {
	box := widgets.NewQMessageBox(nil)
	box.Critical(nil, title, err, widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
}

func ShowConfirm(w *FWindow, title, text string, onOk, onCancel func()) {
	box := widgets.NewQMessageBox(nil)
	rp := box.Question(nil, title, text, widgets.QMessageBox__Ok|widgets.QMessageBox__No, widgets.QMessageBox__Ok)
	switch rp {
	case widgets.QMessageBox__Ok:
		if onOk != nil {
			onOk()
		}
	case widgets.QMessageBox__No:
		if onCancel != nil {
			onCancel()
		}
	}
}

func (f *FWindow) VBox(is ...IView) *FWindow {
	return f.Add(VBox().Append(is...))
}

func (f *FWindow) HBox(is ...IView) *FWindow {
	return f.Add(HBox().Append(is...))
}
