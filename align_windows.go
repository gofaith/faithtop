package faithtop

//todo
func AlignLeft(i IView) *FBox {
	return HBox().Append(
		i,
		HSpace(),
	)
}

func AlignRight(i IView) *FBox {
	return HBox().Append(
		HSpace(),
		i,
	)
}

func AlignHCenter(i IView) *FBox {
	return HBox().Append(
		HSpace(),
		i,
		HSpace(),
	)
}

func AlignTop(i IView) *FBox {
	return VBox().Append(
		i,
		VSpace(),
	)
}

func AlignBottom(i IView) *FBox {
	return VBox().Append(
		VSpace(),
		i,
	)
}

func AlignVCenter(i IView) *FBox {
	return VBox().Append(
		VSpace(),
		i,
		VSpace(),
	)
}

func AlignCenter(i IView) *FBox {
	return AlignHCenter(AlignVCenter(i))
}
