package formatter

import (
	"strings"
	"testing"

	"seqrevcomp/internal/parser"
	"seqrevcomp/internal/reverse_complement"
)

func TestWrapLines(t *testing.T) {
	input := "ATCGATCGATCG"
	wrapped := wrapLines(input, 4)
	expected := "ATCG\nATCG\nATCG"
	if wrapped != expected {
		t.Errorf("expected %q, got %q", expected, wrapped)
	}
}

func TestFormatSequencesFASTA(t *testing.T) {
	seqs := []parser.Sequence{
		{Header: ">seq1", Body: "ATGCATGCATGC"},
	}
	opts := FormatOptions{LineWidth: 4, ToUpper: true}
	out := FormatSequences(seqs, opts, reverse_complement.TypeDNA)
	if !strings.Contains(out, ">seq1") {
		t.Error("missing header")
	}
	if !strings.Contains(out, "ATGC") {
		t.Error("missing sequence")
	}
}

func TestFormatSequencesRaw(t *testing.T) {
	seqs := []parser.Sequence{
		{Body: "atgc"},
	}
	opts := FormatOptions{LineWidth: 0, ToUpper: true}
	out := FormatSequences(seqs, opts, reverse_complement.TypeDNA)
	if out != "ATGC" {
		t.Errorf("expected ATGC, got %q", out)
	}
}

func TestProcessInputDNA(t *testing.T) {
	input := "ATGC"
	opts := DefaultFormatOptions()
	result := ProcessInput(input, opts)
	if result.Output != "GCAT" {
		t.Errorf("expected GCAT, got %q", result.Output)
	}
	if result.SeqType != "DNA" {
		t.Errorf("expected DNA, got %s", result.SeqType)
	}
}

func TestProcessInputRNA(t *testing.T) {
	input := "AUGC"
	opts := DefaultFormatOptions()
	result := ProcessInput(input, opts)
	if result.Output != "GCAU" {
		t.Errorf("expected GCAU, got %q", result.Output)
	}
	if result.SeqType != "RNA" {
		t.Errorf("expected RNA, got %s", result.SeqType)
	}
}

func TestProcessInputMixedTU(t *testing.T) {
	input := "ATUGC"
	opts := DefaultFormatOptions()
	result := ProcessInput(input, opts)
	if !result.HasMixedTU {
		t.Error("expected mixed TU warning")
	}
}

func TestProcessInputFASTA(t *testing.T) {
	input := ">seq1\nATGC"
	opts := DefaultFormatOptions()
	result := ProcessInput(input, opts)
	if !result.HasFASTA {
		t.Error("expected FASTA")
	}
	if !strings.Contains(result.Output, ">seq1") {
		t.Error("expected header in output")
	}
}

func TestBuildStatusMessage(t *testing.T) {
	r := FormatResult{
		SeqType:   "DNA",
		HasFASTA:  false,
		InputLen:  4,
		OutputLen: 4,
	}
	msg := BuildStatusMessage(r)
	if !strings.Contains(msg, "DNA") {
		t.Error("expected DNA in status")
	}
}
