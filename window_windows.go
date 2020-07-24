package faithtop

import (
	"log"

	"github.com/StevenZack/livedata"

	"github.com/gen2brain/beeep"

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
		f.w.SetMinMaxSizePixels(size, size)
		return f
	}
	f.dec.Size.Width = w
	f.dec.Size.Height = h
	f.dec.MaxSize.Width = w
	f.dec.MaxSize.Height = h
	f.dec.MinSize.Width = w
	f.dec.MinSize.Height = h
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

func ShowInfo(w *FWindow, title, info string) {
	walk.MsgBox(w.w, title, info, walk.MsgBoxApplModal|walk.MsgBoxOK)
}

func ShowConfirm(w *FWindow, title, info, ok, cancel string, onYes, onNo func()) {
	i := walk.MsgBox(w.w, title, info, walk.MsgBoxApplModal|walk.MsgBoxYesNo)
	if i == 6 {
		if onYes != nil {
			onYes()
		}
	} else if i == 7 {
		if onNo != nil {
			onNo()
		}
	}
}

func ShowPrompt(w *FWindow, title, label, defaultText, ok, cancel string, onSubmit func(string)) {
	var dlg *walk.Dialog
	var edit *walk.LineEdit
	declarative.Dialog{
		AssignTo: &dlg,
		Title:    title,
		Layout:   declarative.VBox{},
		Children: []declarative.Widget{
			declarative.Label{
				Text: label,
			},
			declarative.LineEdit{
				AssignTo: &edit,
				Text:     defaultText,
				OnKeyUp: func(key walk.Key) {
					if key == walk.KeyReturn {
						if onSubmit != nil {
							onSubmit(edit.Text())
						}
						dlg.Close(0)
					}
				},
			},
			declarative.Composite{
				Layout: declarative.HBox{SpacingZero: true, MarginsZero: true},
				Children: []declarative.Widget{
					declarative.PushButton{
						Text: ok,
						OnClicked: func() {
							if onSubmit != nil {
								onSubmit(edit.Text())
							}
							dlg.Close(0)
						},
					},
					declarative.PushButton{
						Text: cancel,
						OnClicked: func() {
							dlg.Close(0)
						},
					},
				},
			},
		},
	}.Run(w.w)
}

func ShowErr(w *FWindow, title, err string) {
	walk.MsgBox(w.w, title, err, walk.MsgBoxApplModal|walk.MsgBoxOK|walk.MsgBoxIconError)
}

func ShowToast(title, info string) {
	beeep.Notify(title, info, "")
}

func (f *FWindow) BindInfo(title string, l *livedata.String) *FWindow {
	l.ObserveForever(func(s string) {
		if s != "" {
			ShowInfo(f, title, s)
		}
	})
	return f
}
