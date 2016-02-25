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
	do(func() {
		if mask == nil {
			C.cvCopy(src.arr(), dst.arr(), nil)
		} else {
			C.cvCopy(src.arr(), dst.arr(), mask.arr())
		}
	})
}

// ConvertScale converts from src to dst.  Each element is multiplied by scale
// then increased by shift.
func ConvertScale(src, dst Arr, scale, shift float64) {
	do(func() {
		C.cvConvertScale(src.arr(), dst.arr(), C.double(scale), C.double(shift))
	})
}

// SetData copies bytes into the array.  Usually an IplImage's data will be
// packed in interleaved BGR order.  widthStep is the number of bytes per row.
func SetData(arr Arr, data []byte, widthStep int) {
	do(func() {
		C.cvSetData(arr.arr(), unsafe.Pointer(&data[0]), C.int(widthStep))
	})
}
