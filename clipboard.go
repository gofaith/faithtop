package faithtop

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
)

type FClipboard struct {
	clipboard *gtk.Clipboard
}

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

// unstable (causing panic )
// func (f *FClipboard) OnChange(fn func()) *FClipboard {
// clipboard.Connect("owner-change", fn)
// return f
// }
