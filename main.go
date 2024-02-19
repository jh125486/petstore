package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	mw "github.com/oapi-codegen/nethttp-middleware"
	"github.com/swaggest/swgui/v5cdn"

	"github.com/jh125486/petstore/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a new chi router.
	r := chi.NewRouter()
	// Serve the OpenAPI schema.
	spec, err := api.GetSwagger()
	if err != nil {
		log.Fatalf("Error loading swagger spec\n: %s", err)
	}
	r.Get("/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		if err = json.NewEncoder(w).Encode(spec); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	// Serve the Swagger UI.
	r.Handle("/docs/", v5cdn.New("Petstore", "/openapi.json", "/docs/"))

	// Create a new router for the API to isolate the request validator middleware.
	r.Group(func(r chi.Router) {
		// Use our validation middleware to check all requests against the OpenAPI schema.
		r.Use(mw.OapiRequestValidator(spec))
		// Create an instance of our handler which satisfies the generated interface.
		petStoreHandler := api.NewStrictHandler(api.NewPetStore(), nil)
		// Register our petStore above as the handler for the interface
		api.HandlerFromMux(petStoreHandler, r)
	})

	// Serve HTTP until the world ends.
	log.Fatal(http.ListenAndServe(":"+port, r))
}
