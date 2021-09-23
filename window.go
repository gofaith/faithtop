package faithtop

type IWindow interface {
	DeferShow() IWindow
	Assign(v *IWindow) IWindow
	Title(title string) IWindow
	Size(with, height int) IWindow
	CenterWidget(widget IWidget) IWindow
	Show() IWindow
	MenuBar(menubar IMenuBar) IWindow
	OnClose(fn func() bool) IWindow
	IsActiveWindow() bool
	OnChanged(fn func()) IWindow
	Close() 
	RunOnUIThread(fn func()) IWindow
}

var newWindowImpl func() IWindow

func Window() IWindow {
	return newWindowImpl()
}
