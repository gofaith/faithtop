package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FMenuItem struct {
	FBaseView
	v       *gtk.MenuItem
	subMenu *gtk.Menu
}

func MenuItem(s string) *FMenuItem {
	fb := &FMenuItem{}
	fb.v = gtk.NewMenuItemWithMnemonic(s)
	fb.view = fb.v
	fb.widget = &fb.v.Widget
	return fb
}
func (v *FMenuItem) OnClick(f func()) *FMenuItem {
	v.v.Connect("activate", f)
	return v
}
func (v *FMenuItem) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FMenuItem) SubMenu(ms ...*FMenuItem) *FMenuItem {
	if v.subMenu == nil {
		v.subMenu = gtk.NewMenu()
	}
	for _, m := range ms {
		v.subMenu.Append(m.v)
	}
	v.v.SetSubmenu(v.subMenu)
	return v
}
