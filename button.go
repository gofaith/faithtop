package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FButton struct {
	FBaseView
	v *gtk.Button
	onClick func()
}

func Button() *FButton {
	v := gtk.NewButton()
	fb := &FButton{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	fb.v.Connect("clicked", func(){
		if fb.onClick!=nil{
			fb.onClick()
		}
	})
	setupWidget(fb)
	return fb
}

// ================================================================
func (v *FButton) Size(w, h int) *FButton {
	v.FBaseView.Size(w, h)
	return v
}
func (vh *ViewHolder) GetButtonByItemId(itemId string) *FButton {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FButton); ok {
			return lv
		}
	}
	return nil
}
func (v *FButton) SetItemId(lv *FListView, itemId string) *FButton {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetButtonById(id string) *FButton {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FButton); ok {
			return b
		}
	}
	return nil
}
func (v *FButton) getBaseView() *FBaseView {
	return &v.FBaseView
}
func (v *FButton) OnClick(f func()) *FButton {
	v.onClick=f
	return v
}
func (v *FButton) SetId(id string) *FButton {
	idMap[id] = v
	return v
}
func (v *FButton) Expand() *FButton {
	v.expand = true
	return v
}
func (v *FButton) NotFill() *FButton {
	v.notFill = true
	return v
}
func (v *FButton) Disable() *FButton {
	v.v.SetSensitive(false)
	return v
}
func (v *FButton) Enable() *FButton {
	v.v.SetSensitive(true)
	return v
}
func (v *FButton) Visible() *FButton {
	v.v.SetVisible(true)
	return v
}
func (v *FButton) Invisible() *FButton {
	v.FBaseView.Invisible()
	return v
}

func (v *FButton) Tooltips(s string) *FButton {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FButton) Focus() *FButton {
	v.FBaseView.Focus()
	return v
}
func (v *FButton) Padding(i uint) *FButton {
	v.padding = i
	return v
}

func (v *FButton) OnDragDrop(f func([]string)) *FButton {
	v.FBaseView.OnDragDrop(f)
	return v
}

//====================================================================

func (v *FButton) Text(t string) *FButton {
	v.v.SetLabel(t)
	return v
}
func (v *FButton) GetText() string {
	return v.v.GetLabel()
}
func (v *FButton) Image(img *FImage) *FButton {
	v.v.SetImage(img.v)
	v.afterShownFn = img.afterShownFn
	return v
}
