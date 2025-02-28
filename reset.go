package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	// if cfg.PLATFORM != "dev" {
	// 	responseWithError(w, http.StatusForbidden, "This endpoint is only available in development environment", nil)
	// 	return
	// }
	err := cfg.db.DeleteAllUsers(r.Context())
	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Failed to reset database", nil)
		return
	}
	cfg.fileserverHits.Store(0)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hits: %d", cfg.fileserverHits.Load())
}
