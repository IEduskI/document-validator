package documentvalidator

import (
	"context"
	"document-validator/internal"
	"testing"
)

func TestKTNValidator_Validate(t *testing.T) {
	tests := []struct {
		name          string
		doc           internal.Document
		expectedError error
	}{
		{
			name: "valid KTN and issue country",
			doc: internal.Document{
				Value:        "123456789",
				IssueCountry: "US",
			},
			expectedError: nil,
		},
		{
			name: "invalid KTN",
			doc: internal.Document{
				Value:        "12345",
				IssueCountry: "US",
			},
			expectedError: internal.ErrInvalidKTN,
		},
		{
			name: "invalid issue country",
			doc: internal.Document{
				Value:        "123456789",
				IssueCountry: "USA",
			},
			expectedError: internal.ErrInvalidIssueCountry,
		},
	}

	validator := &KTNValidator{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.Validate(context.Background(), tt.doc)
			if err != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
			if err == nil && tt.expectedError != nil {
				t.Errorf("expected error %v, got nil", tt.expectedError)
			}
		})
	}
}
