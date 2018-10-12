package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FButton struct {
	FBaseView
	v *gtk.Button
}

func Button() *FButton {
	v := gtk.NewButton()
	fb := &FButton{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	setupWidget(fb)
	return fb
}

func GetButtonById(id string) *FButton {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FButton); ok {
			return b
		}
	}
	return nil
}

func (vh *ViewHolder) GetButtonByItemId(id string) *FButton {
	if v, ok := vh.vlist[id]; ok {
		if bt, ok := v.(*FButton); ok {
			return bt
		}
	}
	return nil
}

// ----------------------------------------------------------
func (v *FButton) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FButton) SetId(id string) *FButton {
	idMap[id] = v
	return v
}
