package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FEventBox struct {
	FBaseView
	v       *gtk.EventBox
	onClick func()
}

func EventBox(i IView) *FEventBox {
	f := &FEventBox{}
	f.v = gtk.NewEventBox()
	
	f.widget = &f.v.Widget

	f.v.Connect("button-press-event", func() {
		if f.onClick != nil {
			f.onClick()
		}
	})
	
	if i != nil {
		f.Add(i)
	}
	return f
}

// ================================================================
func (v *FEventBox) Size(w, h int) *FEventBox {
	v.FBaseView.Size(w, h)
	return v
}
func (f *FEventBox) Assign(v **FEventBox) *FEventBox {
	*v = f
	return f
}
func (vh *ViewHolder) GetEventBoxByItemId(itemId string) *FEventBox {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FEventBox); ok {
			return lv
		}
	}
	return nil
}
func (v *FEventBox) SetItemId(lv *FListView, itemId string) *FEventBox {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}
func GetEventBoxById(id string) *FEventBox {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FEventBox); ok {
			return b
		}
	}
	return nil
}
func (v *FEventBox) baseView() *FBaseView {
	return &v.FBaseView
}
func (v *FEventBox) SetId(id string) *FEventBox {
	idMap[id] = v
	return v
}
func (v *FEventBox) Expand() *FEventBox {
	v.expand = true
	return v
}
func (v *FEventBox) Disable() *FEventBox {
	v.v.SetSensitive(false)
	return v
}
func (v *FEventBox) Enable() *FEventBox {
	v.v.SetSensitive(true)
	return v
}
func (v *FEventBox) Visible() *FEventBox {
	v.v.SetVisible(true)
	return v
}
func (v *FEventBox) Invisible() *FEventBox {
	v.FBaseView.Invisible()
	return v
}

func (v *FEventBox) Tooltips(s string) *FEventBox {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FEventBox) Focus() *FEventBox {
	v.FBaseView.Focus()
	return v
}

func (v *FEventBox) OnDragDrop(f func([]string)) *FEventBox {
	v.FBaseView.OnDragDrop(f)
	return v
}

//====================================================================
func (v *FEventBox) Add(i IView) *FEventBox {
	v.v.Add(i.baseView().widget)
	return v
}
func (v *FEventBox) OnClick(f func()) *FEventBox {
	v.onClick = f
	return v
}
