package plagiarism

import "fmt"

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

// NewDetector implements the detector interface. Will return a new detector or an error
// if any of the optional arguments fails.
func NewDetector(options ...Option) (*Detector, error) {
	// implement a new detector interface with defaults
	detector := &Detector{N: N, Lang: LANG, StopWords: StopWords[LANG].([]string)}

	// iterrate over options, apply or return an error on failure
	for _, opt := range options {
		if err := opt(detector); err != nil {
			return nil, err
		}
	}
	// retrun the detecor
	return detector, nil
}

// Option applies detector options and returns an error on failure.
type Option func(*Detector) (err error)

// SetN option sets the detector's n-gram size and must be a positive integer larger than 0,
// otherwise an error will be returned. The default n-gram size is 8.
func SetN(n int) Option {
	return func(p *Detector) (err error) {
		// check if n-gram size is larger than 0
		if n > 0 {
			p.N = n
			return
		}
		// otherwise return an error
		return fmt.Errorf("illegal n-gram size %d, must be a positive integer larger than 0 (tip consider using values within range 7-16)", n)
	}
}

// SetLang option sets the detector's language as well as the stopwords for the specified language.
// Use ISO 639-1 formatted language codes (https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes).
// Refer to stopwords.go for all supported languages. If the specified language doesn't exists or
// has no stopwords (empty []string), will return an error. If you want to use a custom
// language or a custom stopwords list use SetStopWords option instead.
func SetLang(lang string) Option {
	return func(p *Detector) (err error) {
		// check if language exists and has stopwords
		if val, ok := StopWords[lang]; ok && val != nil {
			p.Lang = lang
			p.StopWords = val.([]string)
			return
		}
		// otherwise return an error
		return fmt.Errorf("language %s not found or not supported yet (tip consider using a custom stopwords list with SetStopWords option", lang)
	}
}

// SetStopWords option will set a custom language and stopword list.
func SetStopWords(stopWords []string) Option {
	return func(p *Detector) (err error) {
		// check if stopwords list is not empty, otherwise return an error
		if len(stopWords) < 1 {
			return fmt.Errorf("stopwords list cannot be empty")
		}
		p.Lang = "custom"
		p.StopWords = stopWords
		return
	}
}
