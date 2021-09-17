package faithtop

type IApp interface {
	Run() int
	Quit()
	SetQuickStyle(quickStyle QuickStyle)
	SetClipboardText(s string) IApp
	GetClipboardText() string
	OnClipboardTextChanged(fn func(self IApp, s string)) IApp
	SetQuitOnLastWindowClosed(b bool) IApp
	RunOnUIThread(fn func())
}

type QuickStyle string

const (
	QuickStyle_Default   QuickStyle = "Default"
	QuickStyle_Material  QuickStyle = "Material"
	QuickStyle_Fusion    QuickStyle = "Fusion"
	QuickStyle_Imagine   QuickStyle = "Imagine"
	QuickStyle_Universal QuickStyle = "Universal"
)

const (
	MAX_SIZE = 16777215
)

var appImpl func() IApp

func NewApp() IApp {
	if appImpl == nil {
		panic("please build with impl tag: qtdeploy -tags=impl build desktop")
	}
	return appImpl()
}
