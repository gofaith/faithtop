//go:build impl

package faithtop

import "github.com/therecipe/qt/widgets"

type TabWidgetImpl struct {
	*WidgetImpl
	tabWidget *widgets.QTabWidget
}

func init() {
	newTabWidgetImpl = func(pages ...TabPage) ITabWidget {
		v := &TabWidgetImpl{
			tabWidget: widgets.NewQTabWidget(nil),
		}
		v.WidgetImpl = widgetImplFrom(v.tabWidget.QWidget_PTR())

		for _, page := range pages {
			v.tabWidget.AddTab(page.Content.getWidget().(*WidgetImpl).widget, page.Name)
		}
		return v
	}
}

func (t *TabWidgetImpl) Assign(v *ITabWidget) ITabWidget {
	*v = t
	return t
}

func (t *TabWidgetImpl) CurrentIndex(i int) ITabWidget {
	t.tabWidget.SetCurrentIndex(i)
	return t
}

func (t *TabWidgetImpl) GetCurrentIndex() int {
	return t.tabWidget.CurrentIndex()
}

func (t *TabWidgetImpl) OnIndexChanged(fn func(i int)) ITabWidget {
	t.tabWidget.ConnectCurrentChanged(fn)
	return t
}
