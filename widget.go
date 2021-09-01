package faithtop

type IWidget interface {
	Widget() IWidget
	VBox(children ...IWidget) IWidget
}

var newWidgetImpl func() IWidget

func Widget() IWidget {
	return newWidgetImpl()
}
