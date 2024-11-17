package internal

import "errors"

var (
	ErrInvalidNationalID   = errors.New("invalid national id")
	ErrInvalidIssueCountry = errors.New("invalid issue country")
	ErrInvalidIssueDate    = errors.New("invalid issue date")
	ErrInvalidExpiryDate   = errors.New("invalid expiry date")
	ErrInvalidKTN          = errors.New("invalid KTN")
	ErrInvalidDocumentType = errors.New("invalid document type")
)
