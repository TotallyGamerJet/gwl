// +build darwin

package gwl

import(
	"github.com/totallygamerjet/cocoa"
)

func createPlatformWindow(config windowConfig) (ggl Window, err error) {
	return
}

func getStyle(config windowConfig) uint {
	var style = cocoa.NSWindowStyleMaskMiniaturizable
	if config.decorated {
		style |= cocoa.NSWindowStyleMaskTitled | cocoa.NSWindowStyleMaskClosable
	}
	if config.resizable {
		style |= cocoa.NSWindowStyleMaskResizable
	}
	if config.maximized {
		//TODO: fullscreen
	}
	return style
}
