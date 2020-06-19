package faithtop

import (
	"github.com/StevenZack/livedata"
	"github.com/therecipe/qt/widgets"
)

type FTextArea struct {
	FBaseView
	v        *widgets.QTextEdit
	onChange func(string)
}

func TextArea() *FTextArea {
	f := &FTextArea{
		v: widgets.NewQTextEdit(nil),
	}

	f.widget = f.v
	f.v.ConnectTextChanged(func() {
		if f.onChange == nil {
			return
		}
		f.onChange(f.GetText())
	})
	return f
}

func (f *FTextArea) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FTextArea) Assign(v **FTextArea) *FTextArea {
	*v = f
	return f
}

func (f *FTextArea) Expand() *FTextArea {
	f.FBaseView.Expand()
	return f
}

func (f *FTextArea) Size(w, h int) *FTextArea {
	f.FBaseView.Size(w, h)
	return f
}

func (f *FTextArea) Text(text string) *FTextArea {
	f.v.SetText(text)
	return f
}

func (f *FTextArea) OnChange(fn func(text string)) *FTextArea {
	f.onChange = fn
	return f
}

func (f *FTextArea) GetText() string {
	return f.v.ToPlainText()
}

func (f *FTextArea) BindText(str *livedata.String) *FTextArea {
	f.Text(str.Get())
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

func (f *FTextArea) SetItemId(l *FListView, id string) *FTextArea {
	l.vhs[l.currentCreation].vlist[id] = f
	return f
}

func (f *FTextArea) Hint(text string) *FTextArea {
	f.v.SetPlaceholderText(text)
	return f
}

func (f *FTextArea) Disable() *FTextArea {
	f.v.SetEnabled(false)
	return f
}
func (f *FTextArea) Enable() *FTextArea {
	f.v.SetEnabled(true)
	return f
}
