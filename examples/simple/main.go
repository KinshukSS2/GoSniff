package main

import (
	"fmt"
	plagiarism "github.com/KinshukSS2/plag-checker"
)

func main() {
	source := "Plagiarism detection is an important task in academia. This algorithm uses stopwords to find similar content between two texts. The method is based on n-gram analysis of function words."
	target := "This algorithm uses stopwords to find similar content between texts. The method is based on n-gram analysis of function words and can detect plagiarism effectively."

	detector, err := plagiarism.NewDetector(plagiarism.SetN(6))
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
