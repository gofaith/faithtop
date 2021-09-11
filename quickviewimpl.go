//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

type QuickViewImpl struct {
	*WidgetImpl
	quickView *quick.QQuickView
}

func init() {
	newQuickViewImpl = func(win IWindow) IQuickView {
		v := &QuickViewImpl{
			quickView: quick.NewQQuickView(nil),
		}
		var widget *widgets.QWidget
		widget = widget.CreateWindowContainer(v.quickView, win.(*WindowImpl).window, 0)
		v.WidgetImpl = widgetImplFrom(widget)
		return v
	}
}

func (q *QuickViewImpl) Assign(v *IQuickView) IQuickView {
	*v = q
	return q
}

func (q *QuickViewImpl) Source(qmlFile string) IQuickView {
	q.quickView.SetSource(core.NewQUrl3(qmlFile, 0))
	return q
}
