package cv

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

// Release deallocates all of the memory in the pool.
func (s MemStorage) Release() {
	C.cvReleaseMemStorage(&s.s)
}
