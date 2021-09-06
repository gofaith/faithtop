package faithtop

type CheckState int64

const (
	CheckState_Unchecked        CheckState = CheckState(0)
	CheckState_PartiallyChecked CheckState = CheckState(1)
	CheckState_Checked          CheckState = CheckState(2)
)
