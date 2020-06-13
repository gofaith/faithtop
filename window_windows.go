package faithtop

import (
	"log"

	"github.com/gofaith/walk"
	"github.com/gofaith/walk/declarative"
)

type FWindow struct {
	w            *walk.MainWindow
	dec          declarative.MainWindow
	showAfter    bool
	closeClicked func() bool
}

var DefaultWindowIcon *walk.Icon

func Win() *FWindow {
	f := &FWindow{}
	f.dec = declarative.MainWindow{
		AssignTo: &f.w,
		Size:     declarative.Size{Width: 300, Height: 300},
		Layout:   declarative.VBox{MarginsZero: true, SpacingZero: true},
	}
	if DefaultWindowIcon != nil {
		f.dec.Icon = DefaultWindowIcon
	}
	return f
}

func (f *FWindow) Title(s string) *FWindow {
	f.dec.Title = s
	return f
}

func (f *FWindow) Size(w, h int) *FWindow {
	f.dec.Size.Width = w
	f.dec.Size.Height = h
	return f
}

func (f *FWindow) Show() *FWindow {
	currentWindow = f
	if f.w != nil {
		f.w.Show()
		return f
	}
	e := f.dec.Create()
	if e != nil {
		log.Println(e)
		log.Fatal(e)
	}

	f.w.Closing().Attach(func(b *bool, _ walk.CloseReason) {
		if f.closeClicked == nil {
			*b = false
			return
		}
		*b = f.closeClicked()
	})
	return f
}

func (f *FWindow) Close() {
	e := f.w.Close()
	if e != nil {
		log.Println(e)
		return
	}
}

func (f *FWindow) Hide() {
	f.w.Hide()
}

func (f *FWindow) IsVisible() bool {
	return f.w.Visible()
}

func (f *FWindow) DeferShow() *FWindow {
	f.showAfter = true
	return f
}

func (f *FWindow) OnCloseClicked(fn func() bool) *FWindow {
	f.closeClicked = fn
	return f
}

func (f *FWindow) Add(i IView) *FWindow {
	f.dec.Children = []declarative.Widget{i.declarative()}
	if f.showAfter {
		f.Show()
	}
	return f
}
