package faithtop

import (
	"github.com/gofaith/walk"
	"github.com/gofaith/walk/declarative"
)

type FListView struct {
	FBaseView
	scroll          *FScroll
	vhs             []ViewHolder
	currentCreation int
	create          func(*FListView) IView
	bind            func(*ViewHolder, int)
	count           func() int
	gotCount        int
}

type ViewHolder struct {
	root  IView
	vlist map[string]IView
}

func (v *ViewHolder) RootView() IView {
	return v.root
}

func VListView(create func(*FListView) IView, bind func(*ViewHolder, int), count func() int) *FListView {
	f := &FListView{}
	f.scroll = VScroll()
	f.create = create
	f.bind = bind
	f.count = count
	return f.NotifyDataChange()
}

func (f *FListView) Assign(v **FListView) *FListView {
	*v = f
	return f
}

func (f *FListView) NotifyDataChange() *FListView {
	originSize := f.gotCount
	newSize := f.count()
	if newSize < 0 {
		newSize = 0
	}

	if newSize > originSize {
		for i := originSize; i < newSize; i++ {
			f.currentCreation = i
			if i >= len(f.vhs) {
				f.vhs = append(f.vhs, ViewHolder{vlist: make(map[string]IView)})
			}
			createdView := f.create(f)
			f.scroll.Append(createdView)
			f.vhs[i].root = createdView
		}
	} else {
		for i := newSize; i < originSize; i++ {
			f.scroll.v.Children().RemoveAt(0)
			f.vhs = f.vhs[1:len(f.vhs)]
		}
	}
	for i := 0; i < newSize; i++ {
		f.bind(&f.vhs[i], i)
	}
	f.gotCount = newSize
	return f
}

func (f *FListView) Size(w,h int) *FListView {
	f.scroll.Size(w,h)
	return f
}

func (f *FListView) baseView() *FBaseView {
	return &f.FBaseView
}

func (f *FListView) widget(builder *declarative.Builder) walk.Widget {
	if f.scroll.v == nil {
		f.scroll.dec.Create(builder)
	}
	return f.scroll.v
}

func (f *FListView) declarative() declarative.Widget {
	return f.scroll.declarative()
}
