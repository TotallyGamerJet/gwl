package gwl

import "fmt"

// Window is a "pointer" to the window data.
//A zero value means something went wrong
type Window uintptr

func (g Window) set(do func(*windowData)) {
	err := setData(g, do)
	if err != nil {
		fmt.Printf("failed to set: %v\n", err)
	}
}

// SetOnMinimize sets the minimize callback. The callback is called whenever the window
//is minimized
func (g Window) SetOnMinimize(callback func(Window)) {
	g.set(func(win *windowData) { win.callbacks.onMinimize = callback })
}

// SetOnMaximize sets maximize callback. The callback is called whenever the window
//is maximized
func (g Window) SetOnMaximize(callback func(Window)) {
	g.set(func(win *windowData) { win.callbacks.onMaximize = callback })
}

// SetOnFocusChange sets the on focus change callback. The callback will be called whenever focus leaves
//or is returned to the window.
func (g Window) SetOnFocusChange(callback func(window Window, focused bool)) {
	g.set(func(win *windowData) { win.callbacks.onFocusChange = callback })
}

//SetShouldClose takes a bool to set the window's shouldClose field to
func (g Window) SetShouldClose(close bool) {
	g.set(func(win *windowData) { win.shouldClose = close })
}

// ShouldClose returns true if the window should close
func (g Window) ShouldClose() bool {
	return getData(g).shouldClose
}

//MakeContextCurrent makes the opengl context current
func (g Window) MakeContextCurrent() {
	getData(g).MakeContextCurrent()
}

//SwapBuffers swaps the system buffers
func (g Window) SwapBuffers() {
	getData(g).SwapBuffers()
}


type windowData struct {
	context
	window      Window
	shouldClose bool
	callbacks   struct {
		onMinimize    func(Window)
		onMaximize    func(Window)
		onFocusChange func(Window, bool)
	}
}

//windowConfig holds all the information about how to create a window
type windowConfig struct {
	title         string
	width, height uint
	resizable     bool
	maximized     bool
	decorated     bool
	iconified     bool
}

type windowFlag uint32

const (
	None      windowFlag = 0
	Resizable windowFlag = 1 << iota
	Maximized
	Decorated
	Iconified
)

func (bit windowFlag) has(flag windowFlag) bool {
	return bit&flag != 0
}

// CreateWindow creates a window according to the WindowOpts sent in
//It returns a GGLWindow and an error if something went wrong
func CreateWindow(title string, width, height uint, flags windowFlag) (Window, error) {
	config := windowConfig{
		title:     title,
		width:     width,
		height:    height,
		resizable: flags.has(Resizable),
		decorated: flags.has(Decorated),
		maximized: flags.has(Maximized),
		iconified: flags.has(Iconified),
	}
	w, err := createPlatformWindow(config)
	if err != nil {
		return w, err
	}
	return w, nil
}
