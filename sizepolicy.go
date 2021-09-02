package faithtop

type SizePolicyFlag int64

const (
	SizePolicy__GrowFlag   SizePolicyFlag = SizePolicyFlag(1)
	SizePolicy__ExpandFlag SizePolicyFlag = SizePolicyFlag(2)
	SizePolicy__ShrinkFlag SizePolicyFlag = SizePolicyFlag(4)
	SizePolicy__IgnoreFlag SizePolicyFlag = SizePolicyFlag(8)
)

type SizePolicy int64

const (
	SizePolicy__Fixed            SizePolicy = SizePolicy(0)
	SizePolicy__Minimum          SizePolicy = SizePolicy(SizePolicy__GrowFlag)
	SizePolicy__Maximum          SizePolicy = SizePolicy(SizePolicy__ShrinkFlag)
	SizePolicy__Preferred        SizePolicy = SizePolicy(SizePolicy__GrowFlag | SizePolicy__ShrinkFlag)
	SizePolicy__MinimumExpanding SizePolicy = SizePolicy(SizePolicy__GrowFlag | SizePolicy__ExpandFlag)
	SizePolicy__Expanding        SizePolicy = SizePolicy(SizePolicy__GrowFlag | SizePolicy__ShrinkFlag | SizePolicy__ExpandFlag)
	SizePolicy__Ignored          SizePolicy = SizePolicy(SizePolicy__ShrinkFlag | SizePolicy__GrowFlag | SizePolicy__IgnoreFlag)
)
