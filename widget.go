package faithtop

type IWidget interface {
	getWidget() IWidget
	Layout(layout ILayout) IWidget
	Align(align AlignmentFlag)IWidget
}

var newWidgetImpl func() IWidget

func Widget() IWidget {
	return newWidgetImpl()
}
