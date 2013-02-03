package cv

// #include "cv.h"
import "C"

import (
	"image"
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

// NewImage creates a new image.
func NewImage(size Size, depth, channels int) *IplImage {
	// XXX: This should be garbage-collected by Go.
	var i *C.IplImage
	do(func() {
		i = C.cvCreateImage(C.CvSize{C.int(size.Width), C.int(size.Height)}, C.int(depth), C.int(channels))
	})
	return (*IplImage)(unsafe.Pointer(i))
}

func (i *IplImage) arr() unsafe.Pointer {
	return unsafe.Pointer(i)
}

// Size returns the width and height of the image.
func (i *IplImage) Size() Size {
	var s C.CvSize
	do(func() {
		s = C.cvGetSize(i.arr())
	})
	return Size{int(s.width), int(s.height)}
}

// Clone returns an image that has a copy of the i's data.
func (i *IplImage) Clone() *IplImage {
	var ii *C.IplImage
	do(func() {
		ii = C.cvCloneImage((*C.IplImage)(unsafe.Pointer(i)))
	})
	return (*IplImage)(unsafe.Pointer(ii))
}

// SetCOI sets the image's channel of interest.
func (i *IplImage) SetCOI(channel int) {
	do(func() {
		C.cvSetImageCOI((*C.IplImage)(unsafe.Pointer(i)), C.int(channel))
	})
}

// SetROI sets the image's region of interest.
func (i *IplImage) SetROI(r Rect) {
	do(func() {
		C.cvSetImageROI((*C.IplImage)(unsafe.Pointer(i)),
			C.CvRect{C.int(r.X), C.int(r.Y), C.int(r.Width), C.int(r.Height)})
	})
}

// Release destroys the memory associated with the image.
func (i *IplImage) Release() {
	do(func() {
		C.cvReleaseImage((**C.IplImage)(unsafe.Pointer(&i)))
	})
}

// ImageToCV converts a Go image (from the image package) into an IplImage.
// Only the RGB components are preserved.
func ImageToCV(m image.Image) *IplImage {
	const nchannels = 3

	bd := m.Bounds()
	msize := bd.Size()
	data := make([]byte, msize.X*msize.Y*nchannels)
	for y := bd.Min.Y; y < bd.Max.Y; y++ {
		for x := bd.Min.X; x < bd.Max.X; x++ {
			r, g, b, _ := m.At(x, y).RGBA()
			data[(y*msize.X+x)*nchannels+0] = byte(b >> 8)
			data[(y*msize.X+x)*nchannels+1] = byte(g >> 8)
			data[(y*msize.X+x)*nchannels+2] = byte(r >> 8)
		}
	}
	ipl := NewImage(Size{msize.X, msize.Y}, 8, nchannels)
	SetData(ipl, data, msize.X*nchannels)
	return ipl
}
