package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FBox struct {
	FBaseView
	v            *gtk.Box
	isHorizontal bool //false :vertical , true:horizontal
}

func VBox() *FBox {
	f := &FBox{}
	f.v = &gtk.NewVBox(false, 0).Box
	f.view = f.v
	f.widget = &f.v.Widget
	f.isHorizontal = false
	setupWidget(f)
	return f
}

func HBox() *FBox {
	f := &FBox{}
	f.v = &gtk.NewHBox(false, 0).Box
	f.view = f.v
	f.widget = &f.v.Widget
	f.isHorizontal = true
	setupWidget(f)
	return f
}

// ================================================================
func (v *FBox) Size(w, h int) *FBox {
	v.FBaseView.Size(w, h)
	return v
}

func (f *FBox) Assign(v **FBox) *FBox {
	*v = f
	return f
}
func (vh *ViewHolder) GetBoxByItemId(itemId string) *FBox {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FBox); ok {
			return lv
		}
	}
	return nil
}
func (v *FBox) SetItemId(lv *FListView, itemId string) *FBox {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}
func GetBoxById(id string) *FBox {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FBox); ok {
			return b
		}
	}
	return nil
}
func (v *FBox) getBaseView() *FBaseView {
	return &v.FBaseView
}
func (v *FBox) OnClick(f func()) *FBox {
	v.v.Connect("clicked", f)
	return v
}
func (v *FBox) SetId(id string) *FBox {
	idMap[id] = v
	return v
}
func (v *FBox) Expand() *FBox {
	v.expand = true
	return v
}
func (v *FBox) NotFill() *FBox {
	v.notFill = true
	return v
}
func (v *FBox) Disable() *FBox {
	v.v.SetSensitive(false)
	return v
}
func (v *FBox) Enable() *FBox {
	v.v.SetSensitive(true)
	return v
}
func (v *FBox) Visible() *FBox {
	v.v.SetVisible(true)
	return v
}
func (v *FBox) Invisible() *FBox {
	v.FBaseView.Invisible()
	return v
}

func (v *FBox) Tooltips(s string) *FBox {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FBox) Focus() *FBox {
	v.FBaseView.Focus()
	return v
}
func (v *FBox) Padding(i uint) *FBox {
	v.padding = i
	return v
}

func (v *FBox) OnDragDrop(f func([]string)) *FBox {
	v.FBaseView.OnDragDrop(f)
	return v
}

//====================================================================
func (v *FBox) Append(is ...IView) *FBox {
	for _, i := range is {
		if i == nil {
			continue
		}
		v.v.PackStart(i.getBaseView().widget, i.getBaseView().expand, !i.getBaseView().notFill, i.getBaseView().padding)
	}
	return v
}