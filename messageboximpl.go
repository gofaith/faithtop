//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/widgets"
)

func init() {
	newMessageBox_InfoImpl = func(parent IWindow, title, content string, buttons, defaultButton StandardButton) StandardButton {
		return StandardButton(widgets.QMessageBox_Information(parent.(*WindowImpl).window, title, content, widgets.QMessageBox__StandardButton(buttons), widgets.QMessageBox__StandardButton(defaultButton)))
	}
}
