package main

import (
	"gopkg.in/teh-cmc/go-sfml.v24/graphics"
	"gopkg.in/teh-cmc/go-sfml.v24/window"
	"os"
)

func menu() {
	menu_loop()
}

func menu_loop() {
	for window.SfWindow_isOpen(win) > 0 {
		menu_events()
		menu_render()
	}
}

func menu_events() {
	for window.SfWindow_pollEvent(win, eve) > 0 {
		/* Close window: exit */
		if eve.GetXtype() == window.SfEventType(window.SfEvtClosed) {
			os.Exit(0)
		}
		if eve.GetXtype() == window.SfEventType(window.SfEvtKeyPressed) {
			if (int(eve.GetKey().GetCode()) == window.SfKeyEscape) {
				os.Exit(0)
			}
		}
	}
}

func menu_render() {
		graphics.SfRenderWindow_clear(win, graphics.GetSfBlack())
		graphics.SfRenderWindow_display(win)
}
