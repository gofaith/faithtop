package faithtop

import (
	"github.com/StevenZack/livedata"
	"github.com/therecipe/qt/widgets"
)

type FEdit struct {
	FBaseView
	v        *widgets.QLineEdit
	onChange func(string)
}

func Edit() *FEdit {
	f := &FEdit{
		v: widgets.NewQLineEdit(nil),
	}
	f.widget = f.v
	f.v.ConnectTextChanged(func(text string) {
		if f.onChange != nil {
			f.onChange(text)
		}
	})
	return f
}

//base
func (f *FEdit) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FEdit) Assign(v **FEdit) *FEdit {
	*v = f
	return f
}

func (f *FEdit) Expand() *FEdit {
	f.FBaseView.Expand()
	return f
}

func (f *FEdit) Size(w, h int) *FEdit {
	f.FBaseView.Size(w, h)
	return f
}

//edit
func (f *FEdit) Text(text string) *FEdit {
	f.v.SetText(text)
	return f
}

func (f *FEdit) OnChange(fn func(text string)) *FEdit {
	f.onChange = fn
	return f
}

func (f *FEdit) GetText() string {
	return f.v.Text()
}

func (f *FEdit) OnEnter(fn func()) *FEdit {
	f.v.ConnectReturnPressed(fn)
	return f
}

func (f *FEdit) BindText(str *livedata.String) *FEdit {
	str.ObserveForever(func(s string) {
		if f.v.IsVisible() {
			if s == f.GetText() {
				return
			}
			RunOnUIThread(func() {
				f.Text(s)
			})
		}
	})
	f.OnChange(func(text string) {
		if str.Get() == text {
			return
		}
		str.Post(text)
	})
	return f
}
