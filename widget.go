package faithtop

type IWidget interface {
	getWidget() IWidget
	AssignWidget(v *IWidget) IWidget
	SizePolicy(width, height SizePolicy) IWidget
	MinWidth(width int) IWidget
	MinHeight(height int) IWidget
	MaxWidth(width int) IWidget
	MaxHeight(height int) IWidget
	Layout(layout ILayout) IWidget
	Align(align AlignmentFlag) IWidget
	/** Style set a CSS stylesheet string to a widget.
	e.g.
	Button().Style(`QPushButton{background-color:blue;color:white;} QPushButton:hover{background-color:red;} QPushButton:pressed{background-color:white;color:grey}`)
	*/
	Style(styleSheet string) IWidget
	OnDragEnter(fn func()) IWidget
	OnDragLeave(fn func()) IWidget
	OnDrop(fn func(urls []string)) IWidget
	Enabled(b bool) IWidget
}

var newWidgetImpl func() IWidget

func Widget() IWidget {
	return newWidgetImpl()
}
