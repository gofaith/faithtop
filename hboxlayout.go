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

func HBoxAlignTop(widgets ...IWidget) IWidget {
	return Widget().Layout(HBoxLayout().Align(AlignLeading).AppendWidgets(widgets...))
}

func HBoxAlignVCenter(widgets ...IWidget) IWidget {
	return Widget().Layout(HBoxLayout().Align(AlignVCenter).AppendWidgets(widgets...))
}

func HBoxAlignBottom(widgets ...IWidget) IWidget {
	return Widget().Layout(HBoxLayout().Align(AlignTrailing).AppendWidgets(widgets...))
}
