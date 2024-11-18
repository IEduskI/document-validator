package rest

import (
	"document-validator/internal"
	"document-validator/internal/operations/validate"
	"encoding/json"
	"errors"
	"net/http"
)

type ValidateDocumentHandler struct {
	validateService validate.Service
}

func NewValidateDocumentHandler(validateService validate.Service) *ValidateDocumentHandler {
	return &ValidateDocumentHandler{
		validateService: validateService,
	}
}

func (h *ValidateDocumentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Decode the request
	w.Header().Set("Content-Type", "application/json")
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// build the request
	request := buildRequest(req)

	// call the validate service
	if err := h.validateService.Validate(r.Context(), request); err != nil {
		// write the error response
		buildErrorResponse(err, w)
		return
	}

	// write the success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(buildResponse("Documents are valid"))
}

func buildRequest(req Request) validate.Request {
	var documents []validate.Document
	for _, doc := range req.Documents {
		documents = append(documents, validate.Document{
			Type:         doc.Type,
			Value:        doc.Value,
			IssueCountry: doc.IssueCountry,
			IssueDate:    doc.IssueDate,
			ExpiryDate:   doc.ExpiryDate,
		})
	}

	return validate.Request{
		Document: documents,
	}
}

func buildResponse(message string) Response {
	return Response{
		Message: message,
	}
}

func buildErrorResponse(err error, w http.ResponseWriter) {
	var response Error
	switch {
	case errors.Is(err, internal.ErrInvalidNationalID):
		w.WriteHeader(http.StatusBadRequest)
		response = Error{Description: internal.ErrInvalidNationalID.Error()}
	case errors.Is(err, internal.ErrInvalidIssueCountry):
		w.WriteHeader(http.StatusBadRequest)
		response = Error{Description: internal.ErrInvalidIssueCountry.Error()}
	case errors.Is(err, internal.ErrInvalidIssueDate):
		w.WriteHeader(http.StatusBadRequest)
		response = Error{Description: internal.ErrInvalidIssueDate.Error()}
	case errors.Is(err, internal.ErrInvalidExpiryDate):
		w.WriteHeader(http.StatusBadRequest)
		response = Error{Description: internal.ErrInvalidExpiryDate.Error()}
	case errors.Is(err, internal.ErrInvalidKTN):
		w.WriteHeader(http.StatusBadRequest)
		response = Error{Description: internal.ErrInvalidKTN.Error()}
	case errors.Is(err, internal.ErrInvalidDocumentType):
		w.WriteHeader(http.StatusBadRequest)
		response = Error{Description: internal.ErrInvalidDocumentType.Error()}
	case errors.Is(err, internal.ErrValidatorNotFound):
		w.WriteHeader(http.StatusBadRequest)
		response = Error{Description: internal.ErrValidatorNotFound.Error()}
	default:
		w.WriteHeader(http.StatusInternalServerError)
		response = Error{Description: "Internal server error"}
	}

	json.NewEncoder(w).Encode(response)
}
