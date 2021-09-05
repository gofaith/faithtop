//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type TextBrowserImpl struct {
	*WidgetImpl
	textBrowser *widgets.QTextBrowser
}

func init() {
	newTextBrowserImpl = func() ITextBrowser {
		v := &TextBrowserImpl{
			textBrowser: widgets.NewQTextBrowser(nil),
		}
		v.WidgetImpl = widgetImplFrom(v.textBrowser.QWidget_PTR())
		return v
	}
}

func (l *TextBrowserImpl) Text(s string) ITextBrowser {
	l.textBrowser.SetText(s)
	return l
}

func (l *TextBrowserImpl) Interaction(flags TextInteractionFlag) ITextBrowser {
	l.textBrowser.SetTextInteractionFlags(core.Qt__TextInteractionFlag(flags))
	return l
}
