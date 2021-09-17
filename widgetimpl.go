//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"strings"
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

func (w *WidgetImpl) SizePolicy(width, height SizePolicy) IWidget {
	w.widget.SetSizePolicy2(widgets.QSizePolicy__Policy(width), widgets.QSizePolicy__Policy(height))
	return w
}

func (w *WidgetImpl) Style(styleSheet string) IWidget {
	w.widget.SetStyleSheet(styleSheet)
	return w
}

func (w *WidgetImpl) MinWidth(width int) IWidget {
	w.widget.SetMinimumWidth(width)
	return w
}

func (w *WidgetImpl) MinHeight(height int) IWidget {
	w.widget.SetMinimumHeight(height)
	return w
}

func (w *WidgetImpl) MaxWidth(width int) IWidget {
	w.widget.SetMaximumWidth(width)
	return w
}

func (w *WidgetImpl) MaxHeight(height int) IWidget {
	w.widget.SetMaximumHeight(height)
	return w
}

func (w *WidgetImpl) Enabled(b bool) IWidget {
	w.widget.SetEnabled(b)
	return w
}

func (w *WidgetImpl) OnDragEnter(fn func(widget IWidget)) IWidget {
	w.widget.SetAcceptDrops(true)
	w.widget.ConnectDragEnterEvent(func(event *gui.QDragEnterEvent) {
		fn(w)
		event.AcceptProposedAction()
	})
	return w
}

func (w *WidgetImpl) OnDragLeave(fn func(widget IWidget)) IWidget {
	w.widget.ConnectDragLeaveEvent(func(event *gui.QDragLeaveEvent) {
		fn(w)
	})
	return w
}

func (w *WidgetImpl) OnDrop(fn func(widget IWidget, urls []string)) IWidget {
	w.widget.SetAcceptDrops(true)
	w.widget.ConnectDropEvent(func(event *gui.QDropEvent) {
		urls := []string{}
		for _, format := range event.MimeData().Formats() {
			d := event.MimeData().Data(format).Data()
			for _, line := range strings.Split(d, "\r\n") {
				if line != "" {
					urls = append(urls, line)
				}
			}
		}
		fn(w, urls)
		event.AcceptProposedAction()
	})
	return w
}

func (w *WidgetImpl) AssignWidget(v *IWidget) IWidget {
	*v = w
	return w
}

func (w *WidgetImpl) Cursor(shape CursorShape) IWidget {
	w.widget.SetCursor(gui.NewQCursor2(core.Qt__CursorShape(shape)))
	return w
}

func (w *WidgetImpl) OnMousePress(fn func(widget IWidget)) IWidget {
	w.widget.ConnectMousePressEvent(func(event *gui.QMouseEvent) {
		fn(w)
		event.Accept()
	})
	return w
}

func (w *WidgetImpl) OnMouseRelease(fn func(widget IWidget)) IWidget {
	w.widget.ConnectMouseReleaseEvent(func(event *gui.QMouseEvent) {
		fn(w)
		event.Accept()
	})
	return w
}

func (w *WidgetImpl) OnKeyPress(fn func(widget IWidget, key Key)) IWidget {
	w.widget.ConnectKeyPressEvent(func(event *gui.QKeyEvent) {
		fn(w, Key(event.Key()))
		event.Accept()
	})
	return w
}

func (w *WidgetImpl) ToolTip(s string) IWidget {
	w.widget.SetToolTip(s)
	return w
}

func (w *WidgetImpl) Visible(b bool) IWidget {
	w.widget.SetVisible(b)
	return w
}
