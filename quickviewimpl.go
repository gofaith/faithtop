//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
)

type (
	QuickViewImpl struct {
		*WidgetImpl
		quickView *quick.QQuickView
		bridge    *Bridge
	}

	Bridge struct {
		quick.QQuickItem
		_        func()                               `constructor:"init"`
		_        func(funcName string, args []string) `signal:"callGo,auto"`
		_        func(funcName string, args []string) `signal:"callQml"`
		handlers map[string]func(args []string)
	}
)

func init() {
	newQuickViewImpl = func(win IWindow) IQuickView {
		if win == nil {
			panic("the parent window cannot be nil")
		}
		v := &QuickViewImpl{
			quickView: quick.NewQQuickView(nil),
			bridge:    NewBridge(nil),
		}
		var widget *widgets.QWidget
		widget = widget.CreateWindowContainer(v.quickView, win.(*WindowImpl).window, 0)
		v.WidgetImpl = widgetImplFrom(widget)

		v.quickView.RootContext().SetContextProperty("bridge", v.bridge)

		return v
	}
	Bridge_QmlRegisterType2("Bridge", 1, 0, "Bridge")
}

func (b *Bridge) init() {
	b.handlers = make(map[string]func(args []string))
}
func (b *Bridge) callGo(funcName string, args []string) {
	if b.handlers == nil {
		return
	}
	if handler, ok := b.handlers[funcName]; ok {
		handler(args)
	} else {
		panic("funcName [" + funcName + "] is not registered on QuickView")
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

func (q *QuickViewImpl) CallQml(funcName string, args ...string) {
	q.bridge.CallQml(funcName, args)
}

func (q *QuickViewImpl) RegisterFunction(funcName string, fn func(args []string)) IQuickView {
	q.bridge.handlers[funcName] = fn
	return q
}
