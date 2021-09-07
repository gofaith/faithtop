package faithtop

type IFrame interface {
	IWidget
	Shape(shape FrameShape) IFrame
	Shadow(shadow FrameShadow) IFrame
	LineWidth(i int) IFrame
	MidLineWidth(i int) IFrame
}

var newFrameImpl func() IFrame

func Frame() IFrame {
	return newFrameImpl()
}
