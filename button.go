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
func (v *FButton) Size(w, h int) *FButton {
	v.FBaseView.Size(w, h)
	return v
}
func (v *FButton) getBaseView() *FBaseView {
	return &v.FBaseView
}
func GetButtonById(id string) *FButton {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FButton); ok {
			return b
		}
	}
	return nil
}

// ================================================================
func (v *FButton) OnClick(f func()) *FButton {
	v.v.Connect("clicked", f)
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
	v.v.SetVisible(false)
	return v
}

func (v *FButton) Tooltips(s string) *FButton {
	v.v.SetTooltipText(s)
	return v
}
func (v *FButton) Focus() *FButton {
	currentFocus = v.widget
	if currentWin != nil {
		currentWin.SetFocusChild(v.widget)
	}
	return v
}
func (v *FButton) Padding(i uint) *FButton {
	v.padding = i
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
