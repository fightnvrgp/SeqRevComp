// Package validation provides sequence validation and type detection.
package validation

import (
	"strings"
)

// DetectionResult holds the results of sequence analysis.
type DetectionResult struct {
	SeqType       string // "DNA" or "RNA"
	HasFASTA      bool
	HasMixedTU    bool
	InvalidChars  []rune
	CleanedLength int
}

// DetectSequenceType determines if the input is more like DNA or RNA.
// Rules:
//   - If contains U but no T → RNA
//   - If contains T but no U → DNA
//   - If contains both T and U → mixed (warn, default to DNA)
//   - If neither → DNA (default)
func DetectSequenceType(seq string) string {
	upper := strings.ToUpper(seq)
	hasT := strings.Contains(upper, "T")
	hasU := strings.Contains(upper, "U")

	if hasU && !hasT {
		return "RNA"
	}
	if hasT && !hasU {
		return "DNA"
	}
	return "DNA"
}

// DetectMixedTU reports whether the sequence contains both T and U.
func DetectMixedTU(seq string) bool {
	upper := strings.ToUpper(seq)
	return strings.Contains(upper, "T") && strings.Contains(upper, "U")
}

// ValidNucleotides defines all acceptable characters (including IUPAC).
const ValidNucleotides = "ATCGRYSWKMBDHVNatcgryswkmbdhvn"

// FindInvalidChars returns any characters that are not valid nucleotides.
func FindInvalidChars(seq string) []rune {
	invalidMap := make(map[rune]bool)
	var result []rune
	for _, ch := range seq {
		if !strings.ContainsRune(ValidNucleotides, ch) {
			if !invalidMap[ch] {
				invalidMap[ch] = true
				result = append(result, ch)
			}
		}
	}
	return result
}

// Analyze performs full analysis on a cleaned sequence string.
func Analyze(cleanedSeq string) DetectionResult {
	return DetectionResult{
		SeqType:       DetectSequenceType(cleanedSeq),
		HasMixedTU:    DetectMixedTU(cleanedSeq),
		InvalidChars:  FindInvalidChars(cleanedSeq),
		CleanedLength: len(cleanedSeq),
	}
}
