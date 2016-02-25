/*
	Package cv provides idiomatic bindings to OpenCV.
*/
package cv

// #cgo CFLAGS: -Wno-error
// #cgo !windows pkg-config: opencv
// #cgo windows CFLAGS: -IC:/opencv/build/include -IC:/opencv/build/include/opencv -IC:/opencv/build/include/opencv2
// #cgo windows LDFLAGS: -LC:/opencv/build/x86/vc11/bin -lopencv_calib3d2411 -lopencv_contrib2411 -lopencv_core2411 -lopencv_features2d2411 -lopencv_flann2411 -lopencv_gpu2411 -lopencv_highgui2411 -lopencv_imgproc2411 -lopencv_legacy2411 -lopencv_ml2411 -lopencv_nonfree2411 -lopencv_objdetect2411 -lopencv_photo2411 -lopencv_stitching2411 -lopencv_video2411 -lopencv_videostab2411 -lopencv_ffmpeg2411 
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

// PointRect holds 4 points and a Rect object
type PointRect struct {
	Points [4]Point
	Rect   Rect
}

// Scalar holds up to 4 float64s. OpenCV uses scalars for colors.
type Scalar [4]float64

func (s Scalar) cvScalar() C.CvScalar {
	return C.CvScalar{[4]C.double{C.double(s[0]), C.double(s[1]), C.double(s[2]), C.double(s[3])}}
}

// And performs a bitwise AND on src1 and src2 and stores into dst.
func And(src1, src2, dst, mask Arr) {
	do(func() {
		if mask != nil {
			C.cvAnd(src1.arr(), src2.arr(), dst.arr(), mask.arr())
		} else {
			C.cvAnd(src1.arr(), src2.arr(), dst.arr(), nil)
		}
	})
}

// Or performs a bitwise OR on src1 and src2 and stores into dst
func Or(src1, src2, dst, mask Arr) {
	do(func() {
		if mask != nil {
			C.cvOr(src1.arr(), src2.arr(), dst.arr(), mask.arr())
		} else {
			C.cvOr(src1.arr(), src2.arr(), dst.arr(), nil)
		}
	})
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
	var result float64
	do(func() {
		result = float64(C.cvThreshold(src.arr(), dst.arr(), C.double(thresh), C.double(maxVal), C.int(thresholdType)))
	})
	return result
}

// Color space conversions
const (
	RGB2GRAY = C.CV_RGB2GRAY

	BGR2XYZ = C.CV_BGR2XYZ
	RGB2XYZ = C.CV_RGB2XYZ
	XYZ2BGR = C.CV_XYZ2BGR
	XYZ2RGB = C.CV_XYZ2RGB

	BGR2YCrCb = C.CV_BGR2YCrCb
	RGB2YCrCb = C.CV_RGB2YCrCb
	YCrCb2BGR = C.CV_YCrCb2BGR
	YCrCb2RGB = C.CV_YCrCb2RGB

	BGR2HSV = C.CV_BGR2HSV
	RGB2HSV = C.CV_RGB2HSV
	HSV2BGR = C.CV_HSV2BGR
	HSV2RGB = C.CV_HSV2RGB

	BGR2HLS = C.CV_BGR2HLS
	RGB2HLS = C.CV_RGB2HLS
	HLS2BGR = C.CV_HLS2BGR
	HLS2RGB = C.CV_HLS2RGB

	BGR2Lab = C.CV_BGR2Lab
	RGB2Lab = C.CV_RGB2Lab
	Lab2BGR = C.CV_Lab2BGR
	Lab2RGB = C.CV_Lab2RGB

	BGR2Luv = C.CV_BGR2Luv
	RGB2Luv = C.CV_RGB2Luv
	Luv2BGR = C.CV_Luv2BGR
	Luv2RGB = C.CV_Luv2RGB

	BayerBG2BGR = C.CV_BayerBG2BGR
	BayerGB2BGR = C.CV_BayerGB2BGR
	BayerRG2BGR = C.CV_BayerRG2BGR
	BayerGR2BGR = C.CV_BayerGR2BGR
	BayerBG2RGB = C.CV_BayerBG2RGB
	BayerGB2RGB = C.CV_BayerGB2RGB
	BayerRG2RGB = C.CV_BayerRG2RGB
	BayerGR2RGB = C.CV_BayerGR2RGB
)

