package faithtop

type ISystemTray interface {
	Assign(v *ISystemTray) ISystemTray
	/** set icon from file
	e.g.
	Icon(":/qml/logo.png")
	**/
	Icon(fileName string) ISystemTray
	Tooltip(tooltip string) ISystemTray
	Menu(menu IMenu) ISystemTray
	Show()ISystemTray
	ShowMessage(title, message string, icon MessageIcon, millseconds int) ISystemTray
}

type MessageIcon int64

const (
	MessageIcon_NoIcon      MessageIcon = MessageIcon(0)
	MessageIcon_Information MessageIcon = MessageIcon(1)
	MessageIcon_Warning     MessageIcon = MessageIcon(2)
	MessageIcon_Critical    MessageIcon = MessageIcon(3)
)

var newSystemTrayImpl func() ISystemTray

func SystemTray() ISystemTray {
	return newSystemTrayImpl()
}
