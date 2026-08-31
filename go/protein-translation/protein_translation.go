package proteintranslation

import "errors"

var (
	ErrInvalidBase = errors.New("invalid base")
	ErrStop        = errors.New("stop codon")
)

func FromRNA(rna string) ([]string, error) {
	proteins := []string{}

	for i := 0; i < len(rna); i += 3 {
		if i+3 > len(rna) {
			return proteins, ErrInvalidBase
		}

		amino := rna[i : i+3]
		result, err := FromCodon(amino)

		if err == ErrStop {
			return proteins, nil
		}

		if err != nil {
			return proteins, err
		}

		proteins = append(proteins, result)
	}

	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	codonMap := map[string]string{
		"AUG": "Methionine",

		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",

		"UUA": "Leucine",
		"UUG": "Leucine",

		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",

		"UAU": "Tyrosine",
		"UAC": "Tyrosine",

		"UGU": "Cysteine",
		"UGC": "Cysteine",

		"UGG": "Tryptophan",

		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
	protein, ok := codonMap[codon]
	if !ok {
		return "", ErrInvalidBase
	}
	if protein == "STOP" {
		return "", ErrStop
	}

	return protein, nil
}
