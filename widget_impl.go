//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/widgets"
)

type WidgetImpl struct {
	widget        *widgets.QWidget
	alignmentFlag AlignmentFlag
}

func init() {
	newWidgetImpl = func() IWidget {
		return &WidgetImpl{
			widget: widgets.NewQWidget(nil, 0),
		}
	}
}

func (w *WidgetImpl) getWidget() IWidget {
	return w
}

func widgetImplFrom(parent *widgets.QWidget) *WidgetImpl {
	return &WidgetImpl{
		widget: parent,
	}
}

func (w *WidgetImpl) Layout(layout ILayout) IWidget {
	w.widget.SetLayout(layout.getLayout().(*LayoutImpl).layout)
	return w
}

func (w *WidgetImpl) Align(align AlignmentFlag) IWidget {
	w.alignmentFlag = align
	return w
}
