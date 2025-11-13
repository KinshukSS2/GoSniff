package plagiarism

import (
	"bufio"
	"strings"
)

// Tokenize method will split the input string using bufio.Scanner into word tokens
// in order to filter out the unnecessary words. You can always use your own
// tokenizer and provide only the stopwords by using the SetStopWords option instead.
func (p *Detector) Tokenize(input string) []string {
	var output []string
	scanner := bufio.NewScanner(strings.NewReader(strings.ToLower(input)))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	return output
}

// GetStopWords returns the stopwords list for a given token list.
func (p *Detector) GetStopWords(input []string) []string {
	var output []string
	for _, token := range input {
		if p.IsStopWord(token) {
			output = append(output, token)
		}
	}
	return output
}

// IsStopWord will check if a given token is in the stopwords list.
func (p *Detector) IsStopWord(token string) bool {
	for _, stopWord := range p.StopWords {
		if stopWord == token {
			return true
		}
	}
	return false
}

// GetNGrams returns the 2D array representation of the stopword list.
func (p *Detector) GetNGrams(tokens []string) [][]string {
	// implement ngram 2D list
	grams := make([][]string, 0)

	// calculate offset and max for N, length
	offset := int(p.N / 2)
	max := len(tokens)

	// loop through tokens and append to ngram list
	for i := range tokens {
		if i < offset || i+p.N-offset > max {
			continue
		}
		grams = append(grams, tokens[i-offset:i+p.N-offset])
	}

	// return the n-gram list
	return grams
}

// Equal will test if Slices are Equal (NxN).
func (p *Detector) Equal(source, target []string) bool {
	for i := range source {
		if source[i] != target[i] {
			return false
		}
	}
	return true
}
