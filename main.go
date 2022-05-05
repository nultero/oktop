package main

import (
	"ktop/kproc"
	"ktop/ktdata"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

/*
	TODO config somewhere
*/

func init() {
	tcell.SetEncodingFallback(tcell.EncodingFallbackASCII)
}

func main() {

	parseArgs(os.Args[1:])

	stt := ktdata.DefaultState()

	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}
	if err = screen.Init(); err != nil {
		panic(err)
	}

	screen.HideCursor()
	screen.SetStyle(stt.ColorTheme.MainStyle)
	screen.Clear()

	quit := make(chan struct{})
	go func() {
		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape, tcell.KeyEnter, tcell.KeyCtrlC, tcell.KeyCtrlQ:
					close(quit)
					return
				case tcell.KeyCtrlL:
					screen.Sync()
				}
			case *tcell.EventResize:
				stt.NeedsRedraw = true
				screen.Sync()
			}
		}
	}()

	// sty := styles.Matrix()

renderloop:
	for {
		select {
		case <-quit:
			break renderloop

		case <-time.After(stt.PollRate):
		}

		err := kproc.PollCPU(&stt)
		if err != nil {
			panic(err)
		}

		err = kproc.PollMem(&stt)
		if err != nil {
			panic(err)
		}

		if stt.NeedsRedraw {
			redraw(screen, &stt)
			stt.NeedsRedraw = false
		}

		if isDrawable(screen.Size()) {
			// stdDraw(screen, &stt)
			ioDraw(screen, &stt, quadTopRight)
			ioDraw(screen, &stt, quadBottomRight)
			ioDraw(screen, &stt, quadTopLeft)
			ioDraw(screen, &stt, quadBottomLeft)
		}

		screen.Show() // only calling this once ✓
		//  else {
		// 	invalidSzDraw(screen, sty)
		// }
	}

	screen.Fini()
}

func isDrawable(x, y int) bool {
	if x < 30 || y < 16 {
		return false
	}

	return true
}
