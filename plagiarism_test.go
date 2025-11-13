package plagiarism

import "testing"

func TestNewDetector(t *testing.T) {
	detector, err := NewDetector()
	if err != nil {
		t.Errorf("NewDetector() error = %v", err)
		return
	}
	if detector.N != N {
		t.Errorf("NewDetector() N = %v, want %v", detector.N, N)
	}
	if detector.Lang != LANG {
		t.Errorf("NewDetector() Lang = %v, want %v", detector.Lang, LANG)
	}
}

func TestDetectWithStrings(t *testing.T) {
	source := "This is a simple test to check plagiarism detection"
	target := "This is a test to check plagiarism detection"
	
	detector, _ := NewDetector()
	err := detector.DetectWithStrings(source, target)
	
	if err != nil {
		t.Errorf("DetectWithStrings() error = %v", err)
		return
	}
	
	if detector.Score == 0 {
		t.Error("DetectWithStrings() Score should not be 0")
	}
}

func TestTokenize(t *testing.T) {
	detector, _ := NewDetector()
	input := "Hello World Test"
	tokens := detector.Tokenize(input)
	
	if len(tokens) != 3 {
		t.Errorf("Tokenize() returned %d tokens, want 3", len(tokens))
	}
}
