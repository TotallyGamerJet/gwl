// +build windows

package gwl

import (
	"fmt"
	"github.com/totallygamerjet/w32"
	"golang.org/x/sys/windows"
	"unsafe"
)

const gglWindowClass = "GGL Window Class"

func init() {
	registerWindowClass()
}

func registerWindowClass() {
	inst := w32.GetModuleHandle(gglWindowClass)

	wc := w32.WNDCLASSEX{
		WndProc:   windows.NewCallback(wndProc),
		Instance:  inst,
		ClassName: windows.StringToUTF16Ptr(gglWindowClass),
		Size:      uint32(unsafe.Sizeof(w32.WNDCLASSEX{})),
	}

	if atom := w32.RegisterClassEx(&wc); atom == 0 {
		panic("failed to register window class")
	}

}

func wndProc(hwnd w32.HWND, msg uint32, wparam, lparam uintptr) uintptr {
	window := getData(Window(hwnd))
	switch msg {
	case w32.WM_CREATE:
		//https://samulinatri.com/blog/win32-opengl-context/
		pfd := w32.PIXELFORMATDESCRIPTOR{
			uint16(unsafe.Sizeof(w32.PIXELFORMATDESCRIPTOR{})), //  size of this pfd
			1, // version number
			w32.PFD_DRAW_TO_WINDOW | // support window
				w32.PFD_SUPPORT_OPENGL | // support OpenGL
				w32.PFD_DOUBLEBUFFER, // double buffered
			w32.PFD_TYPE_RGBA, // RGBA type
			24,                // 24-bit color depth
			0, 0, 0, 0, 0, 0,  // color bits ignored
			0,          // no alpha buffer
			0,          // shift bit ignored
			0,          // no accumulation buffer
			0, 0, 0, 0, // accum bits ignored
			32,                 // 32-bit z-buffer
			0,                  // no stencil buffer
			0,                  // no auxiliary buffer
			w32.PFD_MAIN_PLANE, // main layer
			0,                  // reserved
			0, 0, 0,            // layer masks ignored
		}
		dc := w32.GetDC(hwnd)

		pf := w32.ChoosePixelFormat(dc, &pfd)
		w32.SetPixelFormat(dc, pf, &pfd)

		rc := w32.WglCreateContext(dc)
		setData(Window(hwnd), func(data *windowData) { data.context = win32Context{dc: dc, rc: rc} })
		return 0
	case w32.WM_DESTROY:
		context, ok := window.context.(win32Context)
		if !ok {
			fmt.Println("Platform err: wrong context")
		}
		Window(hwnd).SetShouldClose(true)
		w32.WglMakeCurrent(0, 0) //make it not current anymore
		w32.ReleaseDC(hwnd, context.dc)
		w32.WglDeleteContext(context.rc)
		w32.PostQuitMessage(0)
		return 0
	case w32.WM_SIZE:
		iconified := wparam == w32.SIZE_MINIMIZED
		if iconified && window.callbacks.onMinimize != nil {
			window.callbacks.onMinimize(Window(hwnd))
		}
		maximized := wparam == w32.SIZE_MAXIMIZED
		if maximized && window.callbacks.onMaximize != nil {
			window.callbacks.onMaximize(Window(hwnd))
		}
		return 0
	case w32.WM_SETFOCUS:
		if window.callbacks.onFocusChange != nil {
			window.callbacks.onFocusChange(Window(hwnd), true)
		}
		//if (window->cursorMode == GLFW_CURSOR_DISABLED)
		//	disableCursor(window);
		return 0
	case w32.WM_KILLFOCUS:
		//if (window->cursorMode == GLFW_CURSOR_DISABLED)
		//     enableCursor(window);
		//
		//if (window->monitor && window->autoIconify)
		//     _glfwPlatformIconifyWindow(window);
		if window.callbacks.onFocusChange != nil {
			window.callbacks.onFocusChange(Window(hwnd), false)
		}
		return 0
	case w32.WM_MOUSELEAVE:
		if window.callbacks.onMouseEnter != nil {
			window.callbacks.onMouseEnter(Window(hwnd), false)
		}
		return 0
	}

	/*err = setData(window) //make sure window is updated
	if err != nil {
		fmt.Printf("failed to set window: %v\n", err)
	}*/
	return w32.DefWindowProc(hwnd, msg, wparam, lparam)
}
