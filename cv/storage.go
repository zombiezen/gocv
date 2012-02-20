package cv

// #cgo CFLAGS: -I/usr/include/opencv -Wno-error
// #cgo LDFLAGS: -lcv -lcxcore
// #include "cv.h"
import "C"

// MemStorage is an OpenCV memory pool.
type MemStorage struct {
	s *C.CvMemStorage
}

// NewMemStorage creates new memory storage. A blockSize of zero uses the
// default block size.
func NewMemStorage(blockSize int) MemStorage {
	return MemStorage{C.cvCreateMemStorage(C.int(blockSize))}
}

func (s MemStorage) Release() {
	C.cvReleaseMemStorage(&s.s)
}
