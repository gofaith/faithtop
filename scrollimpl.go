//go:build impl

package faithtop

import "github.com/therecipe/qt/widgets"

type ScrollImpl struct {
	*WidgetImpl
	scroll *widgets.QScrollArea
}

func init() {
	newScrollImpl = func(child IWidget) IScroll {
		v := &ScrollImpl{
			scroll: widgets.NewQScrollArea(nil),
		}
		v.WidgetImpl = widgetImplFrom(v.scroll.QWidget_PTR())
		v.scroll.SetWidget(child.getWidget().(*WidgetImpl).widget)
		return v
	}
}
