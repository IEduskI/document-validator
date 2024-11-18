package documentvalidator

import "document-validator/internal"

// ValidatorRegistry is a map of document type to document validator
var ValidatorRegistry = map[string]internal.DocumentValidator{
	"I":   &NationalIDValidator{},
	"KTN": &KTNValidator{},
}
