package faithtop

import (
	"github.com/StevenZack/livedata"
	"github.com/mattn/go-gtk/gtk"
)

type FButton struct {
	FBaseView
	v       *gtk.Button
	onClick func()
}

func Button() *FButton {
	v := gtk.NewButton()
	fb := &FButton{}
	fb.v = v
	fb.widget = &v.Widget
	fb.v.Connect("clicked", func() {
		if fb.onClick != nil {
			fb.onClick()
		}
	})

	return fb
}

// ================================================================
func (v *FButton) Size(w, h int) *FButton {
	v.FBaseView.Size(w, h)
	return v
}
func (f *FButton) Assign(v **FButton) *FButton {
	*v = f
	return f
}
func (vh *ViewHolder) GetButton(itemId string) *FButton {
	if v, ok := vh.vlist[itemId]; ok {
		if lv, ok := v.(*FButton); ok {
			return lv
		}
	}
	return nil
}
func (v *FButton) SetItemId(lv *FListView, itemId string) *FButton {
	lv.vhs[lv.currentCreation].vlist[itemId] = v
	return v
}

func GetButtonById(id string) *FButton {
	if v, ok := idMap[id]; ok {
		if b, ok := v.(*FButton); ok {
			return b
		}
	}
	return nil
}
func (v *FButton) baseView() *FBaseView {
	return &v.FBaseView
}
func (v *FButton) OnClick(f func()) *FButton {
	v.onClick = f
	return v
}
func (v *FButton) SetId(id string) *FButton {
	idMap[id] = v
	return v
}
func (v *FButton) Expand() *FButton {
	v.expand = true
	return v
}
func (v *FButton) Disable() *FButton {
	v.v.SetSensitive(false)
	return v
}
func (v *FButton) Enable() *FButton {
	v.v.SetSensitive(true)
	return v
}
func (v *FButton) Visible() *FButton {
	v.v.SetVisible(true)
	return v
}
func (v *FButton) Invisible() *FButton {
	v.FBaseView.Invisible()
	return v
}

func (v *FButton) Tooltips(s string) *FButton {
	v.v.SetTooltipMarkup(s)
	return v
}
func (v *FButton) Focus() *FButton {
	v.FBaseView.Focus()
	return v
}

func (v *FButton) OnDragDrop(f func([]string)) *FButton {
	v.FBaseView.OnDragDrop(f)
	return v
}

//====================================================================

func (v *FButton) Text(t string) *FButton {
	v.v.SetLabel(t)
	return v
}
func (v *FButton) GetText() string {
	return v.v.GetLabel()
}
func (v *FButton) Image(img *FImage) *FButton {
	v.v.SetImage(img.v)
	return v
}

func (f *FButton) BindEnabled(l *livedata.Bool) *FButton {
	l.ObserveForever(func(b bool) {
		RunOnUIThread(func() {
			if b {
				f.Enable()
			} else {
				f.Disable()
			}
		})
	})
	return f
}

func (f *FButton) BindVisible(l *livedata.Bool) *FButton {
	l.ObserveForever(func(b bool) {
		RunOnUIThread(func() {
			if b {
				f.Visible()
			} else {
				f.Invisible()
			}
		})
	})
	return f
}

func (f *FButton) BindInvisible(l *livedata.Bool) *FButton {
	l.ObserveForever(func(b bool) {
		RunOnUIThread(func() {
			if b {
				f.Invisible()
			} else {
				f.Visible()
			}
		})
	})
	return f
}

func (f *FButton) BindText(l *livedata.String)*FButton  {
	l.ObserveForever(func(s string){
		if s==f.GetText(){
			return
		}
		RunOnUIThread(func(){
			f.Text(s)
		})
	})
	return f
}