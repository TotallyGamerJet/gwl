// +build windows

package gwl

import (
	"fmt"
	"github.com/totallygamerjet/w32"
	"golang.org/x/sys/windows"
)

// PollEvents polls the system for os messages. This call doesn't block
func PollEvents() {
	msg := w32.MSG{}
	for w32.PeekMessage(&msg, 0, 0, 0, w32.PM_REMOVE) {
		w32.TranslateMessage(&msg)
		w32.DispatchMessage(&msg)
	}
}

/*func (g Window) ShowWindow() {
	w32.ShowWindow(w32.HWND(g), w32.SW_SHOW)
}*/

func (g Window) SwapBuffers() {
	w32.SwapBuffers(dc)
}

type win32Context struct {
	dc w32.HDC
	rc w32.HGLRC
}

func (w win32Context) swapBuffers() {
	w32.SwapBuffers(w.dc)
}

func createPlatformWindow(config windowConfig) (ggl Window, err error) {
	style := getStyle(config)

	inst := w32.GetModuleHandle(gglWindowClass)
	ggl = Window(w32.CreateWindowEx(
		0,
		windows.StringToUTF16Ptr(gglWindowClass),
		windows.StringToUTF16Ptr(config.title),
		style,
		w32.CW_USEDEFAULT, w32.CW_USEDEFAULT,
		int(config.width), int(config.height),
		0,
		0,
		inst,
		nil,
	))
	if ggl == 0 {
		err = fmt.Errorf("failed to create platform window: %d", w32.GetLastError())
		return
	}
	return
}

func getStyle(config windowConfig) uint {
	var style uint = w32.WS_VISIBLE

	if config.decorated {
		style |= w32.WS_SYSMENU | w32.WS_CAPTION | w32.WS_MINIMIZEBOX
	} else {
		style |= w32.WS_POPUP
	}
	if config.resizable {
		style |= w32.WS_THICKFRAME | w32.WS_MAXIMIZEBOX
	}
	if config.maximized {
		style |= w32.WS_MAXIMIZE
	}
	if config.iconified {
		style |= w32.WS_MINIMIZE
	}

	return style
}
