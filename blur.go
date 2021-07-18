package pixolousAnalyze

import (
	"fmt"

	"gocv.io/x/gocv"
)

func DetectBlur(img gocv.Mat) {
	//img := gocv.IMRead(path, gocv.IMReadGrayScale)
	//defer img.Close()
	laplacian := gocv.NewMat()
	defer laplacian.Close()

	gocv.Laplacian(img, &laplacian, gocv.MatTypeCV64F, 3, 1, 0, gocv.BorderDefault)

	dst := gocv.NewMat()
	defer dst.Close()

	dstStdDev := gocv.NewMat()
	defer dstStdDev.Close()

	gocv.MeanStdDev(laplacian, &dst, &dstStdDev)
	fmt.Println("dst", dst.Cols(), dst.Rows())
	fmt.Println("dstStdDev", dstStdDev.Cols(), dstStdDev.Rows())

}
