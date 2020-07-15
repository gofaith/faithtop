package faithtop

import (
	"sync"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
)

type FClipboard struct {
	clipboard *gtk.Clipboard
}

var (
	clipboardListener       sync.Once
	clipboardChangeHandlers []func(string)
)

func Clipboard() *FClipboard {
	f := &FClipboard{}
	f.clipboard = gtk.NewClipboardGetForDisplay(gdk.DisplayGetDefault(), gdk.SELECTION_CLIPBOARD)
	return f
}
func (f *FClipboard) GetText() string {
	return f.clipboard.WaitForText()
}
func (f *FClipboard) SetText(t string) *FClipboard {
	f.clipboard.SetText(t)
	return f
}

// func (f *FClipboard) OnChange(fn func(str string)) *FClipboard {
// 	clipboardListener.Do(listenClipboard)
// 	clipboardChangeHandlers = append(clipboardChangeHandlers, fn)
// 	if fn != nil {
// 		fn(f.GetText())
// 	}
// 	return f
// }

func (f *FClipboard) OnChange(fn func(string)) *FClipboard {
	if fn == nil {
		return f
	}
	f.clipboard.Connect("owner-change", func() {
		fn(f.GetText())
	})
	return f
}

// func listenClipboard() {
// 	c := Clipboard()
// 	text := c.GetText()
// 	go func() {
// 		for {
// 			time.Sleep(time.Millisecond * 100)
// 			cb := c.GetText()
// 			if cb != text {
// 				text = cb
// 				for _, fn := range clipboardChangeHandlers {
// 					if fn != nil {
// 						fn(text)
// 					}
// 				}
// 			}
// 		}
// 	}()
// }
