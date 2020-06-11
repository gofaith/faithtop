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
		view widgets.QWidget_ITF
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
