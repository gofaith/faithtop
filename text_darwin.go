package faithtop

import (
	"github.com/therecipe/qt/widgets"
)

type FText struct {
	FBaseView
	v *widgets.QLabel
}

func Text(text string) *FText {
	f := &FText{
		v: widgets.NewQLabel2(text, nil, 0),
	}
	
	return f
}

// base
func (f *FText) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FText) Assign(v **FText) *FText {
	*v = f
	return f
}

func (f *FText) Expand() *FText {
	f.FbaseView().expand(true)
	return f
}

// text
func (f *FText) Text(s string) *FText {
	f.v.SetText(s)
	return f
}

func (f *FText) GetText() string {
	return f.v.Text()
}
