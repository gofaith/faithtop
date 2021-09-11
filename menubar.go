package faithtop

type (
	IMenuBar interface {
		Menus(menus ...IMenu) IMenuBar
	}
	IMenu interface {
		Actions(actions ...IAction) IMenu
	}
	IAction interface {
		OnClick(fn func()) IAction
	}
)

var (
	newMenuBarImpl func() IMenuBar
	newMenuImpl    func(text string) IMenu
	newActionImpl  func(text string) IAction
)

func MenuBar(menus ...IMenu) IMenuBar {
	return newMenuBarImpl().Menus(menus...)
}

func Menu(text string) IMenu {
	return newMenuImpl(text)
}

func Action(text string) IAction {
	return newActionImpl(text)
}
