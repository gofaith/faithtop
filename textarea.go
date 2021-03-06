package faithtop

import (
	"github.com/StevenZack/livedata"
	"github.com/mattn/go-gtk/gtk"
)

type FTextArea struct {
	FBaseView
	v *gtk.TextView
}

func TextArea() *FTextArea {
	fb := &FTextArea{}
	v := gtk.NewTextView()
	fb.v = v

	fb.widget = &v.Widget
	fb.v.SetEditable(true)

	return fb
}

// ================================================================
func (v *FTextArea) Size(w, h int) *FTextArea {
	v.FBaseView.Size(w, h)
	return v
}
func (f *FTextArea) Assign(v **FTextArea) *FTextArea {
	*v = f
	return f
}
func (vh *ViewHolder) GetTextArea(itemId string) *FTextArea {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FTextArea); ok {
			return lv
		}
	}
	return nil
}
func (v *FTextArea) SetItemId(lv *FListView, itemId string) *FTextArea {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetTextAreaById(id string) *FTextArea {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FTextArea); ok {
			return b
		}
	}
	return nil
}
func (v *FTextArea) baseView() *FBaseView {
	return &v.FBaseView
}
func (v *FTextArea) SetId(id string) *FTextArea {
	idMap[id] = v
	return v
}
func (v *FTextArea) Expand() *FTextArea {
	v.expand = true
	return v
}
func (v *FTextArea) Disable() *FTextArea {
	v.v.SetSensitive(false)
	return v
}
func (v *FTextArea) Enable() *FTextArea {
	v.v.SetSensitive(true)
	return v
}
func (v *FTextArea) Visible() *FTextArea {
	v.v.SetVisible(true)
	return v
}
func (v *FTextArea) Invisible() *FTextArea {
	v.FBaseView.Invisible()
	return v
}

func (v *FTextArea) Tooltips(s string) *FTextArea {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FTextArea) Focus() *FTextArea {
	v.FBaseView.Focus()
	return v
}
func (v *FTextArea) OnDragDrop(f func([]string)) *FTextArea {
	v.FBaseView.OnDragDrop(f)
	return v
}

//====================================================================
func (f *FTextArea) Text(t string) *FTextArea {
	f.v.GetBuffer().SetText(t)
	return f
}
func (f *FTextArea) GetText() string {
	var start, end gtk.TextIter
	f.v.GetBuffer().GetStartIter(&start)
	f.v.GetBuffer().GetEndIter(&end)
	return f.v.GetBuffer().GetText(&start, &end, false)
}
func (f *FTextArea) Editable(b bool) *FTextArea {
	f.v.SetEditable(b)
	return f
}

func (f *FTextArea) OnChange(fn func(str string)) *FTextArea {
	f.v.Connect("event-after", func() {
		fn(f.GetText())
	})
	return f
}

func (f *FTextArea) BindText(l *livedata.String) *FTextArea {
	l.ObserveForever(func(str string) {
		if str == f.GetText() {
			return
		}
		RunOnUIThread(func() {
			f.Text(str)
		})
	})
	f.v.Connect("event-after", func() {
		text := f.GetText()
		if text != l.Get() {
			l.Post(text)
		}
	})
	return f
}
