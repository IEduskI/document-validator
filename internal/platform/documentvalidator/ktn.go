package documentvalidator

import (
	"context"
	"document-validator/internal"
	"fmt"
	"regexp"
)

// KTNValidator is a validator for KTN documents
type KTNValidator struct{}

func (v *KTNValidator) Validate(ctx context.Context, doc internal.Document) error {
	// validat the ktn value 9 or 10-digit number that can be a combination of letters and numbers
	if reg, _ := regexp.Compile(`^[0-9A-Z]{9,10}$`); !reg.MatchString(doc.Value) {
		fmt.Printf("Invalid KTN: %s\n", doc.Value)
		return internal.ErrInvalidKTN
	}

	// validate the issue country format
	if reg, _ := regexp.Compile(`^[A-Z]{2}$`); !reg.MatchString(doc.IssueCountry) {
		fmt.Printf("Invalid issue country: %s\n", doc.IssueCountry)
		return internal.ErrInvalidIssueCountry
	}

	fmt.Println("KTN document is valid")

	return nil
}
