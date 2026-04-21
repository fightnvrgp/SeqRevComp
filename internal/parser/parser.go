// Package parser handles input parsing, cleaning, and FASTA extraction.
package parser

import (
	"bufio"
	"strings"
	"unicode"
)

// Sequence represents a single sequence with optional FASTA header.
type Sequence struct {
	Header string
	Body   string
}

// ParseResult holds all parsed sequences and any warnings.
type ParseResult struct {
	Sequences    []Sequence
	Warnings     []string
	HasFASTA     bool
	OriginalType string // "fasta" or "raw"
}

// ParseInput cleans and parses raw user input.
func ParseInput(input string) ParseResult {
	result := ParseResult{}
	if strings.TrimSpace(input) == "" {
		result.Warnings = append(result.Warnings, "输入为空")
		return result
	}

	lines := strings.Split(input, "\n")
	var sequences []Sequence
	var current *Sequence
	hasFASTA := false

	flushCurrent := func() {
		if current != nil {
			current.Body = cleanSequence(current.Body)
			if current.Body != "" || current.Header != "" {
				sequences = append(sequences, *current)
			}
			current = nil
		}
	}

	for _, rawLine := range lines {
		line := strings.TrimRight(rawLine, "\r")
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, ">") {
			hasFASTA = true
			flushCurrent()
			current = &Sequence{
				Header: strings.TrimSpace(trimmed),
				Body:   "",
			}
		} else if current != nil {
			current.Body += line
		} else if !hasFASTA {
			current = &Sequence{Header: "", Body: line}
		}
	}
	flushCurrent()

	// Remove empty bodies
	var nonEmpty []Sequence
	for _, s := range sequences {
		if s.Body != "" {
			nonEmpty = append(nonEmpty, s)
		}
	}

	result.HasFASTA = hasFASTA
	result.OriginalType = "raw"
	if hasFASTA {
		result.OriginalType = "fasta"
	}
	result.Sequences = nonEmpty

	if len(nonEmpty) == 0 {
		result.Warnings = append(result.Warnings, "未能从输入中解析出有效序列")
	}

	return result
}

// cleanSequence removes spaces, tabs, digits, and other non-sequence characters,
// keeping only valid nucleotide letters.
func cleanSequence(s string) string {
	var b strings.Builder
	for _, ch := range s {
		if unicode.IsSpace(ch) || unicode.IsDigit(ch) {
			continue
		}
		b.WriteRune(ch)
	}
	return b.String()
}

// SplitRawIntoSequences attempts to split raw text into multiple sequences
// by looking for blank lines or header-like patterns.
func SplitRawIntoSequences(input string) []Sequence {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var groups []Sequence
	var current strings.Builder
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			if current.Len() > 0 {
				groups = append(groups, Sequence{Body: cleanSequence(current.String())})
				current.Reset()
			}
		} else {
			current.WriteString(line)
		}
	}
	if current.Len() > 0 {
		groups = append(groups, Sequence{Body: cleanSequence(current.String())})
	}
	return groups
}
