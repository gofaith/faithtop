package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FStatusBar struct {
	FBaseView
	v *gtk.Statusbar
}

func StatusBar() *FStatusBar {
	v := gtk.NewStatusbar()
	fb := &FStatusBar{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	setupWidget(fb)
	return fb
}

// ================================================================

func (v *FStatusBar) OnDragDrop(f func([]string)) *FStatusBar {
	v.FBaseView.OnDragDrop(f)
	return v
}
func (v *FStatusBar) Size(w, h int) *FStatusBar {
	v.FBaseView.Size(w, h)
	return v
}
func (vh *ViewHolder) GetStatusBarByItemId(itemId string) *FStatusBar {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FStatusBar); ok {
			return lv
		}
	}
	return nil
}
func (v *FStatusBar) SetItemId(lv *FListView, itemId string) *FStatusBar {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetStatusBarById(id string) *FStatusBar {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FStatusBar); ok {
			return b
		}
	}
	return nil
}
func (v *FStatusBar) getBaseView() *FBaseView {
	return &v.FBaseView
}
func (v *FStatusBar) OnClick(f func()) *FStatusBar {
	v.v.Connect("clicked", f)
	return v
}
func (v *FStatusBar) SetId(id string) *FStatusBar {
	idMap[id] = v
	return v
}
func (v *FStatusBar) Expand() *FStatusBar {
	v.expand = true
	return v
}
func (v *FStatusBar) NotFill() *FStatusBar {
	v.notFill = true
	return v
}
func (v *FStatusBar) Disable() *FStatusBar {
	v.v.SetSensitive(false)
	return v
}
func (v *FStatusBar) Enable() *FStatusBar {
	v.v.SetSensitive(true)
	return v
}
func (v *FStatusBar) Visible() *FStatusBar {
	v.v.SetVisible(true)
	return v
}
func (v *FStatusBar) Invisible() *FStatusBar {
	v.FBaseView.Invisible()
	return v
}

func (v *FStatusBar) Tooltips(s string) *FStatusBar {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FStatusBar) Focus() *FStatusBar {
	v.FBaseView.Focus()
	return v
}
func (v *FStatusBar) Padding(i uint) *FStatusBar {
	v.padding = i
	return v
}

//====================================================================
func (v *FStatusBar) Text(s string) *FStatusBar {
	context_id := v.v.GetContextId("content_description")
	v.v.Push(context_id, s)
	return v
}
