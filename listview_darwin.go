package faithtop

type FListView struct {
	FScroll
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
	f.FScroll = *VScroll()
	f.create = create
	f.bind = bind
	f.count = count
	return f.NotifyDataChange()
}

//base
func (f *FListView) Assign(v **FListView) *FListView {
	*v = f
	return f
}

// list

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
			createdView.baseView().Visible()
			f.Append(createdView)
			f.vhs[i].root = createdView
		}
	} else {
		for i := newSize; i < originSize; i++ {
			f.vhs[i].root.baseView().Invisible()
		}
	}

	for i := 0; i < newSize; i++ {
		f.bind(&f.vhs[i], i)
	}
	f.gotCount = newSize
	return f
}

func (v *ViewHolder) GetButton(id string) *FButton {
	return v.vlist[id].(*FButton)
}

func (v *ViewHolder) GetText(id string) *FText {
	return v.vlist[id].(*FText)
}
