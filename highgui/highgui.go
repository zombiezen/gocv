package highgui

// #cgo CFLAGS: -I/usr/include/opencv -Wno-error
// #cgo LDFLAGS: -lcv -lhighgui
// #include "highgui.h"
import "C"

import (
	"bitbucket.org/zombiezen/gocv/cv"
	"errors"
	"unsafe"
)

type Capture struct {
	capture *C.CvCapture
}

func (c Capture) Release() {
	C.cvReleaseCapture(&c.capture)
}

func (c Capture) QueryFrame() (*cv.IplImage, error) {
	image := C.cvQueryFrame(c.capture)
	if image == nil {
		return nil, errors.New("query failed")
	}
	// XXX: The pointer to this memory should not be garbage collected.
	return (*cv.IplImage)(unsafe.Pointer(image)), nil
}

func CaptureFromCAM(device int) (Capture, error) {
	c := C.cvCaptureFromCAM(C.int(device))
	if c == nil {
		return Capture{}, errors.New("Capture failed")
	}
	return Capture{c}, nil
}

func CaptureFromFile(filename string) (Capture, error) {
	s := C.CString(filename)
	defer C.free(unsafe.Pointer(s))

	c := C.cvCaptureFromFile(s)
	if c == nil {
		return Capture{}, errors.New("Capture failed")
	}
	return Capture{c}, nil
}
