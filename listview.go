package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FListView struct {
	FBaseView
	v               *gtk.ScrolledWindow
	lv              *gtk.Box
	vhs             []ViewHolder
	currentCreation int
	createView      func(*FListView) IView
	bindData        func(*ViewHolder, int)
	getCount        func() int
}
type ViewHolder struct {
	vlist map[string]IView
}

// func VlistView(createView func(*FListView) IView, bindData func(*ViewHolder, int), getCount func() int) *FListView {
// 	fb := &FListView{}
// 	v := gtk.NewScrolledWindow(nil, nil)
// 	fb.v = v
// 	fb.view = v
// 	fb.widget = &v.Widget
// 	fb.createView = createView
// 	fb.bindData = bindData
// 	fb.getCount = getCount
// 	fb.v.SetPolicy(gtk.POLICY_NEVER, gtk.POLICY_AUTOMATIC)
// 	fb.lv = &gtk.NewVBox(false, 0).Box
// 	fb.v.Add(fb.lv)
// 	for i := 0; i < getCount(); i++ {
// 		fb.currentCreation = i
// 		fb.vhs = append(fb.vhs, ViewHolder{vlist: make(map[string]IView)})
// 		fb.lv.Add(createView(fb).getBaseView().view)
// 	}
// 	setupWidget(fb)
// 	return fb
// }
