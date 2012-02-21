package cv

// #include "cv.h"
import "C"

import (
	"unsafe"
)

// Arr is the interface that wraps any image-like surface. IplImage and Seq
// implement this interface.
type Arr interface {
	Size() Size
	arr() unsafe.Pointer
}

// Copy copies elements from src to dst. If mask is not nil, then only elements
// that have a non-zero mask element will be copied.
func Copy(src, dst, mask Arr) {
	if mask == nil {
		C.cvCopy(src.arr(), dst.arr(), nil)
	} else {
		C.cvCopy(src.arr(), dst.arr(), mask.arr())
	}
}

// ConvertScale converts from src to dst.  Each element is multiplied by scale
// then increased by shift.
func ConvertScale(src, dst Arr, scale, shift float64) {
	C.cvConvertScale(src.arr(), dst.arr(), C.double(scale), C.double(shift))
}
