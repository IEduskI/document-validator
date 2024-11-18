package documentvalidator

import (
	"document-validator/internal"
	"fmt"
)

type ValidatorFactory interface {
	GetValidator(docType string) (internal.DocumentValidator, error)
}

type DocumentValidatorFactory struct{}

func NewDocumentValidatorFactory() *DocumentValidatorFactory {
	return &DocumentValidatorFactory{}
}

func (vf *DocumentValidatorFactory) GetValidator(docType string) (internal.DocumentValidator, error) {
	switch docType {
	case "I":
		return &NationalIDValidator{}, nil
	case "K":
		return &KTNValidator{}, nil
	default:
		return nil, fmt.Errorf("invalid document type: %s: %w", docType, internal.ErrInvalidDocumentType)
	}
}
