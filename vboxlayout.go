package faithtop

type IVBoxLayout interface {
	IBoxLayout
}

var newVBoxLayoutImpl func() IVBoxLayout

func VBoxLayout() IVBoxLayout {
	return newVBoxLayoutImpl()
}

func VBox(widgets ...IWidget) IWidget {
	return Widget().Layout(VBoxLayout().AppendWidgets(widgets...))
}
