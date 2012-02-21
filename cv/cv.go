/*
	Package cv provides idiomatic bindings to OpenCV.
*/
package cv

// #cgo CFLAGS: -Wno-error
// #cgo pkg-config: opencv
// #include "cv.h"
import "C"

import (
	"unsafe"
)

// Point holds a 2D integer point.
type Point struct {
	X, Y int
}

// Point2D32f holds a 2D float32 point.
type Point2D32f struct {
	X, Y float32
}

// Point3D32f holds a 3D float32 point.
type Point3D32f struct {
	X, Y, Z float32
}

// Point2D64f holds a 2D float64 point.
type Point2D64f struct {
	X, Y float64
}

// Point3D64f holds a 3D float64 point.
type Point3D64f struct {
	X, Y, Z float64
}

// Size holds a 2D integer size.
type Size struct {
	Width, Height int
}

// Size2D32f holds a 2D float32 size.
type Size2D32f struct {
	Width, Height float32
}

// Rect holds a 2D integer rectangle.
type Rect struct {
	X, Y, Width, Height int
}

// Scalar holds up to 4 float64s. OpenCV uses scalars for colors.
type Scalar [4]float64

func (s Scalar) cvScalar() C.CvScalar {
	return C.CvScalar{[4]C.double{C.double(s[0]), C.double(s[1]), C.double(s[2]), C.double(s[3])}}
}

// Types of thresholding
const (
	THRESH_BINARY     = C.CV_THRESH_BINARY
	THRESH_BINARY_INV = C.CV_THRESH_BINARY_INV
	THRESH_TRUNC      = C.CV_THRESH_TRUNC
	THRESH_TOZERO     = C.CV_THRESH_TOZERO
	THRESH_TOZERO_INV = C.CV_THRESH_TOZERO_INV
	THRESH_MASK       = C.CV_THRESH_MASK
	THRESH_OTSU       = C.CV_THRESH_OTSU
)

// Threshold applies a fixed-level threshold to a grayscale image.
func Threshold(src, dst Arr, thresh, maxVal float64, thresholdType int) float64 {
	return float64(C.cvThreshold(src.arr(), dst.arr(), C.double(thresh), C.double(maxVal), C.int(thresholdType)))
}

// Filtering algorithms
const (
	GAUSSIAN_5x5 = C.CV_GAUSSIAN_5x5
)

// PyrDown smooths and down-samples the input image.
func PyrDown(src, dst Arr, filter int) {
	C.cvPyrDown(src.arr(), dst.arr(), C.int(filter))
}

// PyrUp up-samples the input image and smooths the result.
func PyrUp(src, dst Arr, filter int) {
	C.cvPyrDown(src.arr(), dst.arr(), C.int(filter))
}

// IplConvKernel is a convolution kernel.
type IplConvKernel struct {
	NCols   int
	NRows   int
	AnchorX int
	AnchorY int
	Values  uintptr // TODO
	NShiftR int
}

// Dilate applies a maximum filter to the input image one or more times.  If
// element is nil, a 3x3 rectangular element is used.
func Dilate(src, dst Arr, element *IplConvKernel, iterations int) {
	C.cvDilate(src.arr(), dst.arr(), (*C.IplConvKernel)(unsafe.Pointer(element)), C.int(iterations))
}

// Erode applies a minimum filter to the input image one or more times.  If
// element is nil, a 3x3 rectangular element is used.
func Erode(src, dst Arr, element *IplConvKernel, iterations int) {
	C.cvErode(src.arr(), dst.arr(), (*C.IplConvKernel)(unsafe.Pointer(element)), C.int(iterations))
}
