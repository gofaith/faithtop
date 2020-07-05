package faithtop

import (
	"github.com/StevenZack/livedata"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

type FEdit struct {
	FBaseView
	v        *gtk.Entry
	onChange func(text string)
}

func Edit() *FEdit {
	v := gtk.NewEntry()
	fb := &FEdit{}
	fb.v = v

	fb.widget = &v.Widget
	fb.v.Connect("changed", func(ctx *glib.CallbackContext) {
		if fb.onChange != nil {
			fb.onChange(fb.GetText())
		}
	})
	return fb
}

// ================================================================

func (v *FEdit) Size(w, h int) *FEdit {
	v.FBaseView.Size(w, h)
	return v
}

func (f *FEdit) Assign(v **FEdit) *FEdit {
	*v = f
	return f
}
func (vh *ViewHolder) GetEditByItemId(itemId string) *FEdit {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FEdit); ok {
			return lv
		}
	}
	return nil
}
func (v *FEdit) SetItemId(lv *FListView, itemId string) *FEdit {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetEditById(id string) *FEdit {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FEdit); ok {
			return b
		}
	}
	return nil
}
func (v *FEdit) baseView() *FBaseView {
	return &v.FBaseView
}
func (v *FEdit) OnClick(f func()) *FEdit {
	v.v.Connect("clicked", f)
	return v
}
func (v *FEdit) Expand() *FEdit {
	v.expand = true
	return v
}
func (v *FEdit) Disable() *FEdit {
	v.v.SetSensitive(false)
	return v
}
func (v *FEdit) Enable() *FEdit {
	v.v.SetSensitive(true)
	return v
}
func (v *FEdit) Visible() *FEdit {
	v.v.SetVisible(true)
	return v
}
func (v *FEdit) Invisible() *FEdit {
	v.FBaseView.Invisible()
	return v
}

func (v *FEdit) Tooltips(s string) *FEdit {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FEdit) Focus() *FEdit {
	v.FBaseView.Focus()
	return v
}

func (v *FEdit) OnDragDrop(f func([]string)) *FEdit {
	v.FBaseView.OnDragDrop(f)
	return v
}

//====================================================================
func (v *FEdit) Text(t string) *FEdit {
	v.v.SetText(t)
	return v
}
func (v *FEdit) GetText() string {
	return v.v.GetText()
}
func (v *FEdit) Editable(b bool) *FEdit {
	v.v.SetEditable(b)
	return v
}
func (v *FEdit) MaxLength(l int) *FEdit {
	v.v.SetMaxLength(l)
	return v
}
func (v *FEdit) InputTypePassword() *FEdit {
	v.v.SetVisibility(false)
	return v
}
func (v *FEdit) InputTypeNormal() *FEdit {
	v.v.SetVisibility(true)
	return v
}
func (v *FEdit) OnEnter(f func()) *FEdit {
	v.v.Connect("activate", f)
	return v
}

func (f *FEdit) OnChange(fn func(text string)) *FEdit {
	f.onChange = fn
	return f
}

func (f *FEdit) BindText(str *livedata.String) *FEdit {
	f.Text(str.Get())
	str.ObserveForever(func(s string) {
		if s == f.GetText() {
			return
		}
		RunOnUIThread(func() {
			f.Text(s)
		})
	})
	f.OnChange(func(text string) {
		if str.Get() == text {
			return
		}
		str.Post(text)
	})
	return f
}
