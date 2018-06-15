package htmltermui

import (
	cc "ml253/cursesgraphic"
	"strings"

	color "golang.org/x/image/colornames"
)

type mainThread func(tui *TermUI)

type TermUI struct {
	View *cc.Canvas
}

var mainThreadFunction mainThread

func NewTerminalUI(fn mainThread, screenWidth, screenHeight int, fullscreen bool) {
	//tui := new(TermUI);
	//tui.View = cc.NewFullscreenCanvas();
	mainThreadFunction = fn
	cc.CurseGraphicStart(newTermUIContinued, screenWidth, screenHeight, fullscreen)

	//cc.Raw(true);

}

func newTermUIContinued(ctx *cc.Canvas) {
	tui := new(TermUI)
	tui.View = ctx

	mainThreadFunction(tui)
}

func (*TermUI) EndTerminalUI() {
	//cc.EndCanvas();
}

func (t *TermUI) PopupMessage(message string, positiveMessage bool) {
	totalSize := len(message)
	if totalSize < 30 {
		totalSize = 30
	}
	t.View.Save()
	t.View.SetTranslate(0, 0)
	if positiveMessage {
		t.View.SetColor(color.Green, color.Black)
	} else {
		t.View.SetColor(color.Red, color.Black)
	}

	t.View.SetFillChar('~')
	t.View.FillRect((t.View.Width()/2)-(totalSize/2+1),
		t.View.Height()/2-1,
		totalSize+2, 1)
	t.View.FillText(
		"~"+message+strings.Repeat(" ", totalSize-len(message))+"~",
		(t.View.Width()/2)-(totalSize/2+1),
		t.View.Height()/2,
	)
	t.View.FillText(
		"~Press any key to continue..."+strings.Repeat(" ", totalSize-28)+"~",
		(t.View.Width()/2)-(totalSize/2+1),
		t.View.Height()/2+1,
	)
	t.View.SetFillChar('~')
	t.View.FillRect((t.View.Width()/2)-(totalSize/2+1),
		t.View.Height()/2+2,
		totalSize+2, 1)

	//t.View.GetCanvas().GetChar();
	cc.SwapBuffers()
	t.View.GetCharCode()
	t.View.Restore()
	cc.SwapBuffers()
}

func (t *TermUI) DebugPrinter(message string) {
	/*t.View.GetCanvas().MovePrint(t.View.Height()-3,
	  t.View.Width() - t.View.Width() / 4,
	  ("Debug:" + message))*/
	t.View.Save()
	t.View.SetTranslate(0, 0)
	t.View.FillText("Debug:"+message,
		t.View.Width()-t.View.Width()/4,
		t.View.Height()-3)
	t.View.Restore()
}
