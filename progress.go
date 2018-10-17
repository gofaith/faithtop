package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FProgress struct {
	FBaseView
	v *gtk.ProgressBar
}

func Progress() *FProgress {
	v := gtk.NewProgressBar()
	fb := &FProgress{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	setupWidget(fb)
	return fb
}

// ================================================================
func (v *FProgress) Size(w, h int) *FProgress {
	v.FBaseView.Size(w, h)
	return v
}
func (vh *ViewHolder) GetProgressByItemId(itemId string) *FProgress {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FProgress); ok {
			return lv
		}
	}
	return nil
}
func (v *FProgress) SetItemId(lv *FListView, itemId string) *FProgress {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetProgressById(id string) *FProgress {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FProgress); ok {
			return b
		}
	}
	return nil
}
func (v *FProgress) getBaseView() *FBaseView {
	return &v.FBaseView
}
func (v *FProgress) OnClick(f func()) *FProgress {
	v.v.Connect("clicked", f)
	return v
}
func (v *FProgress) SetId(id string) *FProgress {
	idMap[id] = v
	return v
}
func (v *FProgress) Expand() *FProgress {
	v.expand = true
	return v
}
func (v *FProgress) NotFill() *FProgress {
	v.notFill = true
	return v
}
func (v *FProgress) Disable() *FProgress {
	v.v.SetSensitive(false)
	return v
}
func (v *FProgress) Enable() *FProgress {
	v.v.SetSensitive(true)
	return v
}
func (v *FProgress) Visible() *FProgress {
	v.v.SetVisible(true)
	return v
}
func (v *FProgress) Invisible() *FProgress {
	v.v.SetVisible(false)
	return v
}

func (v *FProgress) Tooltips(s string) *FProgress {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FProgress) Focus() *FProgress {
	v.FBaseView.Focus()
	return v
}
func (v *FProgress) Padding(i uint) *FProgress {
	v.padding = i
	return v
}

//====================================================================
func (v *FProgress) Step(f float64) *FProgress {
	v.v.SetFraction(f)
	return v
}
func (v *FProgress) Text(t string) *FProgress {
	v.v.SetText(t)
	return v
}
func (v *FProgress) GetText() string {
	return v.v.GetText()
}
func (v *FProgress) GetStep() float64 {
	return v.v.GetFraction()
}
