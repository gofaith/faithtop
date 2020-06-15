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
		expand bool
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

func (f *FBaseView) baseView() *FBaseView {
	return f
}

// get
func (f *FBaseView) Size(w, h int) {
	f.widget.QWidget_PTR().SetMinimumSize2(w, h)
}
func (f *FBaseView) IsEnabled() bool {
	return f.widget.QWidget_PTR().IsEnabled()
}
func (f *FBaseView) IsVisible() bool {
	return f.widget.QWidget_PTR().IsVisible()
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

func (f *FBaseView) Expand() {
	f.expand = true
}

func (f *FBaseView) Invisible() {
	f.widget.QWidget_PTR().SetVisible(false)
}

func (f *FBaseView) Visible() {
	f.widget.QWidget_PTR().SetVisible(true)
}

// func (f *FBaseView) OnDragDrop(fn func([]string)) *FBaseView {
// 	f.widget.QWidget_PTR().SetAcceptDrops(true)
// 	f.widget.QWidget_PTR().SetAutoFillBackground(true)
// 	f.widget.QWidget_PTR().ConnectDragEnterEvent(func(e *gui.QDragEnterEvent) {
// 		e.AcceptProposedAction()
// 		// pp.Println(e.MimeData())
// 	})
// 	f.widget.QWidget_PTR().ConnectDragMoveEvent(func(e *gui.QDragMoveEvent) {
// 		e.AcceptProposedAction()
// 	})
// 	f.widget.QWidget_PTR().ConnectDropEvent(func(e *gui.QDropEvent) {
// 		var mimeData = e.MimeData()
// 		switch {
// 		// case mimeData.HasImage():
// 		// 		dropArea.SetPixmap(gui.QPixmap_FromImage(gui.NewQImageFromPointer(mimeData.ImageData().ToImage()), 0))
// 		// case mimeData.HasHtml():
// 		// 		dropArea.SetText(mimeData.Html())
// 		// 		dropArea.SetTextFormat(core.Qt__RichText)
// 		case mimeData.HasUrls():
// 			urlList := mimeData.Urls()
// 			out := []string{}
// 			for i := 0; i < len(urlList) && i < 32; i++ {
// 				out = append(out, urlList[i].Path(0))
// 			}
// 			fn(out)
// 			// case mimeData.HasText():
// 			// 	{
// 			// 		dropArea.SetText(mimeData.Text())
// 			// 		dropArea.SetTextFormat(core.Qt__PlainText)
// 			// 	}

// 			// default:
// 			// 	{
// 			// 		dropArea.SetText("can't display data")
// 			// 	}
// 		}

// 		e.AcceptProposedAction()
// 	})
// 	return f
// }
