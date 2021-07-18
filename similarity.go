package pixolousAnalyze

import (
	"fmt"
	"image"
	"strconv"
	"strings"

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
func GetSimilarGrouped(pathToHash map[string]string) [][]string {
	groups := make([][]string, 0)
	visited := make([]string, 0)
	for h := range pathToHash {
		if contains(visited, h) {
			continue

		}
		group := make([]string, 0)

		group = append(group, h)

		visited = append(visited, h)
		for i := range pathToHash {
			if contains(visited, i) {
				continue
			}
			similarity := hashSimilarity(pathToHash[h], pathToHash[i])
			if similarity > 70 {

				group = append(group, i)
				visited = append(visited, i)

			}

		}
		groups = append(groups, group)

	}
	return groups

}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
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

	return parseBinToHex(flattened)

}
func parseBinToHex(s string) string {
	ui, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		return "error"
	}

	return fmt.Sprintf("%x", ui)
}

// func reverseString(str string) string {
// 	byte_str := []rune(str)
// 	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
// 		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
// 	}
// 	return string(byte_str)

// }
func hashSimilarity(hash1 string, hash2 string) float64 {

	result1, _ := strconv.ParseInt(hash1, 16, 64)
	result2, _ := strconv.ParseInt(hash2, 16, 64)

	size := len(hash1)
	if size < len(hash2) {
		size = len(hash2)
	}
	size *= 4

	similitude := strconv.FormatInt(result1^result2, 2)
	percentage := float64(strings.Count(similitude, "0")+size-len(similitude)) / float64(size)
	return float64(percentage * 100)

}

func AHash(image gocv.Mat) string {

	img := prepImage(image)

	avg := pixelAverage(img)

	table := hashTableA(img, avg)

	hash := concatenation(table)
	return hash

}
