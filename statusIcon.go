package faithtop

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

type FStatusIcon struct {
	v        *gtk.StatusIcon
	isClosed bool
}

func StatusIcon() *FStatusIcon {
	s := &FStatusIcon{}
	s.v = gtk.NewStatusIconFromStock(gtk.STOCK_INFO)
	return s
}
func (s *FStatusIcon) Src(url string) *FStatusIcon {
	if StartsWidth(url, "/") {
		s.v.SetFromFile(url)
	} else if StartsWidth(url, "http") {
		go CacheNetFile(url, GetCacheDir(), func(fpath string) {
			RunOnUIThread(func() {
				s.v.SetFromFile(fpath)
			})
		})
	} else if StartsWidth(url, "file://") {
		s.v.SetFromFile(url[len("file://"):])
	} else if StartsWidth(url, "iconName://") {
		s.v.SetFromIconName(url[len("iconName://"):])
	}
	return s
}
func (s *FStatusIcon) SrcIconName(url string) *FStatusIcon {
	s.v.SetFromIconName(url)
	return s
}
func (s *FStatusIcon) SrcFile(url string) *FStatusIcon {
	s.v.SetFromFile(url)
	return s
}
func (s *FStatusIcon) SrcNet(url string) *FStatusIcon {
	go CacheNetFile(url, GetCacheDir(), func(fpath string) {
		RunOnUIThread(func() {
			s.v.SetFromFile(fpath)
		})
	})
	return s
}
func (s *FStatusIcon) Menus(ms ...IMenuItem) *FStatusIcon {
	nm := gtk.NewMenu()
	for _, m := range ms {
		nm.Append(m.getMenuItem())
	}
	nm.ShowAll()
	s.v.Connect("popup-menu", func(cbx *glib.CallbackContext) {
		nm.Popup(nil, nil, gtk.StatusIconPositionMenu, s.v, uint(cbx.Args(0)), uint32(cbx.Args(1)))
	})
	return s
}
func (s *FStatusIcon) Tooltips(t string) *FStatusIcon {
	s.v.SetTooltipMarkup(t)
	return s
}
func (s *FStatusIcon) Title(t string) *FStatusIcon {
	s.v.SetTitle(t)
	return s
}
func (s *FStatusIcon) Close() *FStatusIcon {
	s.v.SetVisible(false)
	return s
}
func (s *FStatusIcon) Show() *FStatusIcon {
	s.v.SetVisible(true)
	return s
}
