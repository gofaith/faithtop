package faithtop

type FSpace FBox

func Space() *FSpace {
	return (*FSpace)(VBox().Expand())
}
