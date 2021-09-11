//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type WindowImpl struct {
	window    *widgets.QMainWindow
	deferShow bool
}

func init() {
	newWindowImpl = func() IWindow {
		return &WindowImpl{
			window: widgets.NewQMainWindow(nil, 0),
		}
	}
}

func (w *WindowImpl) Title(title string) IWindow {
	w.window.SetWindowTitle(title)
	return w
}
func (w *WindowImpl) Size(width, height int) IWindow {
	w.window.SetMinimumSize2(width, height)
	return w
}

func (w *WindowImpl) DeferShow() IWindow {
	w.deferShow = true
	return w
}
func (w *WindowImpl) Show() IWindow {
	w.window.Show()
	return w
}

func (w *WindowImpl) CenterWidget(widget IWidget) IWindow {
	w.window.SetCentralWidget(widget.getWidget().(*WidgetImpl).widget)
	if w.deferShow {
		w.Show()
	}
	return w
}

func (w *WindowImpl) Assign(v *IWindow) IWindow {
	*v = w
	return w
}

func (w *WindowImpl) MenuBar(menubar IMenuBar) IWindow {
	w.window.SetMenuBar(menubar.(*MenuBarImpl).menubar)
	return w
}

func (w *WindowImpl) OnClose(fn func() bool) IWindow {
	w.window..ConnectCloseEvent(func(event *gui.QCloseEvent) {
		if fn(){
			event.Ignore()
		}else{
			event.Accept()
		}
	})
	return w
}
