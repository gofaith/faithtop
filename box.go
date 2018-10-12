package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FBox struct {
	FBaseView
	v *gtk.Box
}

func VBox() *FBox {
	v := gtk.NewVBox(false, 0)
	fb := &FBox{}
	fb.v = &v.Box
	fb.view = v
	fb.widget = &v.Widget
	setupWidget(fb)
	return fb
}
func HBox() *FBox {
	v := gtk.NewHBox(false, 0)
	fb := &FBox{}
	fb.v = &v.Box
	fb.view = v
	fb.widget = &v.Widget
	setupWidget(fb)
	return fb
}

// ---------------------------------------------------------

func (v *FBox) getBaseView() *FBaseView {
	return &v.FBaseView
}

func (v *FBox) SetId(id string) *FBox {
	idMap[id] = v
	return v
}

// ---------------------------------------------------------

func (v *FBox) Append(is ...IView) *FBox {
	var fs []func()
	for _, i := range is {
		v.v.Add(i.getBaseView().view)
		i.getBaseView().alreadyAdded = true
		if i.getBaseView().afterAppend != nil {
			fs = append(fs, i.getBaseView().afterAppend)
			i.getBaseView().afterAppend = nil
		}
	}
	for _, f := range fs {
		f()
	}
	return v
}
