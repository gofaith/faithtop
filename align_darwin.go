package faithtop

type FAlign FBox

func AlignLeft(i IView) *FAlign {
	return (*FAlign)(HBox().Append(i, Space()))
}

func AlignRight(i IView) *FAlign {
	return (*FAlign)(HBox().Append(Space(), i))
}

func AlignHCenter(i IView) *FAlign {
	return (*FAlign)(HBox().Append(Space(), i, Space()))
}

func AlignTop(i IView) *FAlign {
	return (*FAlign)(VBox().Append(i, Space()))
}

func AlignBottom(i IView) *FAlign {
	return (*FAlign)(VBox().Append(Space(), i))
}

func AlignVCenter(i IView) *FAlign {
	return (*FAlign)(VBox().Append(Space(), i, Space()))
}

func AlignCenter(i IView) *FAlign {
	return AlignVCenter(AlignHCenter(i))
}
