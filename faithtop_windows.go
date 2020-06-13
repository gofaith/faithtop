package faithtop

var currentWindow *FWindow

func Main() {
	if currentWindow != nil {
		currentWindow.w.Run()
	} else {
		panic("no window to show")
	}
}
