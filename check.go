package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FCheck struct {
	FBaseView
	v *gtk.CheckButton
}

func Check() *FCheck {
	v := gtk.NewCheckButton()
	fb := &FCheck{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	setupWidget(fb)
	return fb
}

// ================================================================
func (v *FCheck) Size(w, h int) *FCheck {
	v.FBaseView.Size(w, h)
	return v
}
func (vh *ViewHolder) GetCheckByItemId(itemId string) *FCheck {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FCheck); ok {
			return lv
		}
	}
	return nil
}
func (v *FCheck) SetItemId(lv *FListView, itemId string) *FCheck {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetCheckById(id string) *FCheck {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FCheck); ok {
			return b
		}
	}
	return nil
}
func (v *FCheck) getBaseView() *FBaseView {
	return &v.FBaseView
}
func (v *FCheck) OnClick(f func()) *FCheck {
	v.v.Connect("clicked", f)
	return v
}
func (v *FCheck) SetId(id string) *FCheck {
	idMap[id] = v
	return v
}
func (v *FCheck) Expand() *FCheck {
	v.expand = true
	return v
}
func (v *FCheck) NotFill() *FCheck {
	v.notFill = true
	return v
}
func (v *FCheck) Disable() *FCheck {
	v.v.SetSensitive(false)
	return v
}
func (v *FCheck) Enable() *FCheck {
	v.v.SetSensitive(true)
	return v
}
func (v *FCheck) Visible() *FCheck {
	v.v.SetVisible(true)
	return v
}
func (v *FCheck) Invisible() *FCheck {
	v.v.SetVisible(false)
	return v
}

func (v *FCheck) Tooltips(s string) *FCheck {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FCheck) Focus() *FCheck {
	v.FBaseView.Focus()
	return v
}
func (v *FCheck) Padding(i uint) *FCheck {
	v.padding = i
	return v
}

//====================================================================
func (v *FCheck) OnChange(f func(bool)) *FCheck {
	v.v.Connect("toggled", func() {
		if v.v.GetActive() {
			f(true)
		} else {
			f(false)
		}
	})
	return v
}
func (v *FCheck) Checked() *FCheck {
	v.v.SetActive(true)
	return v
}
func (v *FCheck) Unchecked() *FCheck {
	v.v.SetActive(false)
	return v
}
func (v *FCheck) GetChecked() bool {
	return v.v.GetActive()
}
