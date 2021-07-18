package pixolousAnalyze

import (
	"gocv.io/x/gocv"
)

func DetectBlur(path string, threshold float64) bool {
	img := gocv.IMRead(path, gocv.IMReadGrayScale)
	defer img.Close()
	laplacian := gocv.NewMat()
	defer laplacian.Close()

	gocv.Laplacian(img, &laplacian, gocv.MatTypeCV64F, 3, 1, 0, gocv.BorderDefault)

	dst := gocv.NewMat()
	defer dst.Close()

	dstStdDev := gocv.NewMat()
	defer dstStdDev.Close()

	gocv.MeanStdDev(laplacian, &dst, &dstStdDev)
	variance := gocv.Split(dstStdDev)[0].GetDoubleAt(0, 0)
	variance *= variance
	return variance > threshold

}
