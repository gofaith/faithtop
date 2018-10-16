package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FMenuBar struct {
	FBaseView
	v *gtk.MenuBar
}

func MenuBar() *FMenuBar {
	v := gtk.NewMenuBar()
	fb := &FMenuBar{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	setupWidget(fb)
	return fb
}

// ================================================================
func (v *FMenuBar) Size(w, h int) *FMenuBar {
	v.FBaseView.Size(w, h)
	return v
}
func (vh *ViewHolder) GetMenuBarByItemId(itemId string) *FMenuBar {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FMenuBar); ok {
			return lv
		}
	}
	return nil
}
func (v *FMenuBar) SetItemId(lv *FListView, itemId string) *FMenuBar {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetMenuBarById(id string) *FMenuBar {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FMenuBar); ok {
			return b
		}
	}
	return nil
}
func (v *FMenuBar) getBaseView() *FBaseView {
	return &v.FBaseView
}
func (v *FMenuBar) OnClick(f func()) *FMenuBar {
	v.v.Connect("clicked", f)
	return v
}
func (v *FMenuBar) SetId(id string) *FMenuBar {
	idMap[id] = v
	return v
}
func (v *FMenuBar) Expand() *FMenuBar {
	v.expand = true
	return v
}
func (v *FMenuBar) NotFill() *FMenuBar {
	v.notFill = true
	return v
}
func (v *FMenuBar) Disable() *FMenuBar {
	v.v.SetSensitive(false)
	return v
}
func (v *FMenuBar) Enable() *FMenuBar {
	v.v.SetSensitive(true)
	return v
}
func (v *FMenuBar) Visible() *FMenuBar {
	v.v.SetVisible(true)
	return v
}
func (v *FMenuBar) Invisible() *FMenuBar {
	v.v.SetVisible(false)
	return v
}

func (v *FMenuBar) Tooltips(s string) *FMenuBar {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FMenuBar) Focus() *FMenuBar {
	v.FBaseView.Focus()
	return v
}
func (v *FMenuBar) Padding(i uint) *FMenuBar {
	v.padding = i
	return v
}

//====================================================================
func (v *FMenuBar) Menus(is ...IMenuItem) *FMenuBar {
	for _, i := range is {
		v.v.Append(i.getMenuItem())
	}
	return v
}
