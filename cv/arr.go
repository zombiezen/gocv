package cv

// #cgo CFLAGS: -Wno-error
// #cgo pkg-config: opencv
// #include "cv.h"
import "C"

import (
	"unsafe"
)

type Arr interface {
	Size() Size
	arr() unsafe.Pointer
}

func Copy(src, dst, mask Arr) {
	if mask == nil {
		C.cvCopy(src.arr(), dst.arr(), nil)
	} else {
		C.cvCopy(src.arr(), dst.arr(), mask.arr())
	}
}

func ConvertScale(src, dst Arr, scale, shift float64) {
	C.cvConvertScale(src.arr(), dst.arr(), C.double(scale), C.double(shift))
}
