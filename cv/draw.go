package cv

// #cgo CFLAGS: -Wno-error
// #cgo pkg-config: opencv
// #include "cv.h"
import "C"

const AA = C.CV_AA

func PolyLine(img Arr, points [][]Point, closed bool, color Scalar, thickness, lineType, shift int) {
	if len(points) == 0 {
		return
	}

	var cc C.int
	if closed {
		cc = 1
	} else {
		cc = 0
	}

	cvpoints := make([][]C.CvPoint, len(points))
	for i := range points {
		cvpoints[i] = make([]C.CvPoint, len(points[i]))
		for j := range points[i] {
			cvpoints[i][j] = C.CvPoint{C.int(points[i][j].X), C.int(points[i][j].Y)}
		}
	}

	pts := make([]*C.CvPoint, len(points))
	npts := make([]C.int, len(points))
	for i := range points {
		// XXX: Is it safe to pass our point struct as CvPoint?
		pts[i] = &cvpoints[i][0]
		npts[i] = C.int(len(points[i]))
	}

	C.cvPolyLine(img.arr(), &pts[0], &npts[0], C.int(len(points)), cc, color.cvScalar(), C.int(thickness), C.int(lineType), C.int(shift))
}
