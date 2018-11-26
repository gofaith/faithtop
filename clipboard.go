package faithtop

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
)

type FClipboard struct {
	v *gtk.Clipboard
}

func Clipboard() *FClipboard {
	f := &FClipboard{}
	f.v = gtk.NewClipboardGetForDisplay(gdk.DisplayGetDefault(), gdk.SELECTION_CLIPBOARD)
	return f
}
func (f *FClipboard) GetText() string {
	return f.v.WaitForText()
}
func (f *FClipboard) SetText(t string) *FClipboard {
	f.v.SetText(t)
	return f
}
func (f *FClipboard) OnChange(fn func()) *FClipboard {
	f.v.Connect("owner-change", fn)
	return f
}
