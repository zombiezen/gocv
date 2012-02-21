package cv

// #include "highgui.h"
import "C"

import (
	"errors"
	"time"
	"unsafe"
)

type Capture struct {
	capture *C.CvCapture
}

func (c Capture) Release() {
	C.cvReleaseCapture(&c.capture)
}

func (c Capture) QueryFrame() (*IplImage, error) {
	image := C.cvQueryFrame(c.capture)
	if image == nil {
		return nil, errors.New("query failed")
	}
	// XXX: The pointer to this memory should not be garbage collected.
	return (*IplImage)(unsafe.Pointer(image)), nil
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

const (
	LOAD_IMAGE_COLOR     = C.CV_LOAD_IMAGE_COLOR
	LOAD_IMAGE_GRAYSCALE = C.CV_LOAD_IMAGE_GRAYSCALE
	LOAD_IMAGE_UNCHANGED = C.CV_LOAD_IMAGE_UNCHANGED
)

func LoadImage(name string, color int) (*IplImage, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	image := C.cvLoadImage(cname, C.int(color))
	if image == nil {
		return nil, errors.New("LoadImage failed")
	}
	return (*IplImage)(unsafe.Pointer(image)), nil
}

func ShowImage(name string, img Arr) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.cvShowImage(cname, img.arr())
}

// Window flags
const (
	WINDOW_AUTOSIZE = C.CV_WINDOW_AUTOSIZE
)

func NamedWindow(name string, flags int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	// TODO: Check result
	C.cvNamedWindow(cname, C.int(flags))
}

func WaitKey(delay time.Duration) rune {
	return rune(C.cvWaitKey(C.int(delay.Nanoseconds() / 1e6)))
}

func DestroyWindow(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.cvDestroyWindow(cname)
}
