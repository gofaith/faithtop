package faithtop

import (
	"github.com/StevenZack/livedata"
	"github.com/mattn/go-gtk/gtk"
)

type FCombo struct {
	FBaseView
	v                 *gtk.ComboBox
	strs              []string
	isActivatedByCode bool
	currentText       string
}

func Combo() *FCombo {
	v := gtk.NewComboBoxNewText()
	fb := &FCombo{}
	fb.v = v

	fb.widget = &v.Widget

	return fb
}

// ================================================================

func (v *FCombo) Size(w, h int) *FCombo {
	v.FBaseView.Size(w, h)
	return v
}

func (f *FCombo) Assign(v **FCombo) *FCombo {
	*v = f
	return f
}
func (vh *ViewHolder) GetCombo(itemId string) *FCombo {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FCombo); ok {
			return lv
		}
	}
	return nil
}
func (v *FCombo) SetItemId(lv *FListView, itemId string) *FCombo {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetComboById(id string) *FCombo {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FCombo); ok {
			return b
		}
	}
	return nil
}
func (v *FCombo) baseView() *FBaseView {
	return &v.FBaseView
}
func (v *FCombo) OnClick(f func()) *FCombo {
	v.v.Connect("clicked", f)
	return v
}
func (v *FCombo) SetId(id string) *FCombo {
	idMap[id] = v
	return v
}
func (v *FCombo) Expand() *FCombo {
	v.expand = true
	return v
}
func (v *FCombo) Disable() *FCombo {
	v.v.SetSensitive(false)
	return v
}
func (v *FCombo) Enable() *FCombo {
	v.v.SetSensitive(true)
	return v
}
func (v *FCombo) Visible() *FCombo {
	v.v.SetVisible(true)
	return v
}
func (v *FCombo) Invisible() *FCombo {
	v.FBaseView.Invisible()
	return v
}

func (v *FCombo) Tooltips(s string) *FCombo {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FCombo) Focus() *FCombo {
	v.FBaseView.Focus()
	return v
}

func (v *FCombo) OnDragDrop(f func([]string)) *FCombo {
	v.FBaseView.OnDragDrop(f)
	return v
}

//====================================================================
func (v *FCombo) GetSelectedText() string {
	return v.v.GetActiveText()
}
func (v *FCombo) SetData(ts []string) *FCombo {
	for i := 0; i < len(v.strs); i++ {
		v.v.RemoveText(0)
	}
	for _, t := range ts {
		v.v.AppendText(t)
	}
	v.strs = ts
	return v
}
func (f *FCombo) GetSelection() int {
	return f.v.GetActive()
}
func (f *FCombo) Select(index int) *FCombo {
	f.isActivatedByCode = true
	f.v.SetActive(index)
	return f
}
func (v *FCombo) OnChange(f func(str string)) *FCombo {
	v.v.Connect("changed", func() {
		f(v.GetSelectedText())
	})
	return v
}

func (f *FCombo) BindData(signal *livedata.Bool, fn func(*FCombo)) *FCombo {
	signal.ObserveForever(func(b bool) {
		if b {
			RunOnUIThread(func() {
				fn(f)
			})
		}
	})
	return f
}

func (f *FCombo) BindSelection(i *livedata.Int) *FCombo {
	i.ObserveForever(func(idx int) {
		if idx == f.GetSelection() || idx < 0 {
			return
		}
		RunOnUIThread(func() {
			f.Select(idx)
		})
	})
	f.v.Connect("changed", func() {
		sel := f.GetSelection()
		if sel != i.Get() {
			i.Post(sel)
		}
	})
	return f
}
