package faithtop

type IWindow interface {
	Title(title string) IWindow
	Size(with, height int) IWindow
	Show() IWindow
}

var windowImpl func() IWindow

func Window() IWindow {
	return windowImpl()
}
