package faithtop

type FSpace FBox

func Space() *FSpace {
	return VBox().Expand()
}
