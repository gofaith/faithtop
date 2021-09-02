package faithtop

type ITextBrowser interface {
	IWidget
	Text(s string)ITextBrowser
	Interaction(flags TextInteractionFlag)ITextBrowser
}

var newTextBrowserImpl func()ITextBrowser

func TextBrowser()ITextBrowser{
	return newTextBrowserImpl()
}

func TextBrowser2(s string)ITextBrowser{
	return newTextBrowserImpl().Text(s)
}