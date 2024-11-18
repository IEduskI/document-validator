package validate

import (
	"context"
	"document-validator/internal"
	"document-validator/internal/platform/documentvalidator"
)

type Service interface {
	Validate(ctx context.Context, request Request) error
}

type service struct {
	factory documentvalidator.ValidatorFactory
}

func NewService(factory documentvalidator.ValidatorFactory) *service {
	return &service{
		factory: factory,
	}
}

func (s *service) Validate(ctx context.Context, request Request) error {

	// Iterate over the documents
	for _, doc := range request.Document {
		// Validate the document
		docValidator, err := s.factory.GetValidator(doc.Type)
		if err != nil {
			return err
		}

		parsedDoc := buildDocument(doc)

		if err := docValidator.Validate(ctx, parsedDoc); err != nil {
			return err
		}
	}

	return nil
}

func buildDocument(doc Document) internal.Document {
	return internal.Document{
		Type:         doc.Type,
		Value:        doc.Value,
		IssueCountry: doc.IssueCountry,
		IssueDate:    doc.IssueDate,
		ExpiryDate:   doc.ExpiryDate,
	}
}
