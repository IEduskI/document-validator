package documentvalidator

import (
	"context"
	"document-validator/internal"
	"testing"
)

func TestNationalIDValidator_Validate(t *testing.T) {
	tests := []struct {
		name    string
		doc     internal.Document
		wantErr bool
	}{
		{
			name: "valid document",
			doc: internal.Document{
				Value:        "123456789",
				IssueCountry: "US",
				IssueDate:    "2020-01-01",
				ExpiryDate:   "2030-01-01",
			},
			wantErr: false,
		},
		{
			name: "invalid national id",
			doc: internal.Document{
				Value:        "12345",
				IssueCountry: "US",
				IssueDate:    "2020-01-01",
				ExpiryDate:   "2030-01-01",
			},
			wantErr: true,
		},
		{
			name: "invalid issue country",
			doc: internal.Document{
				Value:        "123456789",
				IssueCountry: "USA",
				IssueDate:    "2020-01-01",
				ExpiryDate:   "2030-01-01",
			},
			wantErr: true,
		},
		{
			name: "invalid issue date",
			doc: internal.Document{
				Value:        "123456789",
				IssueCountry: "US",
				IssueDate:    "2020-01-1",
				ExpiryDate:   "2030-01-01",
			},
			wantErr: true,
		},
		{
			name: "invalid expiry date",
			doc: internal.Document{
				Value:        "123456789",
				IssueCountry: "US",
				IssueDate:    "2020-01-01",
				ExpiryDate:   "2030-1-01",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &NationalIDValidator{}
			if err := v.Validate(context.Background(), tt.doc); (err != nil) != tt.wantErr {
				t.Errorf("NationalIDValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
