package faithtop

type IHBoxLayout interface {
	IBoxLayout	
}

var newHBoxLayoutImpl func()IHBoxLayout

func HBoxLayout()IHBoxLayout{
	return newHBoxLayoutImpl()
}