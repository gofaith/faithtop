package faithtop

type FrameShape int64

const (
	FrameShape_NoFrame     FrameShape = FrameShape(0)
	FrameShape_Box         FrameShape = FrameShape(0x0001)
	FrameShape_Panel       FrameShape = FrameShape(0x0002)
	FrameShape_WinPanel    FrameShape = FrameShape(0x0003)
	FrameShape_HLine       FrameShape = FrameShape(0x0004)
	FrameShape_VLine       FrameShape = FrameShape(0x0005)
	FrameShape_StyledPanel FrameShape = FrameShape(0x0006)
)

