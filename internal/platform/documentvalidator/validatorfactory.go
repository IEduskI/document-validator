package documentvalidator

import (
	"document-validator/internal"
	"fmt"
)

func DocumentValidatorFactory(docType string) (internal.DocumentValidator, error) {
	switch docType {
	case "I":
		return &NationalIDValidator{}, nil
	case "K":
		return &KTNValidator{}, nil
	default:
		return nil, fmt.Errorf("invalid document type: %s: %w", docType, internal.ErrInvalidDocumentType)
	}
}
