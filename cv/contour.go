package cv

// #include "cv.h"
import "C"

import (
	"errors"
	"unsafe"
)

// Contour retrieval mode
const (
	RETR_EXTERNAL = C.CV_RETR_EXTERNAL
	RETR_LIST     = C.CV_RETR_LIST
	RETR_CCOMP    = C.CV_RETR_CCOMP
	RETR_TREE     = C.CV_RETR_TREE
)

// Contour approximation method
const (
	CHAIN_CODE             = C.CV_CHAIN_CODE
	CHAIN_APPROX_NONE      = C.CV_CHAIN_APPROX_NONE
	CHAIN_APPROX_SIMPLE    = C.CV_CHAIN_APPROX_SIMPLE
	CHAIN_APPROX_TC89_L1   = C.CV_CHAIN_APPROX_TC89_L1
	CHAIN_APPROX_TC89_KCOS = C.CV_CHAIN_APPROX_TC89_KCOS
	LINK_RUNS              = C.CV_LINK_RUNS
)

// FindContours finds contour lines in a 8-bit, single-channel image. The
// recommended mode and method are RETR_LIST and CHAIN_APPROX_SIMPLE.  offset is
// added to every point in the contour.
func FindContours(image Arr, storage MemStorage, mode, method int, offset Point) (Seq, error) {
	var seq Seq
	result := C.cvFindContours(image.arr(), storage.s, &seq.seq, C.int(unsafe.Sizeof(C.CvContour{})), C.int(mode), C.int(method), C.CvPoint{C.int(offset.X), C.int(offset.Y)})
	if result < 0 {
		// TODO: Get error string
		return Seq{}, errors.New("FindContours failed")
	}
	return seq, nil
}

// Polygon approximation methods
const (
	POLY_APPROX_DP = C.CV_POLY_APPROX_DP
)

// ApproxPoly approximates polygonal curves. POLY_APPROX_DP is the only method
// supported. parameter is the desired approximation accuracy. parameter2
// should be zero to indicate only the given contour.
func ApproxPoly(srcSeq Seq, storage MemStorage, method int, parameter float64, parameter2 int) Seq {
	return Seq{C.cvApproxPoly(unsafe.Pointer(srcSeq.seq), C.int(unsafe.Sizeof(C.CvContour{})), storage.s, C.int(method), C.double(parameter), C.int(parameter2))}
}

// ContourArea returns the area inside contour. If oriented is true, then a
// negative area is returned if the contour is counter-clockwise.
func ContourArea(contour Arr, slice Slice, oriented bool) float64 {
	var corient C.int
	if oriented {
		corient = 1
	} else {
		corient = 0
	}
	return float64(C.cvContourArea(contour.arr(), C.CvSlice{C.int(slice.Start), C.int(slice.End)}, corient))
}

// ArcLength returns the length of the contour. If isClosed is negative, then
// the contour is checked to see whether the contour should be considered
// closed.  If isClosed is zero or one, then it overrides the contour's flags.
func ArcLength(contour Arr, slice Slice, isClosed int) float64 {
	return float64(C.cvArcLength(contour.arr(), C.CvSlice{C.int(slice.Start), C.int(slice.End)}, C.int(isClosed)))
}

// ContourPerimeter returns the length of a closed contour.
func ContourPerimeter(contour Arr) float64 {
	return ArcLength(contour, WHOLE_SEQ, 1)
}

// CheckContourConvexity returns true if the contour is convex.
func CheckContourConvexity(contour Arr) bool {
	return C.cvCheckContourConvexity(contour.arr()) != 0
}
