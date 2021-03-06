package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FAlign struct {
	FBaseView
	v *gtk.Alignment
}

func newAlign(i IView, xa, ya, xs, ys float64) *FAlign {
	f := &FAlign{}
	f.v = gtk.NewAlignment(xa, ya, xs, ys)
	f.widget = &f.v.Widget
	f.v.Add(i.baseView().widget)

	return f
}
func AlignLeft(i IView) *FAlign {
	return newAlign(i, 0, 0, 0, 1)
}
func AlignRight(i IView) *FAlign {
	return newAlign(i, 1, 0, 0, 1)
}
func AlignHCenter(i IView) *FAlign {
	return newAlign(i, 0.5, 0, 0, 1)
}
func AlignTop(i IView) *FAlign {
	return newAlign(i, 0, 0, 1, 0)
}
func AlignBottom(i IView) *FAlign {
	return newAlign(i, 0, 1, 1, 0)
}
func AlignVCenter(i IView) *FAlign {
	return newAlign(i, 0, 0.5, 1, 0)
}
func AlignCenter(i IView) *FAlign {
	return newAlign(i, 0.5, 0.5, 0, 0)
}

// ================================================================
func (v *FAlign) Size(w, h int) *FAlign {
	v.FBaseView.Size(w, h)
	return v
}

func (f *FAlign) Assign(v **FAlign) *FAlign {
	*v = f
	return f
}
func (vh *ViewHolder) GetAlign(itemId string) *FAlign {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FAlign); ok {
			return lv
		}
	}
	return nil
}
func (v *FAlign) SetItemId(lv *FListView, itemId string) *FAlign {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetAlignById(id string) *FAlign {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FAlign); ok {
			return b
		}
	}
	return nil
}
func (v *FAlign) baseView() *FBaseView {
	return &v.FBaseView
}
func (v *FAlign) OnClick(f func()) *FAlign {
	v.v.Connect("clicked", f)
	return v
}
func (v *FAlign) SetId(id string) *FAlign {
	idMap[id] = v
	return v
}
func (v *FAlign) Expand() *FAlign {
	v.expand = true
	return v
}
func (v *FAlign) Disable() *FAlign {
	v.v.SetSensitive(false)
	return v
}
func (v *FAlign) Enable() *FAlign {
	v.v.SetSensitive(true)
	return v
}
func (v *FAlign) Visible() *FAlign {
	v.v.SetVisible(true)
	return v
}
func (v *FAlign) Invisible() *FAlign {
	v.FBaseView.Invisible()
	return v
}

func (v *FAlign) Tooltips(s string) *FAlign {
	v.v.SetTooltipText(s)
	return v
}
func (v *FAlign) Focus() *FAlign {
	v.FBaseView.Focus()
	return v
}

//====================================================================
