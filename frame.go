package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FFrame struct {
	FBaseView
	v *gtk.Frame
}

func Frame(title string, child IView) *FFrame {
	v := gtk.NewFrame(title)
	fb := &FFrame{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	setupWidget(fb)
	if child != nil {
		fb.Add(child)
		fb.afterShownFn = child.getBaseView().afterShownFn
	}
	return fb
}

// ================================================================
func (v *FFrame) Size(w, h int) *FFrame {
	v.FBaseView.Size(w, h)
	return v
}
func (vh *ViewHolder) GetFrameByItemId(itemId string) *FFrame {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FFrame); ok {
			return lv
		}
	}
	return nil
}
func (v *FFrame) SetItemId(lv *FListView, itemId string) *FFrame {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetFrameById(id string) *FFrame {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FFrame); ok {
			return b
		}
	}
	return nil
}
func (v *FFrame) getBaseView() *FBaseView {
	return &v.FBaseView
}
func (v *FFrame) OnClick(f func()) *FFrame {
	v.v.Connect("clicked", f)
	return v
}
func (v *FFrame) SetId(id string) *FFrame {
	idMap[id] = v
	return v
}
func (v *FFrame) Expand() *FFrame {
	v.expand = true
	return v
}
func (v *FFrame) NotFill() *FFrame {
	v.notFill = true
	return v
}
func (v *FFrame) Disable() *FFrame {
	v.v.SetSensitive(false)
	return v
}
func (v *FFrame) Enable() *FFrame {
	v.v.SetSensitive(true)
	return v
}
func (v *FFrame) Visible() *FFrame {
	v.v.SetVisible(true)
	return v
}
func (v *FFrame) Invisible() *FFrame {
	v.v.SetVisible(false)
	return v
}

func (v *FFrame) Tooltips(s string) *FFrame {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FFrame) Focus() *FFrame {
	v.FBaseView.Focus()
	return v
}
func (v *FFrame) Padding(i uint) *FFrame {
	v.padding = i
	return v
}

//====================================================================

func (v *FFrame) Text(t string) *FFrame {
	v.v.SetLabel(t)
	return v
}
func (v *FFrame) GetText() string {
	return v.v.GetLabel()
}
func (v *FFrame) Add(i IView) *FFrame {
	v.v.Add(i.getBaseView().view)
	return v
}
func (v *FFrame) Vbox(is ...IView) *FFrame {
	return v.Add(VBox().Append(is...))
}
func (v *FFrame) HBox(is ...IView) *FFrame {
	return v.Add(HBox().Append(is...))
}
