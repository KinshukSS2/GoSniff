package mainpackage simple


import (
	"fmt"
	plagiarism "github.com/KinshukSS2/plag-checker"
)

func main() {
	source := "The quick brown fox jumps over the lazy dog and runs away"
	target := "The quick brown fox jumps over the lazy dog"

	detector, err := plagiarism.NewDetector()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = detector.DetectWithStrings(source, target)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Similarity Score: %.2f\n", detector.Score)
	fmt.Printf("Similar n-grams: %d\n", detector.Similar)
	fmt.Printf("Total n-grams: %d\n", detector.Total)
}
