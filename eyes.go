package pixolousAnalyze

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

func DetectOpenEyes(path string) float32 {
	face_cascade := gocv.NewCascadeClassifier()
	defer face_cascade.Close()

	if !face_cascade.Load("haarcascade_frontalface_default.xml") {
		fmt.Printf("Error reading cascade file: %v\n", "haarcascade_frontalface_default.xml")
		return -1
	}

	eye_cascade := gocv.NewCascadeClassifier()
	defer eye_cascade.Close()

	if !eye_cascade.Load("haarcascade_eye_tree_eyeglasses.xml") {
		fmt.Printf("Error reading cascade file: %v\n", "haarcascade_eye_tree_eyeglasses.xml")
		return -1
	}
	img := gocv.IMRead(path, gocv.IMReadGrayScale)
	defer img.Close()
	if img.Empty() {
		fmt.Println(path, "empty")
	}
	faces := face_cascade.DetectMultiScaleWithParams(img, 1.3, 5, 0, image.Pt(200, 200), image.Pt(img.Cols(), img.Rows()))
	totalEyes := 0
	for _, r := range faces {
		roi_face := img.Region(r)
		eyes := eye_cascade.DetectMultiScaleWithParams(roi_face, 1.3, 5, 0, image.Pt(50, 50), image.Pt(roi_face.Cols(), roi_face.Rows()))
		totalEyes += len(eyes)

	}
	if len(faces) == 0 {
		return 0
	}
	return float32(totalEyes) / float32(len(faces))

}
