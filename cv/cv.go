package cv

// #cgo CFLAGS: -I/usr/include/opencv -Wno-error
// #cgo LDFLAGS: -lcv -lcxcore
// #include "cv.h"
import "C"

import (
	"unsafe"
)

type Size struct {
	Width  int
	Height int
}

type IplImage struct {
	NSize           int
	ID              int
	NChannels       int
	AlphaChannel    int
	Depth           int
	ColorModel      [4]int8
	ChannelSeq      [4]int8
	DataOrder       int
	Origin          int
	Align           int
	Width           int
	Height          int
	ROI             uintptr // TODO
	MaskROI         *IplImage
	ImageId         uintptr
	TileInfo        uintptr // TODO
	ImageSize       int
	ImageData       uintptr // TODO
	WidthStep       int
	BorderMode      [4]int
	BorderConst     [4]int
	ImageDataOrigin uintptr // TODO
}

type Arr interface {
	Size() Size
	arr() unsafe.Pointer
}

func NewImage(size Size, depth, channels int) *IplImage {
	// XXX: This should be garbage-collected by Go.
	i := C.cvCreateImage(C.CvSize{C.int(size.Width), C.int(size.Height)}, C.int(depth), C.int(channels))
	return (*IplImage)(unsafe.Pointer(i))
}

func (i *IplImage) arr() unsafe.Pointer {
	return unsafe.Pointer(i)
}

func (i *IplImage) Size() Size {
	s := C.cvGetSize(i.arr())
	return Size{int(s.width), int(s.height)}
}

func (i *IplImage) Release() {
	C.cvReleaseImage((**C.IplImage)(unsafe.Pointer(&i)))
}

const (
	THRESH_BINARY = C.CV_THRESH_BINARY
	THRESH_BINARY_INV = C.CV_THRESH_BINARY_INV
)

func Threshold(src, dst Arr, thresh, maxVal float64, thresholdType int) float64 {
	return float64(C.cvThreshold(src.arr(), dst.arr(), C.double(thresh), C.double(maxVal), C.int(thresholdType)))
}
