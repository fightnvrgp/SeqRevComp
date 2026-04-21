package reverse_complement

import (
	"testing"
)

func TestComplementDNA(t *testing.T) {
	tests := []struct {
		input    byte
		expected byte
	}{
		{'A', 'T'}, {'T', 'A'}, {'C', 'G'}, {'G', 'C'},
		{'R', 'Y'}, {'Y', 'R'}, {'S', 'S'}, {'W', 'W'},
		{'K', 'M'}, {'M', 'K'}, {'B', 'V'}, {'V', 'B'},
		{'D', 'H'}, {'H', 'D'}, {'N', 'N'},
		{'a', 't'}, {'t', 'a'}, {'c', 'g'}, {'g', 'c'},
	}
	for _, tc := range tests {
		got := Complement(tc.input, TypeDNA)
		if got != tc.expected {
			t.Errorf("Complement(%q, DNA) = %q; want %q", tc.input, got, tc.expected)
		}
	}
}

func TestComplementRNA(t *testing.T) {
	tests := []struct {
		input    byte
		expected byte
	}{
		{'A', 'U'}, {'U', 'A'}, {'C', 'G'}, {'G', 'C'},
		{'T', 'A'},
		{'R', 'Y'}, {'Y', 'R'}, {'S', 'S'}, {'W', 'W'},
		{'K', 'M'}, {'M', 'K'}, {'B', 'V'}, {'V', 'B'},
		{'D', 'H'}, {'H', 'D'}, {'N', 'N'},
	}
	for _, tc := range tests {
		got := Complement(tc.input, TypeRNA)
		if got != tc.expected {
			t.Errorf("Complement(%q, RNA) = %q; want %q", tc.input, got, tc.expected)
		}
	}
}

func TestReverseComplementDNA(t *testing.T) {
	input := "ATCG"
	expected := "CGAT"
	got := ReverseComplement(input, TypeDNA)
	if got != expected {
		t.Errorf("ReverseComplement(%q, DNA) = %q; want %q", input, got, expected)
	}
}

func TestReverseComplementRNA(t *testing.T) {
	input := "AUCG"
	expected := "CGAU"
	got := ReverseComplement(input, TypeRNA)
	if got != expected {
		t.Errorf("ReverseComplement(%q, RNA) = %q; want %q", input, got, expected)
	}
}

func TestReverseComplementLongDNA(t *testing.T) {
	input := "ATGCGTACGGTAGCTA"
	expected := "TAGCTACCGTACGCAT"
	got := ReverseComplement(input, TypeDNA)
	if got != expected {
		t.Errorf("ReverseComplement(%q, DNA) = %q; want %q", input, got, expected)
	}
}

func TestReverseComplementIUPAC(t *testing.T) {
	input := "RYMK"
	expected := "MKRY"
	got := ReverseComplement(input, TypeDNA)
	if got != expected {
		t.Errorf("ReverseComplement(%q, DNA) = %q; want %q", input, got, expected)
	}
}

func TestToUpper(t *testing.T) {
	input := "atgc"
	expected := "ATGC"
	got := ToUpper(input)
	if got != expected {
		t.Errorf("ToUpper(%q) = %q; want %q", input, got, expected)
	}
}
