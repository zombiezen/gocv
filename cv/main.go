package cv

import (
	"runtime"
)

// OpenCV has some issues with multiple threads.  To overcome this, we use a sneaky approach documented here:
// http://code.google.com/p/go-wiki/wiki/LockOSThread

func init() {
	runtime.LockOSThread()
}

// Main must be called from the program's main function.
func Main() {
	for f := range mainfunc {
		f()
	}
}

var mainfunc = make(chan func())

func do(f func()) {
	done := make(chan struct{})
	mainfunc <- func() {
		f()
		close(done)
	}
	<-done
}
