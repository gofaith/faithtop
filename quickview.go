package faithtop

type IQuickView interface {
	IWidget
	Assign(v *IQuickView) IQuickView
	Source(qmlFile string) IQuickView
}

var newQuickViewImpl func(win IWindow) IQuickView

func QuickView(win IWindow) IQuickView {
	return newQuickViewImpl(win)
}

func QuickView2(win IWindow, source string) IQuickView {
	return newQuickViewImpl(win).Source(source)
}
