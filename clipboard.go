package faithtop

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
)

var (
	clipboard *gtk.Clipboard
)

type FClipboard struct {
	onchangeC chan bool
}

func Clipboard() *FClipboard {
	f := &FClipboard{}
	if clipboard == nil {
		clipboard = gtk.NewClipboardGetForDisplay(gdk.DisplayGetDefault(), gdk.SELECTION_CLIPBOARD)
	}
	return f
}
func (f *FClipboard) GetText() string {
	return clipboard.WaitForText()
}
func (f *FClipboard) SetText(t string) *FClipboard {
	clipboard.SetText(t)
	return f
}

// func (f *FClipboard) OnChange(fn func()) *FClipboard {
// clipboard.Connect("owner-change", fn)
// return f
// }
