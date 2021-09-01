package faithtop

type IWindow interface {
	Title(title string) IWindow
	Size(with, height int) IWindow
	CenterWidget(widget IWidget) IWindow
	Show() IWindow
}

var newWindowImpl func() IWindow

func Window() IWindow {
	return newWindowImpl()
}
