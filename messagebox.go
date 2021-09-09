package faithtop

var (
	newMessageBox_InfoImpl func(parent IWindow, title, content string, standardButtons, defaultButton StandardButton) StandardButton
)

func MessageBox_Info(parent IWindow, title, content string, standardButtons, defaultButton StandardButton) StandardButton {
	return newMessageBox_InfoImpl(parent, title, content, standardButtons, defaultButton)
}
