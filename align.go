package faithtop

type AlignmentFlag int64

const (
	AlignLeft            AlignmentFlag = AlignmentFlag(0x0001)
	AlignLeading         AlignmentFlag = AlignmentFlag(AlignLeft)
	AlignRight           AlignmentFlag = AlignmentFlag(0x0002)
	AlignTrailing        AlignmentFlag = AlignmentFlag(AlignRight)
	AlignHCenter         AlignmentFlag = AlignmentFlag(0x0004)
	AlignJustify         AlignmentFlag = AlignmentFlag(0x0008)
	AlignAbsolute        AlignmentFlag = AlignmentFlag(0x0010)
	AlignHorizontal_Mask AlignmentFlag = AlignmentFlag(AlignLeft | AlignRight | AlignHCenter | AlignJustify | AlignAbsolute)
	AlignTop             AlignmentFlag = AlignmentFlag(0x0020)
	AlignBottom          AlignmentFlag = AlignmentFlag(0x0040)
	AlignVCenter         AlignmentFlag = AlignmentFlag(0x0080)
	AlignBaseline        AlignmentFlag = AlignmentFlag(0x0100)
	AlignVertical_Mask   AlignmentFlag = AlignmentFlag(AlignTop | AlignBottom | AlignVCenter | AlignBaseline)
	AlignCenter          AlignmentFlag = AlignmentFlag(AlignVCenter | AlignHCenter)
)
