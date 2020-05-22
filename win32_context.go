// +build windows

package gwl

import (
	"github.com/totallygamerjet/w32"
)

type win32Context struct {
	dc w32.HDC
	rc w32.HGLRC
}

func (w win32Context) MakeContextCurrent() {
	w32.WglMakeCurrent(w.dc, w.rc)
}

func (w win32Context) SwapBuffers() {
	w32.SwapBuffers(w.dc)
}

func (w win32Context) DeleteContext(win Window) {
	w32.WglMakeCurrent(0, 0) //make it not current anymore
	w32.ReleaseDC(w32.HWND(win), w.dc)
	w32.WglDeleteContext(w.rc)
}
