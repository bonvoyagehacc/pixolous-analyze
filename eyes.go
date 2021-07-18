package pixolousAnalyze

import (
	"fmt"

	"gocv.io/x/gocv"
)

func DetectOpenEyes(path string) {
	face_cascade := gocv.NewCascadeClassifier()
	defer face_cascade.Close()

	if !face_cascade.Load("haarcascade_frontalface_default.xml") {
		fmt.Printf("Error reading cascade file: %v\n", "haarcascade_frontalface_default.xml")
		return
	}

	eye_cascade := gocv.NewCascadeClassifier()
	defer eye_cascade.Close()

	if !eye_cascade.Load("haarcascade_eye_tree_eyeglasses.xml") {
		fmt.Printf("Error reading cascade file: %v\n", "haarcascade_eye_tree_eyeglasses.xml")
		return
	}

}
