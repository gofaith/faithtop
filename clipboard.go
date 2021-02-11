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

func (f *FClipboard) OnChange(fn func(string)) *FClipboard {
	if fn == nil {
		return f
	}
	f.clipboard.Connect("owner-change", func() {
		fn(f.GetText())
	})
	return f
}
