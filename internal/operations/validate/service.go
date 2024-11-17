package validate

import (
	"context"
	"document-validator/internal"
	"document-validator/internal/platform/documentvalidator"
)

type Service interface {
	Validate(ctx context.Context, request Request) error
}

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) Validate(ctx context.Context, request Request) error {
	// Create a document with the request values
	doc := internal.Document{
		Type:         request.Type,
		Value:        request.Value,
		IssueCountry: request.IssueCountry,
		IssueDate:    request.IssueDate,
		ExpiryDate:   request.ExpiryDate,
	}

	// Validate the document
	docValidator, err := documentvalidator.DocumentValidatorFactory(doc.Type)
	if err != nil {
		return err
	}

	return docValidator.Validate(ctx, doc)
}
