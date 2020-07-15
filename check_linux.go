package faithtop

import (
	"github.com/StevenZack/livedata"
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
	fb.widget = &v.Widget

	return fb
}

// ================================================================
func (v *FCheck) Size(w, h int) *FCheck {
	v.FBaseView.Size(w, h)
	return v
}
func (f *FCheck) Assign(v **FCheck) *FCheck {
	*v = f
	return f
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
func (v *FCheck) baseView() *FBaseView {
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
	v.FBaseView.Invisible()
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

func (v *FCheck) OnDragDrop(f func([]string)) *FCheck {
	v.FBaseView.OnDragDrop(f)
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
func (v *FCheck) Checked(b bool) *FCheck {
	v.v.SetActive(b)
	return v
}
func (v *FCheck) GetChecked() bool {
	return v.v.GetActive()
}
func (v *FCheck) Text(t string) *FCheck {
	v.v.SetLabel(t)
	return v
}

func (f *FCheck) BindCheck(l *livedata.Bool) *FCheck {
	l.ObserveForever(func(b bool) {
		RunOnUIThread(func() {
			f.Checked(b)
		})
	})
	f.v.Connect("toggled", func() {
		l.Post(f.GetChecked())
	})
	return f
}
