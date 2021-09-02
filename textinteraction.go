package faithtop

type TextInteractionFlag int64

const (
	NoTextInteraction         TextInteractionFlag = TextInteractionFlag(0)
	TextSelectableByMouse     TextInteractionFlag = TextInteractionFlag(1)
	TextSelectableByKeyboard  TextInteractionFlag = TextInteractionFlag(2)
	LinksAccessibleByMouse    TextInteractionFlag = TextInteractionFlag(4)
	LinksAccessibleByKeyboard TextInteractionFlag = TextInteractionFlag(8)
	TextEditable              TextInteractionFlag = TextInteractionFlag(16)
	TextEditorInteraction     TextInteractionFlag = TextInteractionFlag(TextSelectableByMouse | TextSelectableByKeyboard | TextEditable)
	TextBrowserInteraction    TextInteractionFlag = TextInteractionFlag(TextSelectableByMouse | LinksAccessibleByMouse | LinksAccessibleByKeyboard)
)
