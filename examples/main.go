package main

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"gwl"
	"log"
	"runtime"
)

func main() {
	runtime.LockOSThread()
	win, err := gwl.CreateWindow("My Evil Prog", 1080, 720, gwl.Decorated|gwl.Resizable)
	if err != nil {
		panic(err)
	}

	win.SetOnMinimize(func(window gwl.Window) { fmt.Println("I was just minimized") })
	win.SetOnMaximize(func(window gwl.Window) { fmt.Println("I was just maximized") })
	win.SetOnFocusChange(func(window gwl.Window, focused bool) {
		if focused {
			fmt.Println("Gained focus!")
		} else {
			fmt.Println("Lost focus :(")
		}
	})

	win.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		fmt.Println("init failed: ", err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	for !win.ShouldClose() {
		gl.ClearColor(1, 0, 0, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		win.SwapBuffers()
		gwl.PollEvents()
	}
}
