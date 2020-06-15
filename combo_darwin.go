package faithtop

import (
	"github.com/therecipe/qt/widgets"
)

type FCombo struct {
	FBaseView
	v      *widgets.QComboBox
	change func(string)
}

func Combo() *FCombo {
	f := &FCombo{
		v: widgets.NewQComboBox(nil),
	}
	f.widget = f.v
	f.v.ConnectCurrentTextChanged(func(text string) {
		if f.change != nil {
			f.change(text)
		}
	})
	return f
}

func (f *FCombo) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FCombo) Assign(v **FCombo) *FCombo {
	*v = f
	return f
}

func (f *FCombo) Disable() *FCombo {
	f.v.SetEnabled(false)
	return f
}

func (f *FCombo) Enable() *FCombo {
	f.v.SetEnabled(true)
	return f
}

// combo
func (f *FCombo) SetData(texts []string) *FCombo {
	f.v.Clear()
	f.v.AddItems(texts)
	return f
}

func (f *FCombo) OnChange(fn func(text string)) *FCombo {
	f.change = fn
	return f
}
