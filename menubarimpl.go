//go:build impl

package faithtop

import "github.com/therecipe/qt/widgets"

type (
	MenuBarImpl struct {
		menubar *widgets.QMenuBar
	}
	MenuImpl struct {
		menu *widgets.QMenu
	}
	ActionImpl struct {
		action *widgets.QAction
	}
)

func init() {
	newMenuBarImpl = func() IMenuBar {
		return &MenuBarImpl{
			menubar: widgets.NewQMenuBar(nil),
		}
	}
	newMenuImpl = func(text string) IMenu {
		return &MenuImpl{
			menu: widgets.NewQMenu2(text, nil),
		}
	}
	newActionImpl = func(text string) IAction {
		return &ActionImpl{
			action: widgets.NewQAction2(text, nil),
		}
	}
}

func (m *MenuBarImpl) Menus(menus ...IMenu) IMenuBar {
	for _, menu := range menus {
		m.menubar.AddMenu(menu.(*MenuImpl).menu)
	}
	return m
}

func (m *MenuImpl) Actions(actions ...IAction) IMenu {
	vs := []*widgets.QAction{}
	for _, action := range actions {
		vs = append(vs, action.(*ActionImpl).action)
	}
	m.menu.AddActions(vs)
	return m
}

func (a *ActionImpl) OnClick(fn func()) IAction {
	a.action.ConnectTriggered(func(bool) {
		fn()
	})
	return a
}
