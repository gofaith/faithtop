package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type IMenuItem interface {
	getMenuItem() *gtk.MenuItem
}
type FMenuItem struct {
	FBaseView
	v       *gtk.MenuItem
	subMenu *gtk.Menu
}

func (v *FMenuItem) getMenuItem() *gtk.MenuItem {
	return v.v
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

func (v *FMenuItem) SubMenu(ms ...IMenuItem) *FMenuItem {
	if v.subMenu == nil {
		v.subMenu = gtk.NewMenu()
	}
	for _, m := range ms {
		v.subMenu.Append(m.getMenuItem())
	}
	v.v.SetSubmenu(v.subMenu)
	return v
}

func (f *FMenuItem) Assign(v **FMenuItem) *FMenuItem {
	*v = f
	return f
}
func (v *FMenuItem) SetId(s string) *FMenuItem {
	idMap[s] = v
	return v
}
func GetMenuItemById(id string) *FMenuItem {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FMenuItem); ok {
			return b
		}
	}
	return nil
}

//======================================================================

type FCheckMenuItem struct {
	FBaseView
	v       *gtk.CheckMenuItem
	subMenu *gtk.Menu
}

func (f *FCheckMenuItem) Assign(v **FCheckMenuItem) *FCheckMenuItem {
	*v = f
	return f
}

func (v *FCheckMenuItem) getMenuItem() *gtk.MenuItem {
	return &v.v.MenuItem
}
func CheckMenuItem(s string) *FCheckMenuItem {
	fb := &FCheckMenuItem{}
	fb.v = gtk.NewCheckMenuItemWithMnemonic(s)
	fb.view = fb.v
	fb.widget = &fb.v.Widget
	return fb
}

func (v *FCheckMenuItem) OnChange(f func(bool)) *FCheckMenuItem {
	v.v.Connect("activate", func() {
		f(v.v.GetActive())
	})
	return v
}
func (v *FCheckMenuItem) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FCheckMenuItem) SubMenu(ms ...IMenuItem) *FCheckMenuItem {
	if v.subMenu == nil {
		v.subMenu = gtk.NewMenu()
	}
	for _, m := range ms {
		v.subMenu.Append(m.getMenuItem())
	}
	v.v.SetSubmenu(v.subMenu)
	return v
}
func (v *FCheckMenuItem) Checked() *FCheckMenuItem {
	v.v.SetActive(true)
	return v
}

func (v *FCheckMenuItem) Unchecked() *FCheckMenuItem {
	v.v.SetActive(false)
	return v
}

func (v *FCheckMenuItem) SetId(s string) *FCheckMenuItem {
	idMap[s] = v
	return v
}
func GetCheckMenuItemById(id string) *FCheckMenuItem {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FCheckMenuItem); ok {
			return b
		}
	}
	return nil
}
