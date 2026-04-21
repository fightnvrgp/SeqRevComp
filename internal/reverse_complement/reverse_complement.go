// Package reverse_complement implements DNA/RNA reverse complement logic.
package reverse_complement

import (
	"strings"
)

// SequenceType indicates whether the sequence is DNA or RNA.
type SequenceType int

const (
	TypeDNA SequenceType = iota
	TypeRNA
)

// Complement maps each nucleotide to its complement.
// Supports standard bases and IUPAC ambiguity codes.
func Complement(base byte, seqType SequenceType) byte {
	// DNA complement map
	dnaMap := map[byte]byte{
		'A': 'T', 'T': 'A',
		'C': 'G', 'G': 'C',
		'U': 'A', // U in DNA context is treated as T complement (A)
		'R': 'Y', 'Y': 'R',
		'S': 'S', 'W': 'W',
		'K': 'M', 'M': 'K',
		'B': 'V', 'V': 'B',
		'D': 'H', 'H': 'D',
		'N': 'N',
		// lowercase
		'a': 't', 't': 'a',
		'c': 'g', 'g': 'c',
		'u': 'a',
		'r': 'y', 'y': 'r',
		's': 's', 'w': 'w',
		'k': 'm', 'm': 'k',
		'b': 'v', 'v': 'b',
		'd': 'h', 'h': 'd',
		'n': 'n',
	}

	rnaMap := map[byte]byte{
		'A': 'U', 'U': 'A',
		'C': 'G', 'G': 'C',
		'T': 'A', // T in RNA context is treated as U complement (A)
		'R': 'Y', 'Y': 'R',
		'S': 'S', 'W': 'W',
		'K': 'M', 'M': 'K',
		'B': 'V', 'V': 'B',
		'D': 'H', 'H': 'D',
		'N': 'N',
		// lowercase
		'a': 'u', 'u': 'a',
		'c': 'g', 'g': 'c',
		't': 'a',
		'r': 'y', 'y': 'r',
		's': 's', 'w': 'w',
		'k': 'm', 'm': 'k',
		'b': 'v', 'v': 'b',
		'd': 'h', 'h': 'd',
		'n': 'n',
	}

	if seqType == TypeRNA {
		if c, ok := rnaMap[base]; ok {
			return c
		}
	} else {
		if c, ok := dnaMap[base]; ok {
			return c
		}
	}
	return base
}

// ReverseComplement returns the reverse complement of a sequence.
func ReverseComplement(seq string, seqType SequenceType) string {
	result := make([]byte, len(seq))
	for i := 0; i < len(seq); i++ {
		result[len(seq)-1-i] = Complement(seq[i], seqType)
	}
	return string(result)
}

// ToUpper converts a sequence to uppercase.
func ToUpper(seq string) string {
	return strings.ToUpper(seq)
}
