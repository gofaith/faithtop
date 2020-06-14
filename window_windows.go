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
		Layout:   declarative.VBox{MarginsZero: true, SpacingZero: true},
	}
	f.Size(300, 300)
	if DefaultWindowIcon != nil {
		f.dec.Icon = DefaultWindowIcon
	}
	return f
}

func (f *FWindow) Title(s string) *FWindow {
	if f.w != nil {
		f.w.SetTitle(s)
		return f
	}
	f.dec.Title = s
	return f
}

func (f *FWindow) Size(w, h int) *FWindow {
	if f.w != nil {
		size := walk.Size{Width: w, Height: h}
		f.w.SetSize(size)
		f.w.SetMinMaxSize(size, size)
		return f
	}
	f.dec.Size.Width = w
	f.dec.Size.Height = h
	f.dec.MaxSize.Width = w
	f.dec.MaxSize.Height = h
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
	f.dec.Children = append(f.dec.Children, i.declarative())
	if f.showAfter {
		f.Show()
	}
	return f
}

func (f *FWindow) VBox(is ...IView) *FWindow {
	f.dec.Layout = declarative.VBox{SpacingZero: true, MarginsZero: true}
	for _, i := range is {
		f.dec.Children = append(f.dec.Children, i.declarative())
	}
	if f.showAfter {
		f.Show()
	}
	return f
}
