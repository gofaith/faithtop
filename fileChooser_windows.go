package faithtop

import (
	"github.com/gonutz/wui"
	"github.com/mattn/go-gtk/gtk"
)

type FFileChooser struct {
	title                          string
	parent                         *gtk.Window
	multiple, folderMode, saveMode bool
	filter                         string
	onResult                       func(f string)
	onResults                      func(fs []string)
	showAfter                      bool
}

func FileChooser(w *FWindow) *FFileChooser {
	f := &FFileChooser{}
	f.parent = w.v
	return f
}
func ShowFileChooser(w *FWindow) *FFileChooser {
	return FileChooser(w).DeferShow()
}
func (f *FFileChooser) Filter(pattern string) *FFileChooser {
	f.filter = pattern
	return f
}
func (f *FFileChooser) TypeSelectFile(rp func(fname string)) *FFileChooser {
	f.multiple = false
	f.folderMode = false
	f.saveMode = false
	f.onResult = rp
	if f.showAfter {
		f.Show()
	}
	return f
}
func (f *FFileChooser) TypeSelectFiles(rp func(fs []string)) *FFileChooser {
	f.multiple = true
	f.folderMode = false
	f.saveMode = false
	f.onResults = rp
	if f.showAfter {
		f.Show()
	}
	return f
}
func (f *FFileChooser) TypeSelectFolder(rp func(dir string)) *FFileChooser {
	f.multiple = false
	f.folderMode = true
	f.saveMode = false
	f.onResult = rp
	if f.showAfter {
		f.Show()
	}
	return f
}
func (f *FFileChooser) TypeSelectFolders(rp func(dirs []string)) *FFileChooser {
	f.multiple = true
	f.folderMode = true
	f.saveMode = false
	f.onResults = rp
	if f.showAfter {
		f.Show()
	}
	return f
}
func (f *FFileChooser) TypeSaveFile(rp func(path string)) *FFileChooser {
	f.multiple = false
	f.folderMode = false
	f.saveMode = true
	f.onResult = rp
	if f.showAfter {
		f.Show()
	}
	return f
}
func (f *FFileChooser) TypeSaveFolder(rp func(dir string)) *FFileChooser {
	f.multiple = false
	f.folderMode = true
	f.saveMode = true
	f.onResult = rp
	if f.showAfter {
		f.Show()
	}
	return f
}
func (f *FFileChooser) Show() *FFileChooser {
	if f.saveMode {
		if f.folderMode {
			w := gtk.NewFileChooserDialog(
				f.title,
				f.parent,
				gtk.FILE_CHOOSER_ACTION_CREATE_FOLDER,
				gtk.STOCK_OK,
				gtk.RESPONSE_ACCEPT)
			if w.Run() == gtk.RESPONSE_ACCEPT {
				f.onResult(w.GetFilename())
			}
			w.Destroy()
		} else {
			w := wui.NewFileSaveDialog()
			w.SetTitle(f.title)
			b, s := w.Execute(nil)
			if b {
				f.onResult(s)
			}
		}
	} else { //open mode
		if f.folderMode {
			if f.multiple { //open folders
				w := gtk.NewFileChooserDialog(
					f.title,
					f.parent,
					gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER,
					gtk.STOCK_OK,
					gtk.RESPONSE_ACCEPT)
				w.SetSelectMultiple(true)
				if w.Run() == gtk.RESPONSE_ACCEPT {
					f.onResults(w.GetFilenames())
				}
				w.Destroy()
			} else { // open folder
				w := wui.NewFolderSelectDialog()
				w.SetTitle(f.title)
				b, s := w.Execute(nil)
				if b {
					f.onResult(s)
				}
			}
		} else { //open file
			if f.multiple {
				w := wui.NewFileOpenDialog()
				w.SetTitle(f.title)
				b, s := w.ExecuteMultiSelection(nil)
				if b {
					f.onResults(s)
				}
			} else {
				w := wui.NewFileOpenDialog()
				w.SetTitle(f.title)
				b, s := w.ExecuteSingleSelection(nil)
				if b {
					f.onResult(s)
				}
			}
		}
	}
	return f
}
func (f *FFileChooser) DeferShow() *FFileChooser {
	f.showAfter = true
	return f
}
func (f *FFileChooser) Close() *FFileChooser {
	return f
}
func (f *FFileChooser) Title(s string) *FFileChooser {
	f.title = s
	return f
}
