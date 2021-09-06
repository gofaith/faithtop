package faithtop

type IHBoxLayout interface {
	IBoxLayout
}

var newHBoxLayoutImpl func() IHBoxLayout

func HBoxLayout() IHBoxLayout {
	return newHBoxLayoutImpl()
}

func HBox(widgets ...IWidget) IWidget {
	return Widget().Layout(HBoxLayout().AppendWidgets(widgets...))
}
