package repository

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func NewSheetService(client *http.Client) *sheets.Service {
	ctx := context.Background()
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	return srv
}
