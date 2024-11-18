package main

import (
	"document-validator/internal/operations/validate"
	"document-validator/internal/platform/documentvalidator"
	"document-validator/internal/platform/rest"
	"document-validator/utils"
	"log"
	"net/http"
)

func main() {
	config := utils.NewConfig()

	factory, err := documentvalidator.NewDocumentValidatorFactory(config.GetServiceDocumentsTypes())
	if err != nil {
		log.Fatalf("Error creating document validator factory: %v", err)
	}

	service := validate.NewService(factory)
	handler := rest.NewValidateDocumentHandler(service)

	http.HandleFunc("POST /document/validate", handler.ServeHTTP)

	log.Println("Servidor iniciado en http://localhost:90")
	if err := http.ListenAndServe(":90", nil); err != nil {
		log.Fatalf("Error iniciando el servidor: %v", err)
	}
}
