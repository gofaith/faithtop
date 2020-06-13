package faithtop

import (
	"github.com/therecipe/qt/widgets"
)

type FSpace struct {
	FBaseView
	v *widgets.QWidget
}

func Space() *FSpace {
	f := &FSpace{
		v: widgets.NewQWidget(nil, 0),
	}
	f.widget = f.v
	f.Expand(true)
	return f
}

func (f *FSpace) baseView() *FBaseView {
	return &f.FBaseView
}
