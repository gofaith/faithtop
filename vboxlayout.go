package faithtop

type IVBoxLayout interface {
	IBoxLayout
}

var newVBoxLayoutImpl func() IVBoxLayout

func VBoxLayout() IVBoxLayout {
	return newVBoxLayoutImpl()
}

func VBox(widgets ...IWidget) IWidget {
	return Widget().SizePolicy(SizePolicy__Expanding, SizePolicy__Fixed).Layout(VBoxLayout().AppendWidgets(widgets...))
}

func VBoxAlignLeft(widgets ...IWidget) IWidget {
	return Widget().Layout(VBoxLayout().Align(AlignLeading).AppendWidgets(widgets...))
}

func VBoxAlignHCenter(widgets ...IWidget) IWidget {
	return Widget().Layout(VBoxLayout().Align(AlignHCenter).AppendWidgets(widgets...))
}

func VBoxAlignRight(widgets ...IWidget) IWidget {
	return Widget().Layout(VBoxLayout().Align(AlignTrailing).AppendWidgets(widgets...))
}
