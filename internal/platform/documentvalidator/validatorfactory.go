package documentvalidator

import (
	"document-validator/internal"
)

// ValidatorFactory is an interface for getting a document validator
type ValidatorFactory interface {
	// GetValidator returns a document validator for the given document type
	GetValidator(docType string) (internal.DocumentValidator, error)
}

// DocumentValidatorFactory is a factory for getting a document validator
type DocumentValidatorFactory struct {
	validators map[string]internal.DocumentValidator
}

// NewDocumentValidatorFactory creates a new DocumentValidatorFactory
func NewDocumentValidatorFactory(docs []string) (*DocumentValidatorFactory, error) {
	factory := &DocumentValidatorFactory{
		validators: make(map[string]internal.DocumentValidator),
	}

	for _, doc := range docs {
		validator, exists := ValidatorRegistry[doc]
		if !exists {
			return nil, internal.ErrValidatorNotFound
		}
		factory.validators[doc] = validator
	}

	return factory, nil
}

// GetValidator returns a document validator for the given document type
func (f *DocumentValidatorFactory) GetValidator(docType string) (internal.DocumentValidator, error) {
	validator, exists := f.validators[docType]
	if !exists {
		return nil, internal.ErrInvalidDocumentType
	}
	return validator, nil
}
