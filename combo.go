package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FCombo struct {
	FBaseView
	v    *gtk.ComboBox
	strs []string
}

func Combo() *FCombo {
	v := gtk.NewComboBoxNewText()
	fb := &FCombo{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	setupWidget(fb)
	return fb
}

// ================================================================

func (v *FCombo) Size(w, h int) *FCombo {
	v.FBaseView.Size(w, h)
	return v
}
func (vh *ViewHolder) GetComboByItemId(itemId string) *FCombo {
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
func (v *FCombo) getBaseView() *FBaseView {
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
func (v *FCombo) NotFill() *FCombo {
	v.notFill = true
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
func (v *FCombo) Padding(i uint) *FCombo {
	v.padding = i
	return v
}

func (v *FCombo) OnDragDrop(f func([]string)) *FCombo {
	v.FBaseView.OnDragDrop(f)
	return v
}

//====================================================================
func (v *FCombo) GetActiveText() string {
	return v.v.GetActiveText()
}
func (v *FCombo) GetAllText() []string {
	return v.strs
}
func (v *FCombo) AppendText(ts ...string) *FCombo {
	for i := 0; i < len(v.strs); i++ {
		v.v.RemoveText(0)
	}
	for _, t := range ts {
		v.v.InsertText(t, 0)
	}
	v.strs = ts
	if len(ts) > 0 {
		v.v.SetActive(0)
	}
	return v
}
func (f *FCombo) ActiveText(index int) *FCombo {
	f.v.SetActive(index)
	return f
}
func (v *FCombo) OnChange(f func(string)) *FCombo {
	v.v.Connect("changed", func() {
		f(v.GetActiveText())
	})
	return v
}
