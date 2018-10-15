package faithtop

type FListView struct {
	FScroll
	vhs             []ViewHolder
	currentCreation int
	createView      func(*FListView) IView
	bindData        func(*ViewHolder, int)
	getCount        func() int
}

type ViewHolder struct {
	vlist map[string]IView
}

func VlistView(createView func(*FListView) IView, bindData func(*ViewHolder, int), getCount func() int) *FListView {
	fb := &FListView{}
	fb.FScroll = *VScroll()
	fb.createView = createView
	fb.bindData = bindData
	fb.getCount = getCount
	for i := 0; i < getCount(); i++ {
		fb.currentCreation = i
		fb.vhs = append(fb.vhs, ViewHolder{vlist: make(map[string]IView)})
		fb.Append(createView(fb))
	}
	fb.execBindData()
	return fb
}
func HlistView(createView func(*FListView) IView, bindData func(*ViewHolder, int), getCount func() int) *FListView {
	fb := &FListView{}
	fb.FScroll = *HScroll()
	fb.createView = createView
	fb.bindData = bindData
	fb.getCount = getCount
	for i := 0; i < getCount(); i++ {
		fb.currentCreation = i
		fb.vhs = append(fb.vhs, ViewHolder{vlist: make(map[string]IView)})
		fb.Append(createView(fb))
	}
	fb.execBindData()
	return fb
}

// ================================================================
func (v *FListView) Size(w, h int) *FListView {
	v.FBaseView.Size(w, h)
	return v
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
func (v *FListView) getBaseView() *FBaseView {
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
func (v *FListView) NotFill() *FListView {
	v.notFill = true
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
	v.v.SetVisible(false)
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
func (v *FListView) Padding(i uint) *FListView {
	v.padding = i
	return v
}

//====================================================================
func (v *FListView) execBindData() *FListView {
	for i := 0; i < v.getCount(); i++ {
		v.bindData(&v.vhs[i], i)
	}
	return v
}
