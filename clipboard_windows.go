package faithtop

import (
	"github.com/gofaith/walk"
)

type FClipboard struct {
	v *walk.ClipboardService
}

func Clipboard() *FClipboard {
	return &FClipboard{
		v: walk.Clipboard(),
	}
}

func (f *FClipboard) SetText(s string) *FClipboard {
	f.v.SetText(s)
	return f
}

func (f *FClipboard) GetText() string {
	s, _ := f.v.Text()
	return s
}

func (f *FClipboard) OnChange(fn func(string))*FClipboard  {
	f.v.ContentsChanged().Attach(func(){
		if fn!=nil{
			fn(f.GetText())
		}
	})
	return f
}