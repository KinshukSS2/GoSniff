# Plagiarism Detector

A Go library for detecting plagiarism using stopwords and n-grams analysis.

## Overview

This library implements a plagiarism detection algorithm that uses stopwords and n-grams to identify similar content between texts. The algorithm analyzes the position and frequency of function words (stopwords) to calculate a similarity score.

### Similarity Network Visualization

The algorithm creates relationships between similar documents, allowing you to visualize networks of similar content and identify potential plagiarism chains.

![Similarity Network](assets/similarity-network.png)

*The similarity network shows how documents are connected based on their similarity scores. Each node represents a document, and edges show the similarity percentage between documents.*

### The Chain of Misinformation

This algorithm can track how content spreads and evolves through reproduction, paraphrasing, and falsification - helping identify the chain of misinformation.

![The Chain of Misinformation](assets/chain-of-misinformation.png)

*Tracking how original content transforms through reproduction, paraphrasing, new article creation, falsification, patchwork plagiarism, and ultimately misinformation.*

## Use Cases

- **Academic Integrity**: Detect plagiarism in student papers and research articles
- **Content Monitoring**: Track content copying and unauthorized republishing
- **Misinformation Tracking**: Identify how false information spreads and mutates
- **Copyright Protection**: Find unauthorized use of copyrighted content
- **SEO Analysis**: Detect duplicate content across websites

## Features

- Fast plagiarism detection using stopwords n-gram analysis
- Configurable n-gram size (default: 8)
- English stopwords support
- Simple and intuitive API
- Detects paraphrasing and patchwork plagiarism

## Installation

```bash
go get github.com/KinshukSS2/plag-checker
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/KinshukSS2/plag-checker"
)

func main() {
    source := "This is a test document for plagiarism detection"
    target := "This is a test document for detection"

    detector, _ := plagiarism.NewDetector()
    err := detector.DetectWithStrings(source, target)
    
    if err != nil {
        panic(err)
    }

    fmt.Printf("Similarity Score: %.2f\n", detector.Score)
    fmt.Printf("Similar n-grams: %d\n", detector.Similar)
    fmt.Printf("Total n-grams: %d\n", detector.Total)
}
```

## API Documentation

### Creating a Detector

```go
// Create with default settings (n=8, language="en")
detector, err := plagiarism.NewDetector()

// Create with custom n-gram size
detector, err := plagiarism.NewDetector(plagiarism.SetN(12))

// Create with custom stopwords
customStopwords := []string{"the", "a", "is", "are"}
detector, err := plagiarism.NewDetector(plagiarism.SetStopWords(customStopwords))
```

### Detection Methods

```go
// Detect using text strings
err := detector.DetectWithStrings(sourceText, targetText)

// Detect using pre-filtered stopwords
err := detector.DetectWithStopWords(sourceStopwords, targetStopwords)
```

### Results

After detection, access results through the detector:
- `detector.Score` - Similarity score (0.0 to 1.0)
- `detector.Similar` - Number of matching n-grams
- `detector.Total` - Total number of n-grams compared

## How It Works

1. **Tokenization**: Text is split into individual words (tokens)
2. **Stopword Filtering**: Only stopwords (function words like "the", "is", "and") are kept
3. **N-gram Generation**: Stopwords are converted into n-grams (sequences of N words)
4. **Similarity Calculation**: N-grams are compared for positional matches
5. **Score**: Ratio of matching n-grams to total n-grams

## Testing

Run tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

## Examples

See the `examples/` directory for complete examples.

## License

MIT License - See LICENSE file for details

## Credits

Algorithm inspired by research in plagiarism detection using stopwords n-grams.

By:Kinshuk 