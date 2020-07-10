package faithtop

type FListView struct {
	FScroll
	vhs             []ViewHolder
	currentCreation int
	createView      func(*FListView) IView
	bindData        func(*ViewHolder, int)
	getCount        func() int
	gotCount        int
}

type ViewHolder struct {
	root  IView
	vlist map[string]IView
}

func (v *ViewHolder) RootView() IView {
	return v.root
}
func VListView(createView func(*FListView) IView, bindData func(*ViewHolder, int), getCount func() int) *FListView {
	fb := &FListView{}
	fb.FScroll = *VScroll()
	fb.createView = createView
	fb.bindData = bindData
	fb.getCount = getCount
	fb.gotCount = getCount()
	for i := 0; i < fb.gotCount; i++ {
		fb.currentCreation = i
		fb.vhs = append(fb.vhs, ViewHolder{vlist: make(map[string]IView)})
		createdView := createView(fb)
		fb.Append(createdView)
		fb.vhs[i].root = createdView
	}
	fb.execBindData()
	return fb
}
func HListView(createView func(*FListView) IView, bindData func(*ViewHolder, int), getCount func() int) *FListView {
	fb := &FListView{}
	fb.FScroll = *HScroll()
	fb.createView = createView
	fb.bindData = bindData
	fb.getCount = getCount
	fb.gotCount = getCount()
	for i := 0; i < getCount(); i++ {
		fb.currentCreation = i
		fb.vhs = append(fb.vhs, ViewHolder{vlist: make(map[string]IView)})
		createdView := createView(fb)
		fb.Append(createdView)
		fb.vhs[i].root = createdView
	}
	fb.execBindData()
	return fb
}

// ================================================================
func (v *FListView) Size(w, h int) *FListView {
	v.FBaseView.Size(w, h)
	return v
}
func (f *FListView) Assign(v **FListView) *FListView {
	*v = f
	return f
}
func (vh *ViewHolder) GetListViewByItemId(itemId string) *FListView {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FListView); ok {
			return lv
		}
	}
	return nil
}
func (v *FListView) SetItemId(lv *FListView, itemId string) *FListView {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}
func GetListViewById(id string) *FListView {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FListView); ok {
			return b
		}
	}
	return nil
}
func (v *FListView) baseView() *FBaseView {
	return &v.FBaseView
}

func (v *FListView) OnClick(f func()) *FListView {
	v.v.Connect("clicked", f)
	return v
}
func (v *FListView) SetId(id string) *FListView {
	idMap[id] = v
	return v
}
func (v *FListView) Expand() *FListView {
	v.expand = true
	return v
}
func (v *FListView) Disable() *FListView {
	v.v.SetSensitive(false)
	return v
}
func (v *FListView) Enable() *FListView {
	v.v.SetSensitive(true)
	return v
}
func (v *FListView) Visible() *FListView {
	v.v.SetVisible(true)
	return v
}
func (v *FListView) Invisible() *FListView {
	v.FBaseView.Invisible()
	return v
}

func (v *FListView) Tooltips(s string) *FListView {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FListView) Focus() *FListView {
	v.FBaseView.Focus()
	return v
}

func (v *FListView) OnDragDrop(f func([]string)) *FListView {
	v.FBaseView.OnDragDrop(f)
	return v
}

//====================================================================
func (v *FListView) execBindData() *FListView {
	for i := 0; i < v.getCount(); i++ {
		v.bindData(&v.vhs[i], i)
	}
	return v
}
func (fb *FListView) OnDataSetChanged() *FListView {
	origin_size := fb.gotCount
	new_size := fb.getCount()
	if new_size > origin_size {
		for i := origin_size; i < new_size; i++ {
			fb.currentCreation = i
			if i >= len(fb.vhs) {
				fb.vhs = append(fb.vhs, ViewHolder{vlist: make(map[string]IView)})
			}
			createdView := fb.createView(fb)
			createdView.baseView().widget.ShowAll()
			fb.Append(createdView)
			fb.vhs[i].root = createdView
		}
	} else {
		for i := new_size; i < origin_size; i++ {
			fb.vhs[i].root.baseView().Invisible()

		}
	}
	for i := 0; i < new_size; i++ {
		fb.bindData(&fb.vhs[i], i)
	}
	fb.gotCount = new_size
	return fb
}
