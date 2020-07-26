package faithtop

func HSpace() *FBox {
	return VBox().Expand()
}

func VSpace() *FBox {
	return HBox().Expand()
}
