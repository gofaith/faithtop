//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type WidgetImpl struct {
	widget *widgets.QWidget
}

func init() {
	newWidgetImpl = func() IWidget {
		return &WidgetImpl{
			widget: widgets.NewQWidget(nil, 0),
		}
	}
}

func (w *WidgetImpl) Widget() IWidget {
	return w
}

func widgetImplFrom(parent *widgets.QWidget) *WidgetImpl {
	return &WidgetImpl{
		widget: parent,
	}
}

func (w *WidgetImpl) VBox(children ...IWidget) IWidget {
	layout := widgets.NewQVBoxLayout()
	for _, child := range children {
		widget := child.Widget().(*WidgetImpl)
		layout.AddWidget(widget.widget, 1, core.Qt__AlignCenter)
	}
	w.widget.SetLayout(layout)
	return w
}
