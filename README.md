# Plagiarism Detector

A Go library for detecting plagiarism using stopwords and n-grams analysis.

## Overview

This library implements a plagiarism detection algorithm that uses stopwords and n-grams to identify similar content between texts.

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
    detector, _ := plagiarism.NewDetector()
    // More examples coming soon
}
```

## License

MIT License - See LICENSE file for details

## Credits

Algorithm inspired by research in plagiarism detection using stopwords n-grams.
Original implementation by Civic Information Office (cvcio/go-plagiarism).
