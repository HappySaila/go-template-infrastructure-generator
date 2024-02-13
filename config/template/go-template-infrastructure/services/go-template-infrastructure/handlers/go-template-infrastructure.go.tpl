package handlers

import (
	"context"
	"dev/[[.Dashed]].git/foundation/web"
	"net/http"
)

// create ...
func ([[.Letter]] [[.CapsCamel]]) create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	// Get request data
	var request = struct {
		Value string `validate:"required"`
	}{}

	// Decode, sanitize & validate request
	err := web.Decode(r, &request)
	if err != nil {
		return err
	}

	// Log
	[[.Letter]].Service.Log.Printf("Creating...")

	// Create
	err = [[.Letter]].Service.Create(ctx)
	if err != nil {
		return err
	}

	// Send response data
	response := struct {
		Status string `json:"Status"`
	}{
		Status: "Success",
	}
	return web.Respond(ctx, w, response, http.StatusOK)

}
