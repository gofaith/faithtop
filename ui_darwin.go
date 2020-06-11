package faithtop

import (
	"github.com/therecipe/qt/widgets"
)

type IView interface {
	baseView() *FBaseView
}

type FBaseView struct {
	view widgets.QWidget_ITF
}

