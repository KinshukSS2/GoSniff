package plagiarism

import "fmt"

// DeepEquaility something like Jaccard coefficient but with strict position.
// Instead of intersection / union we use position / union == 1.0
func (p *Detector) DeepEquaility(source, target *[][]string) [][]Set {
	// Copy Temp Slices, I > J
	tempI := *source
	tempJ := *target

	// initilize source sets and set scores to 0
	setI := make([]Set, len(tempI))
	for i := range tempI {
		setI[i] = Set{NGram: tempI[i], Score: 0}
	}

	// initilize target sets and set scores to 0
	setJ := make([]Set, len(tempJ))
	for j := range tempJ {
		setJ[j] = Set{NGram: tempJ[j], Score: 0}
	}

	// find equals for I/J and set score to 1
	for i := range setI {
		for j := range setJ {
			if p.Equal(setI[i].NGram, setJ[j].NGram) {
				setI[i].Score = 1
				setJ[j].Score = 1
			}
		}
	}

	// return the sets
	return [][]Set{setI, setJ}
}

// Detect will read values directly from the detector interface bypassing
// GetStopWords and Tokenize methods assuming that you already provided a
// a list of stopwords for each string (source, target). Will return an
// error on failure.
func (p *Detector) Detect() error {
	// check if any of source or target stopwords list is an empty string array and return an error
	if (len(p.SourceStopWords) < 1 || len(p.TargetStopWords) < 1) && (p.SourceText == "" || p.TargetText == "") {
		return fmt.Errorf("you should at least define source and target texts")
	}

	if len(p.SourceStopWords) > 0 && len(p.TargetStopWords) > 0 {
		return p.DetectWithStopWords(p.SourceStopWords, p.TargetStopWords)
	}

	// check if any of source or target text is an empty string and return an error
	if p.SourceText != "" && p.TargetText != "" {
		return p.DetectWithStrings(p.SourceText, p.TargetText)
	}

	return fmt.Errorf("empty strings cannot continue")
}

// DetectWithStrings returns an error on failure, otherwise will invoke
// DetectWithStopWords method.
func (p *Detector) DetectWithStrings(source, target string) error {
	// check if any of source or target text is an empty string and return an error
	if source == "" || target == "" {
		return fmt.Errorf("both, source and target text cannot be empty")
	}

	// assign detector values
	p.SourceText = source
	p.TargetText = target

	// tokenize sting, filter stopwords and return DetectWithStopWords method
	return p.DetectWithStopWords(
		p.GetStopWords(p.Tokenize(p.SourceText)),
		p.GetStopWords(p.Tokenize(p.TargetText)),
	)
}

// DetectWithStopWords returns an error on failure, otherwise will set Score, Similar and Total
// values to the detector interface.
func (p *Detector) DetectWithStopWords(source, target []string) error {
	// check if any of source or target stopwords list is an empty string array and return an error
	if len(source) < 1 || len(target) < 1 {
		return fmt.Errorf("both, source and target stopwords list cannot be empty")
	}

	// assign detector values
	p.SourceStopWords = source
	p.TargetStopWords = target

	// get the n-grams and assign detector values
	p.SourceNGrams = p.GetNGrams(p.SourceStopWords)
	p.TargetNGrams = p.GetNGrams(p.TargetStopWords)

	// test n-grams equality
	equality := p.DeepEquaility(&p.SourceNGrams, &p.TargetNGrams)

	// sum source similarity score
	for i := range equality[0] {
		p.Similar += equality[0][i].Score
	}

	// sum target similarity score
	for j := range equality[1] {
		p.Similar += equality[1][j].Score
	}

	// sum totals
	p.Total = len(equality[0]) + len(equality[1])

	// calculate probability
	p.Score = float64(p.Similar) / float64(p.Total)

	return nil
}
