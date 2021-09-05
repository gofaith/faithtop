package faithtop

type (
	ITabWidget interface {
		IWidget
		Assign(v *ITabWidget) ITabWidget
		CurrentIndex(i int) ITabWidget
		GetCurrentIndex() int
		OnIndexChanged(fn func(i int)) ITabWidget
	}
	TabPage struct {
		Name    string
		Content IWidget
	}
)

var newTabWidgetImpl func(pages ...TabPage) ITabWidget

func TabWidget(pages ...TabPage) ITabWidget {
	return newTabWidgetImpl(pages...)
}
