package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FTabLayout struct {
	FBaseView
	v *gtk.Notebook
}

func TabLayout() *FTabLayout {
	v := gtk.NewNotebook()
	fb := &FTabLayout{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	setupWidget(fb)
	return fb
}

// ================================================================
func (f *FTabLayout) Assign(v **FTabLayout) *FTabLayout {
	*v = f
	return f
}
func (v *FTabLayout) OnDragDrop(f func([]string)) *FTabLayout {
	v.FBaseView.OnDragDrop(f)
	return v
}
func (v *FTabLayout) Size(w, h int) *FTabLayout {
	v.FBaseView.Size(w, h)
	return v
}
func (vh *ViewHolder) GetTabsByItemId(itemId string) *FTabLayout {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FTabLayout); ok {
			return lv
		}
	}
	return nil
}
func (v *FTabLayout) SetItemId(lv *FListView, itemId string) *FTabLayout {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetTabsById(id string) *FTabLayout {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FTabLayout); ok {
			return b
		}
	}
	return nil
}
func (v *FTabLayout) getBaseView() *FBaseView {
	return &v.FBaseView
}
func (v *FTabLayout) OnClick(f func()) *FTabLayout {
	v.v.Connect("clicked", f)
	return v
}
func (v *FTabLayout) SetId(id string) *FTabLayout {
	idMap[id] = v
	return v
}
func (v *FTabLayout) Expand() *FTabLayout {
	v.expand = true
	return v
}
func (v *FTabLayout) NotFill() *FTabLayout {
	v.notFill = true
	return v
}
func (v *FTabLayout) Disable() *FTabLayout {
	v.v.SetSensitive(false)
	return v
}
func (v *FTabLayout) Enable() *FTabLayout {
	v.v.SetSensitive(true)
	return v
}
func (v *FTabLayout) Visible() *FTabLayout {
	v.v.SetVisible(true)
	return v
}
func (v *FTabLayout) Invisible() *FTabLayout {
	v.FBaseView.Invisible()
	return v
}

func (v *FTabLayout) Tooltips(s string) *FTabLayout {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FTabLayout) Focus() *FTabLayout {
	v.FBaseView.Focus()
	return v
}
func (v *FTabLayout) Padding(i uint) *FTabLayout {
	v.padding = i
	return v
}

//====================================================================
type FTab struct {
	title   *gtk.Label
	content IView
}

func Tab(title string, view IView) *FTab {
	fb := &FTab{}
	fb.title = gtk.NewLabel(title)
	fb.content = view
	return fb
}

func (v *FTabLayout) Tabs(ps ...*FTab) *FTabLayout {
	for _, p := range ps {
		v.v.AppendPage(p.content.getBaseView().view, p.title)
	}
	return v
}
func (v *FTabLayout) OnSwitchPage(f func()) *FTabLayout {
	v.v.Connect("switch-page", f)
	return v
}
