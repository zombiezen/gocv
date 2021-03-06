package cv

// #include "cv.h"
import "C"

const AA = C.CV_AA

// PolyLine draws one or more lines onto img.
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

	cvpoints := make([][]C.CvPoint, 0, len(points))
	pts := make([]*C.CvPoint, 0, len(points))
	npts := make([]C.int, 0, len(points))

	for i := range points {

		if len(points[i]) == 0 {
			continue
		}

		cvpoints[i] = make([]C.CvPoint, len(points[i]))

		for j := range points[i] {
			cvpoints[i][j] = C.CvPoint{C.int(points[i][j].X), C.int(points[i][j].Y)}
		}

		pts = append(pts, &cvpoints[i][0])
		npts = append(npts, C.int(len(points[i])))

	}

	if len(pts) == 0 {
		return
	}

	do(func() {
		C.cvPolyLine(img.arr(), &pts[0], &npts[0], C.int(len(points)), cc, color.cvScalar(), C.int(thickness), C.int(lineType), C.int(shift))
	})
}
