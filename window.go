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

// SetCallbacks is used to set all the callbacks. Multiple callbacks may be set in one call
// to SetCallback.
func (g Window) SetCallbacks(do func(*Callbacks)) {
	err := setData(g, func(data *windowData) {
		do(&data.callbacks)
	})
	if err != nil {
		fmt.Printf("failed to set: %v\n", err)
	}
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
	getData(g).context.MakeContextCurrent()
}

//SwapBuffers swaps the system buffers
func (g Window) SwapBuffers() {
	getData(g).context.SwapBuffers()
}

//Callbacks holds the struct of functions to that will be called on a callback
type Callbacks struct {
	//OnMinimize will be called when the window is minimized
	OnMinimize func(Window)
	//OnMaximize will be called when the window is maximized
	OnMaximize func(Window)
	//OnFocusChange will be called when the window gains or loses focus
	OnFocusChange func(Window, bool)
	//OnMouseEnter will be called when the mouse leaves and enters the window
	OnMouseEnter func(Window, bool)
}

type windowData struct {
	context     context
	window      Window
	shouldClose bool
	callbacks   Callbacks
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

const None windowFlag = 0 //separate so that iota starts at 0
const (
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
