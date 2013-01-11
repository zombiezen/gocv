package cv

// #include "highgui.h"
import "C"

import (
	"errors"
	"time"
	"unsafe"
)

// Capture is an image capture source.
type Capture struct {
	capture *C.CvCapture
}

// Release closes the capture source.
func (c Capture) Release() {
	do(func() {
		C.cvReleaseCapture(&c.capture)
	})
}

// QueryFrame returns a new frame from the capture source.
func (c Capture) QueryFrame() (*IplImage, error) {
	var image *C.IplImage
	do(func() {
		image = C.cvQueryFrame(c.capture)
	})
	if image == nil {
		return nil, errors.New("query failed")
	}
	// XXX: The pointer to this memory should not be garbage collected.
	return (*IplImage)(unsafe.Pointer(image)), nil
}

// CaptureFromCAM creates a new capture source for the given device.
func CaptureFromCAM(device int) (Capture, error) {
	var c *C.CvCapture
	do(func() {
		c = C.cvCaptureFromCAM(C.int(device))
	})
	if c == nil {
		return Capture{}, errors.New("Capture failed")
	}
	return Capture{c}, nil
}

// CaptureFromFile creates a new capture source for a given file.
func CaptureFromFile(filename string) (Capture, error) {
	s := C.CString(filename)
	defer C.free(unsafe.Pointer(s))

	var c *C.CvCapture
	do(func() {
		c = C.cvCaptureFromFile(s)
	})
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

// LoadImage creates an image from a file.
func LoadImage(name string, color int) (*IplImage, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	var image *C.IplImage
	do(func() {
		image = C.cvLoadImage(cname, C.int(color))
	})
	if image == nil {
		return nil, errors.New("LoadImage failed")
	}
	return (*IplImage)(unsafe.Pointer(image)), nil
}

// ShowImage displays img to the window called name.
func ShowImage(name string, img Arr) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	do(func() {
		C.cvShowImage(cname, img.arr())
	})
}

// Window flags
const (
	WINDOW_AUTOSIZE = C.CV_WINDOW_AUTOSIZE
)

// NamedWindow creates a new window called name.
func NamedWindow(name string, flags int) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	// TODO: Check result
	do(func() {
		C.cvNamedWindow(cname, C.int(flags))
	})
}

// WaitKey obtains key input from the user. If delay is non-zero, then the
// function will return if a key has not been hit before delay has elapsed.
func WaitKey(delay time.Duration) rune {
	var key rune
	do(func() {
		key = rune(C.cvWaitKey(C.int(delay.Nanoseconds() / 1e6)))
	})
	return key
}

// DestroyWindow will close the window called name.
func DestroyWindow(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	do(func() {
		C.cvDestroyWindow(cname)
	})
}
