package faithtop

import (
	"github.com/StevenZack/livedata"
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
	f.widget = f.v
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
	f.FBaseView.Expand()
	return f
}

func (f *FText) SetItemId(l *FListView, id string) *FText {
	l.vhs[l.currentCreation].vlist[id] = f
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

func (f *FText) BindText(t *livedata.String) *FText {
	t.ObserveForever(func(s string) {
		RunOnUIThread(func() {
			f.Text(s)
		})
	})
	return f
}

