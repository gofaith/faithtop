package faithtop

type IWidget interface {
	getWidget() IWidget
	Size(width, height SizePolicy) IWidget
	Layout(layout ILayout) IWidget
	Align(align AlignmentFlag) IWidget
}

var newWidgetImpl func() IWidget

func Widget() IWidget {
	return newWidgetImpl()
}
