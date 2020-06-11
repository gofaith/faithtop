package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FFixed struct {
	FBaseView
	v *gtk.Fixed
}

func Fixed() *FFixed {
	f := &FFixed{}
	f.v = gtk.NewFixed()
	f.view = f.v
	f.widget = &f.v.Widget
	setupWidget(f)
	return f
}

// ================================================================
func (v *FFixed) Size(w, h int) *FFixed {
	v.FBaseView.Size(w, h)
	return v
}
func (f *FFixed) Assign(v **FFixed) *FFixed {
	*v = f
	return f
}
func (vh *ViewHolder) GetFixedByItemId(itemId string) *FFixed {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FFixed); ok {
			return lv
		}
	}
	return nil
}
func (v *FFixed) SetItemId(lv *FListView, itemId string) *FFixed {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}
func GetFixedById(id string) *FFixed {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FFixed); ok {
			return b
		}
	}
	return nil
}
func (v *FFixed) getBaseView() *FBaseView {
	return &v.FBaseView
}
func (v *FFixed) OnClick(f func()) *FFixed {
	v.v.Connect("clicked", f)
	return v
}
func (v *FFixed) SetId(id string) *FFixed {
	idMap[id] = v
	return v
}
func (v *FFixed) Expand() *FFixed {
	v.expand = true
	return v
}
func (v *FFixed) NotFill() *FFixed {
	v.notFill = true
	return v
}
func (v *FFixed) Disable() *FFixed {
	v.v.SetSensitive(false)
	return v
}
func (v *FFixed) Enable() *FFixed {
	v.v.SetSensitive(true)
	return v
}
func (v *FFixed) Visible() *FFixed {
	v.v.SetVisible(true)
	return v
}
func (v *FFixed) Invisible() *FFixed {
	v.FBaseView.Invisible()
	return v
}

func (v *FFixed) Tooltips(s string) *FFixed {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FFixed) Focus() *FFixed {
	v.FBaseView.Focus()
	return v
}
func (v *FFixed) Padding(i uint) *FFixed {
	v.padding = i
	return v
}

func (v *FFixed) OnDragDrop(f func([]string)) *FFixed {
	v.FBaseView.OnDragDrop(f)
	return v
}

//====================================================================
func (v *FFixed) Append(is ...IView) *FFixed {
	for _, i := range is {
		v.v.Add(i.getBaseView().widget)
		var iv = i
		v.v.Connect("size_allocate", func() {
			a := v.v.GetAllocation()
			iv.getBaseView().widget.SetSizeRequest(a.Width, a.Height)
		})
	}
	return v
}
