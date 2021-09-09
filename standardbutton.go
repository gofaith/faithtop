package faithtop

type StandardButton int64

var (
	StandardButton_NoButton        StandardButton = StandardButton(0x00000000)
	StandardButton_Ok              StandardButton = StandardButton(0x00000400)
	StandardButton_Save            StandardButton = StandardButton(0x00000800)
	StandardButton_SaveAll         StandardButton = StandardButton(0x00001000)
	StandardButton_Open            StandardButton = StandardButton(0x00002000)
	StandardButton_Yes             StandardButton = StandardButton(0x00004000)
	StandardButton_YesToAll        StandardButton = StandardButton(0x00008000)
	StandardButton_No              StandardButton = StandardButton(0x00010000)
	StandardButton_NoToAll         StandardButton = StandardButton(0x00020000)
	StandardButton_Abort           StandardButton = StandardButton(0x00040000)
	StandardButton_Retry           StandardButton = StandardButton(0x00080000)
	StandardButton_Ignore          StandardButton = StandardButton(0x00100000)
	StandardButton_Close           StandardButton = StandardButton(0x00200000)
	StandardButton_Cancel          StandardButton = StandardButton(0x00400000)
	StandardButton_Discard         StandardButton = StandardButton(0x00800000)
	StandardButton_Help            StandardButton = StandardButton(0x01000000)
	StandardButton_Apply           StandardButton = StandardButton(0x02000000)
	StandardButton_Reset           StandardButton = StandardButton(0x04000000)
	StandardButton_RestoreDefaults StandardButton = StandardButton(0x08000000)
	StandardButton_FirstButton     StandardButton = StandardButton(StandardButton_Ok)
	StandardButton_LastButton      StandardButton = StandardButton(StandardButton_RestoreDefaults)
	StandardButton_YesAll          StandardButton = StandardButton(StandardButton_YesToAll)
	StandardButton_NoAll           StandardButton = StandardButton(StandardButton_NoToAll)
	StandardButton_Default         StandardButton = StandardButton(0x00000100)
	StandardButton_Escape          StandardButton = StandardButton(0x00000200)
	StandardButton_FlagMask        StandardButton = StandardButton(0x00000300)
	StandardButton_ButtonMask      StandardButton = StandardButton(0)
)
