package main

import (
	"gopkg.in/teh-cmc/go-sfml.v24/graphics"
	"gopkg.in/teh-cmc/go-sfml.v24/window"
)

func menu() {
	menu_events()
}
func menu_events() {
	for window.SfWindow_pollEvent(win, eve) > 0 {
		/* Close window: exit */
		if eve.GetXtype() == window.SfEventType(window.SfEvtClosed) {
			return
		}
	}
}

func menu_render() {
		graphics.SfRenderWindow_clear(win, graphics.GetSfBlack())
		graphics.SfRenderWindow_display(win)
}
