//go:build impl

package faithtop

import (
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type SystemTrayImpl struct {
	tray *widgets.QSystemTrayIcon
}

func init() {
	newSystemTrayImpl = func() ISystemTray {
		return &SystemTrayImpl{
			tray: widgets.NewQSystemTrayIcon(nil),
		}
	}
}

func (s *SystemTrayImpl) Assign(v *ISystemTray) ISystemTray {
	*v = s
	return s
}

func (s *SystemTrayImpl) Icon(fileName string) ISystemTray {
	s.tray.SetIcon(gui.NewQIcon5(fileName))
	return s
}

func (s *SystemTrayImpl) Tooltip(tooltip string) ISystemTray {
	s.tray.SetToolTip(tooltip)
	return s
}

func (s *SystemTrayImpl) Menu(menu IMenu) ISystemTray {
	s.tray.SetContextMenu(menu.(*MenuImpl).menu)
	return s
}

func (s *SystemTrayImpl) Show() ISystemTray {
	s.tray.Show()
	return s
}

func (s *SystemTrayImpl) ShowMessage(title, message string, icon MessageIcon, millseconds int) ISystemTray {
	s.tray.ShowMessage(title, message, widgets.QSystemTrayIcon__MessageIcon(icon), millseconds)
	return s
}
