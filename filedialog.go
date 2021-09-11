package faithtop

type IFileDialog interface {
	FileMode(fileMode FileMode) IFileDialog
	NameFilter(pattern string) IFileDialog
	Directory(dir string) IFileDialog
	AcceptMode(acceptMode AcceptMode) IFileDialog
	Options(options FileDialogOption) IFileDialog
	Show() IFileDialog
	SelectedFiles() []string
}

type FileMode int64

const (
	FileMode_AnyFile       FileMode = FileMode(0)
	FileMode_ExistingFile  FileMode = FileMode(1)
	FileMode_Directory     FileMode = FileMode(2)
	FileMode_ExistingFiles FileMode = FileMode(3)
	FileMode_DirectoryOnly FileMode = FileMode(4)
)

type AcceptMode int64

const (
	AcceptMode_AcceptOpen AcceptMode = AcceptMode(0)
	AcceptMode_AcceptSave AcceptMode = AcceptMode(1)
)

type FileDialogOption int64

const (
	FileDialogOption_ShowDirsOnly                FileDialogOption = FileDialogOption(0x00000001)
	FileDialogOption_DontResolveSymlinks         FileDialogOption = FileDialogOption(0x00000002)
	FileDialogOption_DontConfirmOverwrite        FileDialogOption = FileDialogOption(0x00000004)
	FileDialogOption_DontUseSheet                FileDialogOption = FileDialogOption(0x00000008)
	FileDialogOption_DontUseNativeDialog         FileDialogOption = FileDialogOption(0x00000010)
	FileDialogOption_ReadOnly                    FileDialogOption = FileDialogOption(0x00000020)
	FileDialogOption_HideNameFilterDetails       FileDialogOption = FileDialogOption(0x00000040)
	FileDialogOption_DontUseCustomDirectoryIcons FileDialogOption = FileDialogOption(0x00000080)
)

var (
	newFileDialogImpl  func(parent IWindow) IFileDialog
	saveFileDialogImpl func(parent IWindow, title, defaultName, filter, selectedFilter string, options FileDialogOption) string
)

func FileDialog(parent IWindow) IFileDialog {
	return newFileDialogImpl(parent)
}

func SaveFile(parent IWindow, title, defaultName, filter, selectedFilter string, options FileDialogOption) string {
	return saveFileDialogImpl(parent, title, defaultName, filter, selectedFilter, options)
}
