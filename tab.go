package faithtop

import (
	"github.com/mattn/go-gtk/gtk"
)

type FTabs struct {
	FBaseView
	v *gtk.Notebook
}

func Tabs() *FTabs {
	v := gtk.NewNotebook()
	fb := &FTabs{}
	fb.v = v
	fb.view = v
	fb.widget = &v.Widget
	setupWidget(fb)
	return fb
}

// ================================================================
func (v *FTabs) Size(w, h int) *FTabs {
	v.FBaseView.Size(w, h)
	return v
}
func (vh *ViewHolder) GetTabsByItemId(itemId string) *FTabs {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FTabs); ok {
			return lv
		}
	}
	return nil
}
func (v *FTabs) SetItemId(lv *FListView, itemId string) *FTabs {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetTabsById(id string) *FTabs {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FTabs); ok {
			return b
		}
	}
	return nil
}
func (v *FTabs) getBaseView() *FBaseView {
	return &v.FBaseView
}
func (v *FTabs) OnClick(f func()) *FTabs {
	v.v.Connect("clicked", f)
	return v
}
func (v *FTabs) SetId(id string) *FTabs {
	idMap[id] = v
	return v
}
func (v *FTabs) Expand() *FTabs {
	v.expand = true
	return v
}
func (v *FTabs) NotFill() *FTabs {
	v.notFill = true
	return v
}
func (v *FTabs) Disable() *FTabs {
	v.v.SetSensitive(false)
	return v
}
func (v *FTabs) Enable() *FTabs {
	v.v.SetSensitive(true)
	return v
}
func (v *FTabs) Visible() *FTabs {
	v.v.SetVisible(true)
	return v
}
func (v *FTabs) Invisible() *FTabs {
	v.v.SetVisible(false)
	return v
}

func (v *FTabs) Tooltips(s string) *FTabs {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FTabs) Focus() *FTabs {
	v.FBaseView.Focus()
	return v
}
func (v *FTabs) Padding(i uint) *FTabs {
	v.padding = i
	return v
}

//====================================================================
type FTabPage struct {
	title   *gtk.Label
	content gtk.IWidget
}

func Page(title string, view IView) *FTabPage {
	fb := &FTabPage{}
	fb.title = gtk.NewLabel(title)
	fb.content = view.getBaseView().view
	return fb
}

func (v *FTabs) Pages(ps ...*FTabPage) *FTabs {
	for _, p := range ps {
		v.v.AppendPage(p.content, p.title)
	}
	return v
}
