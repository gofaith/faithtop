package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FLabel struct {
	FBaseView
	v *gtk.Label
}

func Label() *FLabel {
	fb := &FLabel{}
	v := gtk.NewLabel("")
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	setupWidget(fb)
	return fb
}

// ================================================================
func (v *FLabel) Size(w, h int) *FLabel {
	v.FBaseView.Size(w, h)
	return v
}
func (vh *ViewHolder) GetLabelByItemId(itemId string) *FLabel {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FLabel); ok {
			return lv
		}
	}
	return nil
}
func (v *FLabel) SetItemId(lv *FListView, itemId string) *FLabel {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetLabelById(id string) *FLabel {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FLabel); ok {
			return b
		}
	}
	return nil
}
func (v *FLabel) getBaseView() *FBaseView {
	return &v.FBaseView
}
func (v *FLabel) OnClick(f func()) *FLabel {
	v.v.Connect("clicked", f)
	return v
}
func (v *FLabel) SetId(id string) *FLabel {
	idMap[id] = v
	return v
}
func (v *FLabel) Expand() *FLabel {
	v.expand = true
	return v
}
func (v *FLabel) NotFill() *FLabel {
	v.notFill = true
	return v
}
func (v *FLabel) Disable() *FLabel {
	v.v.SetSensitive(false)
	return v
}
func (v *FLabel) Enable() *FLabel {
	v.v.SetSensitive(true)
	return v
}
func (v *FLabel) Visible() *FLabel {
	v.v.SetVisible(true)
	return v
}
func (v *FLabel) Invisible() *FLabel {
	v.v.SetVisible(false)
	return v
}

func (v *FLabel) Tooltips(s string) *FLabel {
	v.v.SetTooltipText(s)
	return v
}
func (v *FLabel) Focus() *FLabel {
	v.FBaseView.Focus()
	return v
}
func (v *FLabel) Padding(i uint) *FLabel {
	v.padding = i
	return v
}

//====================================================================

func (v *FLabel) Text(t string) *FLabel {
	v.v.SetMarkup(t)
	return v
}
func (v *FLabel) GetText() string {
	return v.v.GetText()
}
