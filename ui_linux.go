package faithtop

import (
	"strings"
	"unsafe"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

type IView interface {
	getBaseView() *FBaseView
}
type FBaseView struct {
	view                  gtk.IWidget
	padding               uint
	widget                *gtk.Widget
	initWidth, initHeight int
	expand, notFill       bool
}

func setupWidget(f IView) {
	// f.getBaseView().widget.SetCanFocus(true)
	f.getBaseView().initWidth = -56
	f.getBaseView().initHeight = -56
}
func RunOnUIThread(f func()) {
	glib.IdleAdd(f)
}
func (v *FBaseView) Size(w, h int) {
	if w <= 0 {
		w = -1
	}
	if h <= 0 {
		h = -1
	}
	v.widget.SetSizeRequest(w, h)
}
func (v *FBaseView) IsEnabled() bool {
	return v.widget.GetSensitive()
}
func (v *FBaseView) IsVisible() bool {
	return v.widget.GetVisible()
}
func (v *FBaseView) GetX() int {
	return v.widget.GetAllocation().X
}
func (v *FBaseView) GetY() int {
	return v.widget.GetAllocation().Y
}
func (v *FBaseView) Focus() {
	v.widget.GrabFocus()
}
// func (v *FBaseView) OnEnter(f func()) {
// 	v.widget.Connect("activate", f)
// }
func (v *FBaseView) Invisible() {
	RunOnUIThread(func() {
		v.widget.SetVisible(false)
	})
}

//fix fs list empty
func (v *FBaseView) OnDragDrop(f func([]string)) {
	targets := []gtk.TargetEntry{
		{"text/uri-list", 0, 0},
		{"STRING", 0, 1},
		{"text/plain", 0, 2},
	}
	v.widget.DragDestSet(
		gtk.DEST_DEFAULT_MOTION|
			gtk.DEST_DEFAULT_HIGHLIGHT|
			gtk.DEST_DEFAULT_DROP,
		targets,
		gdk.ACTION_COPY)
	v.widget.DragDestAddUriTargets()
	v.widget.Connect("drag-data-received", func(ctx *glib.CallbackContext) {
		sdata := gtk.NewSelectionDataFromNative(unsafe.Pointer(ctx.Args(3)))
		if sdata != nil {
			a := (*[2000]uint8)(sdata.GetData())
			names := string(a[0 : sdata.GetLength()-1])
			files := strings.Split(names, "\n")
			var fs []string
			for i := range files {
				filename, _, _ := glib.FilenameFromUri(files[i])
				str := strings.Replace(filename, "\r", "", -1)
				if str == "" {
					continue
				}
				fs = append(fs, str)
			}
			f(fs)
		}
	})
}
