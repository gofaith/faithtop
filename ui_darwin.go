package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type (
	IView interface {
		baseView() *FBaseView
	}
	FBaseView struct {
		widget widgets.QWidget_ITF
		layout widgets.QLayout_ITF
		expand bool
		align  core.Qt__AlignmentFlag
	}
	qmlBridge struct {
		core.QObject
		_ func(fn func()) `signal:sendToQml`
	}
)

var qml = NewQmlBridge(nil)

func init() {
	qml.ConnectSendToQml(func(f func()) {
		if f != nil {
			f()
		}
	})
}

func RunOnUIThread(f func()) {
	qml.SendToQml(f)
}

// get
func (f *FBaseView) isLayout() bool {
	return f.layout != nil
}

func (f *FBaseView) Size(w, h int) {
	f.widget.QWidget_PTR().SetMinimumSize2(w, h)
}
func (f *FBaseView) IsEnabled() bool {
	return f.widget.QWidget_PTR().IsEnabled()
}
func (f *FBaseView) IsVisible() bool {
	return f.widget.QWidget_PTR().IsVisible()
}
func (f *FBaseView) GetWidth() int {
	return f.widget.QWidget_PTR().Width()
}
func (f *FBaseView) GetHeight() int {
	return f.widget.QWidget_PTR().Height()
}
func (f *FBaseView) GetX() int {
	return f.widget.QWidget_PTR().X()
}
func (f *FBaseView) GetY() int {
	return f.widget.QWidget_PTR().Y()
}
func (f *FBaseView) Focus() {
	f.widget.QWidget_PTR().SetFocus2()
}

// base set

func (f *FBaseView) Expand(b bool) {
	f.expand = b
}

func (f *FBaseView) AlignStart() {
	f.align = core.Qt__AlignLeading
}

func (f *FBaseView) AlignEnd() {
	f.align = core.Qt__AlignTrailing
}

func (f *FBaseView) AlignCenter() {
	f.align = core.Qt__AlignCenter
}

func (f *FBaseView) AlignStretch() {
	f.align = 0
}
