package analysis

import (
	"fmt"
	"image"
	"strconv"

	"gocv.io/x/gocv"
)

const resizeWidth int = 8
const resizeHeight int = 8

func prepImage(img gocv.Mat) gocv.Mat {

	resizedImg := gocv.NewMat()

	gocv.Resize(img, &resizedImg, image.Pt(resizeWidth, resizeHeight), 0, 0, gocv.InterpolationDefault)
	gocv.CvtColor(resizedImg, &resizedImg, gocv.ColorBGRToGray)
	return resizedImg

}

func average2D(table2D gocv.Mat) float32 {

	var total float32
	for i := 0; i < resizeWidth; i++ {
		for j := 0; j < resizeHeight; j++ {

			total += table2D.GetFloatAt(i, j)

		}
	}
	return total / float32(resizeWidth*resizeHeight)
}
func pixelAverage(mat gocv.Mat) int {
	total := 0
	channel := gocv.Split(mat)[0]
	for i := 0; i < resizeHeight; i++ {
		for j := 0; j < resizeWidth; j++ {

			total += int(channel.GetUCharAt(0, 0))

		}
	}

	return total / (resizeWidth * resizeHeight)
}

func hashTableA(mat gocv.Mat, pixAvg int) [resizeHeight][resizeWidth]int {
	channel := gocv.Split(mat)[0]
	var hashTable [resizeHeight][resizeWidth]int

	for i := 0; i < resizeHeight; i++ {
		for j := 0; j < resizeWidth; j++ {
			if channel.GetUCharAt(i, j) > uint8(pixAvg) {
				hashTable[j][i] = 1

			} else {
				hashTable[j][i] = 0

			}
		}
	}
	return hashTable

}

func concatenation(table [resizeHeight][resizeWidth]int) string {
	var flattened string
	for i := 0; i < resizeHeight; i++ {
		for j := 0; j < resizeWidth; j++ {
			flattened += strconv.Itoa(table[i][j])

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

	avg := pixelAverage(img)

	table := hashTableA(img, avg)

	return concatenation(table)

	fmt.Println(table)
	return ""

}