// CvtColor converts an image from one color space to another.
func CvtColor(src, dst Arr, code int) {
	do(func() {
		C.cvCvtColor(src.arr(), dst.arr(), C.int(code))
	})
}

// Split copies each of src's channels into the destinations.
//
// If the source array has N channels then if the first N destination channels
// are not nil, they all are extracted from the source array; if only a single
// destination channel of the first N is not nil, this particular channel is
// extracted; otherwise an error is raised. The rest of the destination channels
// (beyond the first N) must always be nil.
func Split(src, dst0, dst1, dst2, dst3 Arr) {
	var p0, p1, p2, p3 unsafe.Pointer
	if dst0 != nil {
		p0 = dst0.arr()
	}
	if dst1 != nil {
		p1 = dst1.arr()
	}
	if dst2 != nil {
		p2 = dst2.arr()
	}
	if dst3 != nil {
		p3 = dst3.arr()
	}
	do(func() {
		C.cvSplit(src.arr(), p0, p1, p2, p3)
	})
}

// Filtering algorithms
const (
	GAUSSIAN_5x5 = C.CV_GAUSSIAN_5x5
)

// PyrDown smooths and down-samples the input image.
func PyrDown(src, dst Arr, filter int) {
	do(func() {
		C.cvPyrDown(src.arr(), dst.arr(), C.int(filter))
	})
}

// PyrUp up-samples the input image and smooths the result.
func PyrUp(src, dst Arr, filter int) {
	do(func() {
		C.cvPyrDown(src.arr(), dst.arr(), C.int(filter))
	})
}

const (
	SHAPE_RECT = C.CV_SHAPE_RECT
)

// IplConvKernel is a convolution kernel.
type IplConvKernel struct {
	NCols   int
	NRows   int
	AnchorX int
	AnchorY int
	Shape   int
	Values  uintptr // TODO
	NShiftR int
}

func ReleaseStructuringElement(element *IplConvKernel) {
	do(func() {
		C.cvReleaseStructuringElement((**C.IplConvKernel)(unsafe.Pointer(element)))
	})
}

// Morphology constants
type Morphology int

const (
	MORPH_OPEN     Morphology = C.CV_MOP_OPEN
	MORPH_CLOSE    Morphology = C.CV_MOP_CLOSE
	MORPH_GRADIENT Morphology = C.CV_MOP_GRADIENT
	MORPH_TOPHAT   Morphology = C.CV_MOP_TOPHAT
	MORPH_BLACKHAT Morphology = C.CV_MOP_BLACKHAT
)

// Dilate applies a maximum filter to the input image one or more times.  If
// element is nil, a 3x3 rectangular element is used.
func Dilate(src, dst Arr, element *IplConvKernel, iterations int) {
	do(func() {
		C.cvDilate(src.arr(), dst.arr(), (*C.IplConvKernel)(unsafe.Pointer(element)), C.int(iterations))
	})
}

// Erode applies a minimum filter to the input image one or more times.  If
// element is nil, a 3x3 rectangular element is used.
func Erode(src, dst Arr, element *IplConvKernel, iterations int) {
	do(func() {
		C.cvErode(src.arr(), dst.arr(), (*C.IplConvKernel)(unsafe.Pointer(element)), C.int(iterations))
	})
}

func MorphologyEx(src, dst, temp Arr, element *IplConvKernel, operation Morphology, iterations int) {
	do(func() {
		C.cvMorphologyEx(src.arr(), dst.arr(), temp.arr(), (*C.IplConvKernel)(unsafe.Pointer(element)), C.int(operation), C.int(iterations))
	})
}
