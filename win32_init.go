// +build windows

package gwl

import (
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
			Size:    uint16(unsafe.Sizeof(w32.PIXELFORMATDESCRIPTOR{})), //  size of this pfd
			Version: 1,                                                  // version number
			DwFlags: w32.PFD_DRAW_TO_WINDOW | // support window
				w32.PFD_SUPPORT_OPENGL | // support OpenGL
				w32.PFD_DOUBLEBUFFER, // double buffered
			IPixelType: w32.PFD_TYPE_RGBA, // RGBA type
			ColorBits:  24,
			DepthBits:  32,
			ILayerType: w32.PFD_MAIN_PLANE,
		}
		dc := w32.GetDC(hwnd)

		pf := w32.ChoosePixelFormat(dc, &pfd)
		w32.SetPixelFormat(dc, pf, &pfd)

		rc := w32.WglCreateContext(dc)
		Window(hwnd).set(func(data *windowData) { data.context = win32Context{dc: dc, rc: rc} })
		return 0
	case w32.WM_DESTROY:
		context := window.context
		Window(hwnd).SetShouldClose(true)
		context.DeleteContext(Window(hwnd))
		w32.PostQuitMessage(0)
		return 0
	case w32.WM_SIZE:
		iconified := wparam == w32.SIZE_MINIMIZED
		if iconified && window.callbacks.OnMinimize != nil {
			window.callbacks.OnMinimize(Window(hwnd))
		}
		maximized := wparam == w32.SIZE_MAXIMIZED
		if maximized && window.callbacks.OnMaximize != nil {
			window.callbacks.OnMaximize(Window(hwnd))
		}
		return 0
	case w32.WM_SETFOCUS:
		if window.callbacks.OnFocusChange != nil {
			window.callbacks.OnFocusChange(Window(hwnd), true)
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
		if window.callbacks.OnFocusChange != nil {
			window.callbacks.OnFocusChange(Window(hwnd), false)
		}
		return 0
	case w32.WM_MOUSELEAVE:
		if window.callbacks.OnMouseEnter != nil {
			window.callbacks.OnMouseEnter(Window(hwnd), false)
		}
		return 0
	}

	/*err = setData(window) //make sure window is updated
	if err != nil {
		fmt.Printf("failed to set window: %v\n", err)
	}*/
	return w32.DefWindowProc(hwnd, msg, wparam, lparam)
}
