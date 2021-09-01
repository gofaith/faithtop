//go:build !dev

package faithtop

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

type AppImpl struct {
	app *widgets.QApplication
}

func init() {
	appImpl = func() IApp {
		return &AppImpl{
			app: widgets.NewQApplication(len(os.Args), os.Args),
		}
	}
}

func (a *AppImpl) Run() int {
	return a.app.Exec()
}
