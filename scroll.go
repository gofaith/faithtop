package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FScroll struct {
	FBaseView
	v   *gtk.ScrolledWindow
	box *FBox
}

func newScroll() *FScroll {
	v := gtk.NewScrolledWindow(nil, nil)
	f := &FScroll{}

	f.v = v
	f.widget = &v.Widget
	return f
}

func Scroll(child IView) *FScroll {
	fb := newScroll()
	fb.v.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	fb.v.Add(child.baseView().widget)
	return fb
}
func VScroll() *FScroll {
	fb := newScroll()
	fb.v.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)
	fb.box = VBox()
	fb.v.AddWithViewPort(fb.box.widget)
	return fb
}

func HScroll() *FScroll {
	fb := newScroll()
	fb.v.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_NEVER)
	fb.box = HBox()
	fb.v.AddWithViewPort(fb.box.widget)
	return fb
}

// ================================================================

func (v *FScroll) OnDragDrop(f func([]string)) *FScroll {
	v.FBaseView.OnDragDrop(f)
	return v
}
func (v *FScroll) Size(w, h int) *FScroll {
	v.FBaseView.Size(w, h)
	return v
}
func (f *FScroll) Assign(v **FScroll) *FScroll {
	*v = f
	return f
}
func (vh *ViewHolder) GetScroll(itemId string) *FScroll {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FScroll); ok {
			return lv
		}
	}
	return nil
}
func (v *FScroll) SetItemId(lv *FListView, itemId string) *FScroll {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetScrollById(id string) *FScroll {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FScroll); ok {
			return b
		}
	}
	return nil
}
func (v *FScroll) baseView() *FBaseView {
	return &v.FBaseView
}

func (v *FScroll) OnClick(f func()) *FScroll {
	v.v.Connect("clicked", f)
	return v
}
func (v *FScroll) SetId(id string) *FScroll {
	idMap[id] = v
	return v
}
func (v *FScroll) Expand() *FScroll {
	v.expand = true
	return v
}
func (v *FScroll) Disable() *FScroll {
	v.v.SetSensitive(false)
	return v
}
func (v *FScroll) Enable() *FScroll {
	v.v.SetSensitive(true)
	return v
}
func (v *FScroll) Visible() *FScroll {
	v.v.SetVisible(true)
	return v
}
func (v *FScroll) Invisible() *FScroll {
	v.FBaseView.Invisible()
	return v
}

func (v *FScroll) Tooltips(s string) *FScroll {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FScroll) Focus() *FScroll {
	v.FBaseView.Focus()
	return v
}

//====================================================================
func (v *FScroll) Append(is ...IView) *FScroll {
	if v.box == nil {
		return v
	}
	v.box.Append(is...)
	return v
}
