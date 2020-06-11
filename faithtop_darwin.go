package faithtop

import (
	"os"

	"github.com/therecipe/qt/widgets"
)

var app *widgets.QApplication

func init() {
	app = widgets.NewQApplication(len(os.Args), os.Args)
}

func MainQuit() {
	app.Quit()
}

func Main() {
	widgets.QApplication_Exec()
}
