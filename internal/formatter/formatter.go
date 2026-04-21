// Package formatter handles output formatting.
package formatter

import (
	"fmt"
	"strings"

	"seqrevcomp/internal/parser"
	"seqrevcomp/internal/reverse_complement"
)

// FormatOptions controls how output is formatted.
type FormatOptions struct {
	LineWidth int  // 0 means no wrapping
	ToUpper   bool // default true
}

// DefaultFormatOptions returns default formatting options.
func DefaultFormatOptions() FormatOptions {
	return FormatOptions{
		LineWidth: 60,
		ToUpper:   true,
	}
}

// FormatSequences formats sequences for output.
// If FASTA input, preserves headers and wraps lines.
// If raw input, outputs plain sequence text.
func FormatSequences(seqs []parser.Sequence, opts FormatOptions, seqType reverse_complement.SequenceType) string {
	if len(seqs) == 0 {
		return ""
	}

	var parts []string
	hasAnyHeader := false
	for _, s := range seqs {
		if s.Header != "" {
			hasAnyHeader = true
			break
		}
	}

	for i, s := range seqs {
		body := s.Body
		if opts.ToUpper {
			body = strings.ToUpper(body)
		}

		if hasAnyHeader && s.Header != "" {
			parts = append(parts, s.Header)
		}

		if opts.LineWidth > 0 {
			parts = append(parts, wrapLines(body, opts.LineWidth))
		} else {
			parts = append(parts, body)
		}

		// Add blank line between FASTA entries
		if hasAnyHeader && i < len(seqs)-1 {
			parts = append(parts, "")
		}
	}

	return strings.Join(parts, "\n")
}

// FormatResult holds the complete processing result for display.
type FormatResult struct {
	Output      string
	SeqType     string
	HasFASTA    bool
	HasMixedTU  bool
	InputLen    int
	OutputLen   int
	Warnings    []string
}

// ProcessInput performs the full pipeline: parse, analyze, reverse complement, format.
func ProcessInput(input string, opts FormatOptions) FormatResult {
	result := FormatResult{}

	// Parse
	parseResult := parser.ParseInput(input)
	result.HasFASTA = parseResult.HasFASTA
	result.Warnings = parseResult.Warnings

	if len(parseResult.Sequences) == 0 {
		result.Warnings = append(result.Warnings, "未找到有效序列")
		return result
	}

	// Determine overall sequence type from concatenated bodies
	var allBodies strings.Builder
	for _, s := range parseResult.Sequences {
		allBodies.WriteString(s.Body)
	}
	concat := allBodies.String()

	// Detect type
	seqTypeStr := "DNA"
	if strings.Contains(strings.ToUpper(concat), "U") && !strings.Contains(strings.ToUpper(concat), "T") {
		seqTypeStr = "RNA"
	}
	result.SeqType = seqTypeStr

	if strings.Contains(strings.ToUpper(concat), "T") && strings.Contains(strings.ToUpper(concat), "U") {
		result.HasMixedTU = true
		result.Warnings = append(result.Warnings, "同时检测到 T 和 U，默认按 DNA 规则处理（U 视为 T 处理）")
	}

	result.InputLen = len(concat)

	// Reverse complement each sequence
	var rcSeqs []parser.Sequence
	seqType := reverse_complement.TypeDNA
	if seqTypeStr == "RNA" {
		seqType = reverse_complement.TypeRNA
	}

	for _, s := range parseResult.Sequences {
		rcBody := reverse_complement.ReverseComplement(s.Body, seqType)
		if opts.ToUpper {
			rcBody = strings.ToUpper(rcBody)
		}
		rcSeqs = append(rcSeqs, parser.Sequence{
			Header: s.Header,
			Body:   rcBody,
		})
	}

	result.Output = FormatSequences(rcSeqs, opts, seqType)
	result.OutputLen = len(strings.ReplaceAll(result.Output, "\n", ""))
	// Subtract headers from output length
	if result.HasFASTA {
		total := 0
		for _, s := range rcSeqs {
			total += len(s.Body)
		}
		result.OutputLen = total
	}

	return result
}

func wrapLines(s string, width int) string {
	if width <= 0 || len(s) <= width {
		return s
	}
	var parts []string
	for i := 0; i < len(s); i += width {
		end := i + width
		if end > len(s) {
			end = len(s)
		}
		parts = append(parts, s[i:end])
	}
	return strings.Join(parts, "\n")
}

// BuildStatusMessage creates a human-readable status string.
func BuildStatusMessage(r FormatResult) string {
	var parts []string
	parts = append(parts, fmt.Sprintf("序列类型: %s", r.SeqType))
	if r.HasFASTA {
		parts = append(parts, "格式: FASTA")
	} else {
		parts = append(parts, "格式: 裸序列")
	}
	if r.HasMixedTU {
		parts = append(parts, "警告: 同时含 T 和 U")
	}
	parts = append(parts, fmt.Sprintf("输入长度: %d", r.InputLen))
	parts = append(parts, fmt.Sprintf("输出长度: %d", r.OutputLen))
	if len(r.Warnings) > 0 {
		parts = append(parts, fmt.Sprintf("提示: %s", strings.Join(r.Warnings, "; ")))
	}
	return strings.Join(parts, " | ")
}
