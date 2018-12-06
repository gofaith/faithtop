package faithtop

import (
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gtk"
)

var (
	clipboard     *gtk.Clipboard
	clipboardText string
)

type FClipboard struct {
}

func Clipboard() *FClipboard {
	if clipboard == nil {
		clipboard = gtk.NewClipboardGetForDisplay(gdk.DisplayGetDefault(), gdk.SELECTION_CLIPBOARD)
	}
	f := &FClipboard{}
	RunOnUIThread(func() {
		clipboard.Connect("owner-change", func() {
			clipboardText = clipboard.WaitForText()
		})
		clipboardText = clipboard.WaitForText()
	})
	return f
}
func (f *FClipboard) GetText() string {
	return clipboardText
}
func (f *FClipboard) SetText(t string) *FClipboard {
	RunOnUIThread(func() {
		clipboardText = t
		clipboard.SetText(t)
	})
	return f
}
func (f *FClipboard) OnChange(fn func()) *FClipboard {
	clipboard.Connect("owner-change", func() {
		go fn()
	})
	return f
}
