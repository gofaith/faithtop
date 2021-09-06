package faithtop

type (
	IButtonGroup interface {
		OnClicked(fn func(text string)) IButtonGroup
		BindValue(text *string) IButtonGroup
	}
)

var newButtonGroupImpl func() IButtonGroup

func ButtonGroup() IButtonGroup {
	return newButtonGroupImpl()
}
