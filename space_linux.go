package faithtop

type FSpace FBox

func Space() *FSpace {
	return (*FSpace)(VBox().Expand())
}

func (f *FSpace) baseView() *FBaseView {
	return &f.FBaseView
}
