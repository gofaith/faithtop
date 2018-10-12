package faithtop

import "github.com/mattn/go-gtk/gtk"

type FBox struct {
	FBaseView
	v           *gtk.Box
	orientation bool //false :vertical , true:horizontal
}

func (v *FBox) getBaseView() *FBaseView {
	return &v.FBaseView
}

func VBox() *FBox {
	f := &FBox{}
	f.v = &gtk.NewVBox(false, 0).Box
	f.view = f.v
	f.widget = &f.v.Widget
	f.orientation = false
	setupWidget(f)
	return f
}

func HBox() *FBox {
	f := &FBox{}
	f.v = &gtk.NewHBox(false, 0).Box
	f.view = f.v
	f.widget = &f.v.Widget
	f.orientation = true
	setupWidget(f)
	return f
}

func (v *FBox) Size(w, h int) *FBox {
	v.FBaseView.Size(w, h)
	return v
}
func (v *FBox) Append(is ...IView) *FBox {
	var fs []func()
	for _, i := range is {
		v.v.PackStart(i.getBaseView().widget, i.getBaseView().expand, !i.getBaseView().notFill, i.getBaseView().padding)
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
