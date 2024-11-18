package documentvalidator

import (
	"context"
	"document-validator/internal"
	"fmt"
	"regexp"
)

// NationalIDValidator is a validator for national id documents
type NationalIDValidator struct{}

func (v *NationalIDValidator) Validate(ctx context.Context, doc internal.Document) error {
	// Validate the national id
	if reg, _ := regexp.Compile(`^[0-9]{9}$`); !reg.MatchString(doc.Value) {
		fmt.Printf("Invalid national id: %s\n", doc.Value)
		return fmt.Errorf("invalid national id: %w", internal.ErrInvalidNationalID)
	}

	// validate the issue country format
	if reg, _ := regexp.Compile(`^[A-Z]{2}$`); !reg.MatchString(doc.IssueCountry) {
		fmt.Printf("Invalid issue country: %s\n", doc.IssueCountry)
		return fmt.Errorf("invalid issue country: %w", internal.ErrInvalidIssueCountry)
	}

	// validate the issue date and expiry date format
	reg, _ := regexp.Compile(`^[0-9]{4}-[0-9]{2}-[0-9]{2}$`)

	if !reg.MatchString(doc.IssueDate) {
		fmt.Printf("Invalid issue date: %s\n", doc.IssueDate)
		return fmt.Errorf("invalid issue date: %w", internal.ErrInvalidIssueDate)
	}

	if !reg.MatchString(doc.ExpiryDate) {
		fmt.Printf("Invalid expiry date: %s\n", doc.ExpiryDate)
		return fmt.Errorf("invalid expiry date: %w", internal.ErrInvalidExpiryDate)
	}

	fmt.Println("National ID document is valid")

	return nil
}
