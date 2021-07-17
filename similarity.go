package analysis

import (
	"image"
	"strconv"

	"gocv.io/x/gocv"
)

var resizeWidth int = 8
var resizeHeight int = 8

func prepImage(img gocv.Mat) gocv.Mat {

	resizedImg := gocv.NewMat()

	gocv.Resize(img, &resizedImg, image.Pt(resizeWidth, resizeHeight), 0, 0, gocv.InterpolationDefault)
	gocv.CvtColor(resizedImg, &resizedImg, gocv.ColorBGRToGray)
	return resizedImg

}

func average2D(table2D gocv.Mat) float32 {

	var total int32
	for i := 0; i < resizeWidth; i++ {
		for j := 0; j < resizeHeight; j++ {
			total += table2D.GetIntAt(i, j)
		}
	}
	return float32(total) / float32(resizeWidth*resizeHeight)
}
func hashTableA(table2D gocv.Mat, averageHash float32) gocv.Mat {
	for i := 0; i < resizeWidth; i++ {
		for j := 0; j < resizeHeight; j++ {
			if table2D.GetFloatAt(i, j) > averageHash {
				table2D.SetIntAt(i, j, 1)

			} else {
				table2D.SetIntAt(i, j, 0)
			}
		}
	}
	return table2D

}

func concatenation(table2D gocv.Mat) string {
	var flattened string
	for i := 0; i < resizeWidth; i++ {
		for j := 0; j < resizeHeight; j++ {
			flattened += strconv.Itoa(int(table2D.GetIntAt(i, j)))

		}
	}
	result, _ := strconv.ParseInt(flattened, 2, 64)
	return reverseString(strconv.FormatInt(result, 16))

}
func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)

}

func AHash(image gocv.Mat) string {

	img := prepImage(image)

	//avg := average2D(img)

	//hashedImg := hashTableA(img, avg)

	//return concatenation(hashedImg)

	return strconv.Itoa(img.Cols())

}
