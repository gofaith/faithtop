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
	newMenuImpl = func() IMenu {
		return &MenuImpl{
			menu: widgets.NewQMenu(nil),
		}
	}
	newActionImpl = func(text string) IAction {
		return &ActionImpl{
			action: widgets.NewQAction2(text, nil),
		}
	}
	showPopupMenuImpl = func(win IWindow, anchor IWidget, menu IMenu) {
		target := anchor.getWidget().(*WidgetImpl).widget
		pos := target.MapToGlobal(win.(*WindowImpl).window.CentralWidget().Pos())
		pos.SetY(pos.Y() + target.Height())
		menu.(*MenuImpl).menu.Exec2(pos, nil)
	}
}

func (m *MenuBarImpl) Menus(menus ...IMenu) IMenuBar {
	for _, menu := range menus {
		m.menubar.AddMenu(menu.(*MenuImpl).menu)
	}
	return m
}

func (m *MenuImpl) Title(title string) IMenu {
	m.menu.SetTitle(title)
	return m
}

func (m *MenuImpl) Actions(actions ...IAction) IMenu {
	vs := []*widgets.QAction{}
	for _, action := range actions {
		vs = append(vs, action.(*ActionImpl).action)
	}
	m.menu.Clear()
	m.menu.AddActions(vs)
	return m
}

func (a *ActionImpl) OnClick(fn func()) IAction {
	a.action.ConnectTriggered(func(bool) {
		fn()
	})
	return a
}

func (a *MenuImpl) Assign(v *IMenu) IMenu {
	*v = a
	return a
}

func (a *ActionImpl) Assign(v *IAction) IAction {
	*v = a
	return a
}

func (a *ActionImpl) Text(text string) IAction {
	a.action.SetText(text)
	return a
}

func (a *ActionImpl) GetText() string {
	return a.action.Text()
}

func (m *MenuImpl) AddAction(i IAction) IMenu {
	m.menu.AddActions([]*widgets.QAction{i.(*ActionImpl).action})
	return m
}

func (m *MenuImpl) Clear() IMenu {
	m.menu.Clear()
	return m
}

func (m *MenuImpl) RemoveAction(action IAction) IMenu {
	m.menu.RemoveAction(action.(*ActionImpl).action)
	return m
}
