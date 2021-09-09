package faithtop

type IScroll interface {
	IWidget
}

var newScrollImpl func(child IWidget) IScroll

func Scroll(child IWidget) IScroll {
	return newScrollImpl(child)
}
