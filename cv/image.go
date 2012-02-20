package cv

// #cgo CFLAGS: -I/usr/include/opencv -Wno-error
// #cgo LDFLAGS: -lcv -lcxcore
// #include "cv.h"
import "C"

import (
	"unsafe"
)

// IplImage stores an image.
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

func (i *IplImage) Clone() *IplImage {
	return (*IplImage)(unsafe.Pointer(C.cvCloneImage((*C.IplImage)(unsafe.Pointer(i)))))
}

func (i *IplImage) SetCOI(channel int) {
	C.cvSetImageCOI((*C.IplImage)(unsafe.Pointer(i)), C.int(channel))
}

func (i *IplImage) SetROI(r Rect) {
	C.cvSetImageROI((*C.IplImage)(unsafe.Pointer(i)),
		C.CvRect{C.int(r.X), C.int(r.Y), C.int(r.Width), C.int(r.Height)})
}

func (i *IplImage) Release() {
	C.cvReleaseImage((**C.IplImage)(unsafe.Pointer(&i)))
}