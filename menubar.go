package faithtop

type (
	IMenuBar interface {
		Menus(menus ...IMenu) IMenuBar
	}
	IMenu interface {
		Title(title string) IMenu
		Actions(actions ...IAction) IMenu
	}
	IAction interface {
		OnClick(fn func()) IAction
	}
)

var (
	newMenuBarImpl    func() IMenuBar
	newMenuImpl       func() IMenu
	newActionImpl     func(text string) IAction
	showPopupMenuImpl func(win IWindow, anchor IWidget, menu IMenu)
)

func MenuBar(menus ...IMenu) IMenuBar {
	return newMenuBarImpl().Menus(menus...)
}

func Menu() IMenu {
	return newMenuImpl()
}

func Menu2(text string) IMenu {
	return newMenuImpl().Title(text)
}

func Action(text string) IAction {
	return newActionImpl(text)
}

func ShowPopupMenu(win IWindow, anchor IWidget, menu IMenu) {
	showPopupMenuImpl(win, anchor, menu)
}
