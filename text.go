package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FText struct {
	FBaseView
	v *gtk.Label
}

func Text(t string) *FText {
	fb := &FText{}
	v := gtk.NewLabel("")
	v.SetMarkup(t)
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	setupWidget(fb)
	return fb
}

// ================================================================
func (v *FText) Size(w, h int) *FText {
	v.FBaseView.Size(w, h)
	return v
}
func (vh *ViewHolder) GetTextByItemId(itemId string) *FText {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FText); ok {
			return lv
		}
	}
	return nil
}
func (v *FText) SetItemId(lv *FListView, itemId string) *FText {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetTextById(id string) *FText {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FText); ok {
			return b
		}
	}
	return nil
}
func (v *FText) getBaseView() *FBaseView {
	return &v.FBaseView
}
func (v *FText) SetId(id string) *FText {
	idMap[id] = v
	return v
}
func (v *FText) Expand() *FText {
	v.expand = true
	return v
}
func (v *FText) NotFill() *FText {
	v.notFill = true
	return v
}
func (v *FText) Disable() *FText {
	v.v.SetSensitive(false)
	return v
}
func (v *FText) Enable() *FText {
	v.v.SetSensitive(true)
	return v
}
func (v *FText) Visible() *FText {
	v.v.SetVisible(true)
	return v
}
func (v *FText) Invisible() *FText {
	v.FBaseView.Invisible()
	return v
}

func (v *FText) Tooltips(s string) *FText {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FText) Focus() *FText {
	v.FBaseView.Focus()
	return v
}
func (v *FText) Padding(i uint) *FText {
	v.padding = i
	return v
}

//====================================================================

func (v *FText) Text(t string) *FText {
	v.v.SetMarkup(t)
	return v
}
func (v *FText) PlainText(t string) *FText {
	v.v.SetText(t)
	return v
}
func (v *FText) GetText() string {
	return v.v.GetText()
}
func (v *FText) Selectable() *FText {
	v.v.SetSelectable(true)
	return v
}
func (v *FText) Unselectable() *FText {
	v.v.SetSelectable(false)
	return v
}
func (v *FText) LeftJustify() *FText {
	v.v.SetJustify(gtk.JUSTIFY_LEFT)
	return v
}
func (v *FText) RightJustify() *FText {
	v.v.SetJustify(gtk.JUSTIFY_RIGHT)
	return v
}
func (v *FText) CenterJustify() *FText {
	v.v.SetJustify(gtk.JUSTIFY_CENTER)
	return v
}
func (v *FText) FillJustify() *FText {
	v.v.SetJustify(gtk.JUSTIFY_FILL)
	return v
}
