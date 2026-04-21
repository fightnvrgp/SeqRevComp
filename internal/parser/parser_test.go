package parser

import (
	"testing"
)

func TestParseInputEmpty(t *testing.T) {
	result := ParseInput("")
	if len(result.Sequences) != 0 {
		t.Errorf("expected 0 sequences, got %d", len(result.Sequences))
	}
}

func TestParseInputRaw(t *testing.T) {
	input := "ATGC\nCGTA"
	result := ParseInput(input)
	if len(result.Sequences) != 1 {
		t.Fatalf("expected 1 sequence, got %d", len(result.Sequences))
	}
	expected := "ATGCCGTA"
	if result.Sequences[0].Body != expected {
		t.Errorf("expected %q, got %q", expected, result.Sequences[0].Body)
	}
	if result.HasFASTA {
		t.Error("expected no FASTA")
	}
}

func TestParseInputFASTA(t *testing.T) {
	input := ">seq1\nATGC\nCGTA\n>seq2\nGGGG"
	result := ParseInput(input)
	if len(result.Sequences) != 2 {
		t.Fatalf("expected 2 sequences, got %d", len(result.Sequences))
	}
	if result.Sequences[0].Header != ">seq1" {
		t.Errorf("expected header >seq1, got %q", result.Sequences[0].Header)
	}
	if result.Sequences[0].Body != "ATGCCGTA" {
		t.Errorf("expected body ATGCCGTA, got %q", result.Sequences[0].Body)
	}
	if result.Sequences[1].Body != "GGGG" {
		t.Errorf("expected body GGGG, got %q", result.Sequences[1].Body)
	}
	if !result.HasFASTA {
		t.Error("expected FASTA")
	}
}

func TestParseInputWithSpacesAndNumbers(t *testing.T) {
	input := "1 ATGC\n2 CGTA"
	result := ParseInput(input)
	if len(result.Sequences) != 1 {
		t.Fatalf("expected 1 sequence, got %d", len(result.Sequences))
	}
	expected := "ATGCCGTA"
	if result.Sequences[0].Body != expected {
		t.Errorf("expected %q, got %q", expected, result.Sequences[0].Body)
	}
}

func TestParseInputWithTabs(t *testing.T) {
	input := "ATGC\tCGTA"
	result := ParseInput(input)
	expected := "ATGCCGTA"
	if result.Sequences[0].Body != expected {
		t.Errorf("expected %q, got %q", expected, result.Sequences[0].Body)
	}
}

func TestCleanSequence(t *testing.T) {
	input := "A T 1 G C\n5"
	expected := "ATGC"
	got := cleanSequence(input)
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}

func TestSplitRawIntoSequences(t *testing.T) {
	input := "ATGC\n\nCGTA\n\nGGGG"
	seqs := SplitRawIntoSequences(input)
	if len(seqs) != 3 {
		t.Fatalf("expected 3 sequences, got %d", len(seqs))
	}
	if seqs[0].Body != "ATGC" {
		t.Errorf("expected ATGC, got %q", seqs[0].Body)
	}
	if seqs[1].Body != "CGTA" {
		t.Errorf("expected CGTA, got %q", seqs[1].Body)
	}
}
