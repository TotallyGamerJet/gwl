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
