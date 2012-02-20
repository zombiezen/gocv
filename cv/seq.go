package cv

// #cgo CFLAGS: -I/usr/include/opencv -Wno-error
// #cgo LDFLAGS: -lcv -lcxcore
// #include "cv.h"
import "C"

import (
	"unsafe"
)

type Seq struct {
	seq *C.CvSeq
}

// Len returns the number of items in the sequence.
func (s Seq) Len() int {
	return int(s.seq.total)
}

func (s Seq) IsZero() bool {
	return s.seq == nil
}

// Prev returns the previous linked sequence.
func (s Seq) Prev() Seq {
	return Seq{(*C.CvSeq)(s.seq.h_prev)}
}

// Next returns the next linked sequence.
func (s Seq) Next() Seq {
	return Seq{(*C.CvSeq)(s.seq.h_next)}
}

// At returns the element at i.
func (s Seq) At(i int) unsafe.Pointer {
	if i < 0 || i >= int(s.seq.total) {
		panic("Seq index out of bounds")
	}
	return unsafe.Pointer(C.cvGetSeqElem(s.seq, C.int(i)))
}

// PointAt returns the element at i converted to a point.
func (s Seq) PointAt(i int) Point {
	pt := (*C.CvPoint)(s.At(i))
	return Point{int(pt.x), int(pt.y)}
}

func (s Seq) arr() unsafe.Pointer {
	return unsafe.Pointer(s.seq)
}

func (s Seq) Size() Size {
	sz := C.cvGetSize(s.arr())
	return Size{int(sz.width), int(sz.height)}
}

type Slice struct {
	Start, End int
}

var WHOLE_SEQ = Slice{0, C.CV_WHOLE_SEQ_END_INDEX}
