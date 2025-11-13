package plagiarismpackage plagiarismdetector


const (
	// N default n-gram size
	N = 8
	// LANG default language
	LANG = "en"
)

// Set of n-grams and scores
type Set struct {
	NGram []string
	Score int
}

// Detector struct holds the configuration and results of plagiarism detection
type Detector struct {
	N               int
	Lang            string
	StopWords       []string
	SourceText      string
	TargetText      string
	SourceStopWords []string
	TargetStopWords []string
	SourceNGrams    [][]string
	TargetNGrams    [][]string
	Score           float64
	Similar         int
	Total           int
}
