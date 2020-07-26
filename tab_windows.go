package faithtop

import (
	"github.com/gofaith/walk"
	"github.com/gofaith/walk/declarative"
)

type FTabLayout struct {
	FBaseView
	v   *walk.TabWidget
	dec declarative.TabWidget
}

func TabLayout() *FTabLayout {
	f := &FTabLayout{}
	f.dec = declarative.TabWidget{
		AssignTo: &f.v,
	}
	return f
}

func (f *FTabLayout) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FTabLayout) widget(builder *declarative.Builder) walk.Widget {
	if f.v == nil {
		f.dec.Create(builder)
	}
	return f.v
}

func (f *FTabLayout) declarative() declarative.Widget {
	return f.dec
}

func (f *FTabLayout) Tabs(tabs ...*FTab) *FTabLayout {
	for _, tab := range tabs {
		f.dec.Pages = append(f.dec.Pages, declarative.TabPage{
			Title:  tab.title,
			Layout: declarative.VBox{SpacingZero: true, MarginsZero: true},
			Children: []declarative.Widget{
				tab.content.declarative(),
			},
		})
	}
	return f
}

func (f *FTabLayout) Expand() *FTabLayout {
	f.expand=true
	return f
}

type FTab struct {
	title   string
	content IView
}

func Tab(title string, view IView) *FTab {
	return &FTab{
		title:   title,
		content: view,
	}
}
