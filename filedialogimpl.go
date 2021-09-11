//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type FileDialogImpl struct {
	dialog *widgets.QFileDialog
}

func init() {
	newFileDialogImpl = func(parent IWindow) IFileDialog {
		return &FileDialogImpl{
			dialog: widgets.NewQFileDialog(parent.(*WindowImpl).window, core.Qt__Dialog),
		}
	}
	saveFileDialogImpl = func(parent IWindow, title, defaultName, filter, selectedFilter string, options FileDialogOption) string {
		var dialog *widgets.QFileDialog
		return dialog.GetSaveFileName(parent.(*WindowImpl).window, title, defaultName, filter, selectedFilter, widgets.QFileDialog__Option(options))
	}
}
func (f *FileDialogImpl) FileMode(mode FileMode) IFileDialog {
	f.dialog.SetFileMode(widgets.QFileDialog__FileMode(mode))
	return f
}

func (f *FileDialogImpl) NameFilter(pattern string) IFileDialog {
	f.dialog.SetNameFilter(pattern)
	return f
}

func (f *FileDialogImpl) Directory(dir string) IFileDialog {
	f.dialog.SetDirectory(dir)
	return f
}

func (f *FileDialogImpl) AcceptMode(acceptMode AcceptMode) IFileDialog {
	f.dialog.SetAcceptMode(widgets.QFileDialog__AcceptMode(acceptMode))
	return f
}

func (f *FileDialogImpl) Options(options FileDialogOption) IFileDialog {
	f.dialog.SetOptions(widgets.QFileDialog__Option(options))
	return f
}
func (f *FileDialogImpl) Show() IFileDialog {
	f.dialog.Exec()
	return f
}
func (f *FileDialogImpl) SelectedFiles() []string {
	return f.dialog.SelectedFiles()
}
