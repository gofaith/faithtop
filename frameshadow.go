package faithtop

type FrameShadow int64

const (
	FrameShadow_Plain  FrameShadow = FrameShadow(0x0010)
	FrameShadow_Raised FrameShadow = FrameShadow(0x0020)
	FrameShadow_Sunken FrameShadow = FrameShadow(0x0030)
)
