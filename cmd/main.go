package main

import (
	"document-validator/internal/operations/validate"
	"document-validator/internal/platform/documentvalidator"
	"document-validator/internal/platform/rest"
	"log"
	"net/http"
)

func main() {
	factory := documentvalidator.NewDocumentValidatorFactory()
	service := validate.NewService(factory)
	handler := rest.NewValidateDocumentHandler(service)

	http.HandleFunc("POST /document/validate", handler.ServeHTTP)

	log.Println("Servidor iniciado en http://localhost:90")
	if err := http.ListenAndServe(":90", nil); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}
