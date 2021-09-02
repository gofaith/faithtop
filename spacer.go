package faithtop

type ISpacer interface {
	ILayoutItem
}

var newHSpacerImpl, newVSpacerImpl func() ISpacer

func HSpacer() ISpacer {
	return newHSpacerImpl()
}

func VSpacer() ISpacer {
	return newVSpacerImpl()
}
