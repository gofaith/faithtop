package faithtop

import (
	"log"

	"github.com/gofaith/walk"
)

type IView interface {
	baseView() *FBaseView
}
type FBaseView struct {
	widget walk.Widget
	expand bool
}

func (f *FBaseView) baseView() *FBaseView {
	return f
}

func (f *FBaseView) IsEnabled() bool {
	return f.widget.Enabled()
}

func (f *FBaseView) IsVisible() bool {
	return f.widget.Visible()
}

func (f *FBaseView) GetX() int {
	return f.widget.X()
}

func (f *FBaseView) GetY() int {
	return f.widget.Y()
}

func (f *FBaseView) Focus() {
	e := f.widget.SetFocus()
	if e != nil {
		log.Println(e)
		return
	}
}

//
func (f *FBaseView) Expand()  {
	f.expand=true	
}