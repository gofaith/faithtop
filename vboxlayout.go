package faithtop

type IVBoxLayout interface {
	IBoxLayout
}

var newVBoxLayoutImpl func() IVBoxLayout

func VBoxLayout() IVBoxLayout {
	return newVBoxLayoutImpl()
}
