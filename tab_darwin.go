package faithtop

import (
	"github.com/therecipe/qt/widgets"
)

type FTabLayout struct {
	FBaseView
	v *widgets.QTabWidget
}
type FTab struct {
	title   string
	content IView
}

func TabLayout() *FTabLayout {
	f := &FTabLayout{
		v: widgets.NewQTabWidget(nil),
	}
	f.widget = f.v
	return f
}

func Tab(title string, v IView) *FTab {
	f := &FTab{title: title, content: v}
	return f
}

func (f *FTabLayout) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FTabLayout) Assign(v **FTabLayout) *FTabLayout {
	*v = f
	return f
}

func (f *FTabLayout) Expand() *FTabLayout {
	f.FBaseView.Expand()
	return f
}

func (f *FTabLayout) Tabs(ts ...*FTab) *FTabLayout {
	for _, t := range ts {
		f.v.AddTab(t.content.baseView().widget, t.title)
	}
	return f
}


func (f *FTabLayout) OnSwitchPage(fn func())  *FTabLayout{
	f.v.ConnectCurrentChanged(func(i int){
		if fn!=nil{
			fn()
		}
	})
	return f
}

func (f *FTabLayout) CurrentIndex() int {
	return f.v.CurrentIndex()
}