package faithtop

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

type FRadio struct {
	FBaseView
	v               *gtk.RadioButton
	groupId         string
	onGroupSelected func(string)
}

func Radio(groupId string) *FRadio {
	var group *glib.SList
	if radio, ok := groupIdMap[groupId]; ok {
		group = radio.v.GetGroup()
	}
	v := gtk.NewRadioButton(group)
	fb := &FRadio{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	fb.groupId = groupId
	setupWidget(fb)
	if group == nil {
		groupIdMap[groupId] = fb
	}
	fb.v.Connect("toggled", func() {
		if fb.v.GetActive() && groupIdMap[groupId].onGroupSelected != nil {
			groupIdMap[groupId].onGroupSelected(fb.v.GetLabel())
		}
	})
	return fb
}

// ================================================================
func (v *FRadio) OnEnter(f func()) *FRadio {
	v.FBaseView.OnEnter(f)
	return v
}
func (v *FRadio) Size(w, h int) *FRadio {
	v.FBaseView.Size(w, h)
	return v
}
func (vh *ViewHolder) GetRadioByItemId(itemId string) *FRadio {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FRadio); ok {
			return lv
		}
	}
	return nil
}
func (v *FRadio) SetItemId(lv *FListView, itemId string) *FRadio {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetRadioById(id string) *FRadio {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FRadio); ok {
			return b
		}
	}
	return nil
}
func (v *FRadio) getBaseView() *FBaseView {
	return &v.FBaseView
}
func (v *FRadio) OnClick(f func()) *FRadio {
	v.v.Connect("clicked", f)
	return v
}
func (v *FRadio) SetId(id string) *FRadio {
	idMap[id] = v
	return v
}
func (v *FRadio) Expand() *FRadio {
	v.expand = true
	return v
}
func (v *FRadio) NotFill() *FRadio {
	v.notFill = true
	return v
}
func (v *FRadio) Disable() *FRadio {
	v.v.SetSensitive(false)
	return v
}
func (v *FRadio) Enable() *FRadio {
	v.v.SetSensitive(true)
	return v
}
func (v *FRadio) Visible() *FRadio {
	v.v.SetVisible(true)
	return v
}
func (v *FRadio) Invisible() *FRadio {
	v.FBaseView.Invisible()
	return v
}

func (v *FRadio) Tooltips(s string) *FRadio {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FRadio) Focus() *FRadio {
	v.FBaseView.Focus()
	return v
}
func (v *FRadio) Padding(i uint) *FRadio {
	v.padding = i
	return v
}

//====================================================================

func (v *FRadio) Text(t string) *FRadio {
	v.v.SetLabel(t)
	return v
}
func (v *FRadio) GetText() string {
	return v.v.GetLabel()
}
func (v *FRadio) Selected() *FRadio {
	v.v.SetActive(true)
	return v
}
func (v *FRadio) OnGroupSelected(f func(string)) *FRadio {
	groupIdMap[v.groupId].onGroupSelected = f
	return v
}
