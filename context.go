package gwl

type context interface {
	SwapBuffers()
	MakeContextCurrent()
	DeleteContext(Window)
}
