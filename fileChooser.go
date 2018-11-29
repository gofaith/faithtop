package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FFileChooser struct {
	v         *gtk.FileChooserDialog
	showAfter bool
}

func FileChooser(w *FWindow) *FFileChooser {
	f := &FFileChooser{}
	f.v = gtk.NewFileChooserDialog(
		"Choose FIle...", w.v,
		gtk.FILE_CHOOSER_ACTION_OPEN,
		gtk.STOCK_OK,
		gtk.RESPONSE_ACCEPT)
	return f
}
func ShowFileChooser(w *FWindow)*FFileChooser{
	return FileChooser(w).DeferShow()
}
func (f *FFileChooser) Filter(pattern string) *FFileChooser {
	filter := gtk.NewFileFilter()
	filter.AddPattern(pattern)
	f.v.AddFilter(filter)
	return f
}
func (f *FFileChooser) TypeSelectFile(rp func(fname string)) *FFileChooser {
	f.v.SetAction(gtk.FILE_CHOOSER_ACTION_OPEN)
	f.v.SetSelectMultiple(false)
	f.v.Response(func() {
		rp(f.v.GetFilename())
		f.v.Destroy()
	})
	if f.showAfter{
		f.Show()
	}
	return f
}
func (f *FFileChooser) TypeSelectFiles(rp func(fs []string)) *FFileChooser {
	f.v.SetAction(gtk.FILE_CHOOSER_ACTION_OPEN)
	f.v.SetSelectMultiple(true)
	f.v.Response(func() {
		rp(f.v.GetFilenames())
		f.v.Destroy()
	})
	if f.showAfter{
		f.Show()
	}
	return f
}
func (f *FFileChooser) TypeSelectFolder(rp func(dir string)) *FFileChooser {
	f.v.SetAction(gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER)
	f.v.SetSelectMultiple(false)
	f.v.Response(func() {
		rp(f.v.GetFilename())
		f.v.Destroy()
	})
	if f.showAfter{
		f.Show()
	}
	return f
}
func (f *FFileChooser) TypeSelectFolders(rp func(dirs []string)) *FFileChooser {
	f.v.SetAction(gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER)
	f.v.SetSelectMultiple(true)
	f.v.Response(func() {
		rp(f.v.GetFilenames())
		f.v.Destroy()
	})
	if f.showAfter{
		f.Show()
	}
	return f
}
func (f *FFileChooser) TypeSaveFile(rp func(path string)) *FFileChooser {
	f.v.SetAction(gtk.FILE_CHOOSER_ACTION_SAVE)
	f.v.Response(func() {
		rp(f.v.GetFilename())
		f.v.Destroy()
	})
	if f.showAfter{
		f.Show()
	}
	return f
}
func (f *FFileChooser) TypeSaveFolder(rp func(dir string)) *FFileChooser {
	f.v.SetAction(gtk.FILE_CHOOSER_ACTION_CREATE_FOLDER)
	f.v.Response(func() {
		rp(f.v.GetFilename())
		f.v.Destroy()
	})
	if f.showAfter{
		f.Show()
	}
	return f
}
func (f *FFileChooser) Show() *FFileChooser {
	f.v.Run()
	return f
}
func (f *FFileChooser) DeferShow() *FFileChooser {
	f.showAfter=true
	return f
}
func (f *FFileChooser) Close() *FFileChooser {
	f.v.Destroy()
	return f
}