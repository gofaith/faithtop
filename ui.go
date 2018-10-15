package faithtop

import (
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
	afterShownFn          func()
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
func (v *FBaseView) GetWidth() int {
	return v.widget.GetAllocation().Width
}
func (v *FBaseView) GetHeight() int {
	return v.widget.GetAllocation().Height
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
