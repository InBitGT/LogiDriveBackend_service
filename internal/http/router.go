package http

import (
	"net/http"

	"backend_service/internal/deliveries"
	"backend_service/internal/routes"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Health
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Backend OK"))
	})

	// Rutas (dominio)
	mux.HandleFunc("/routes", routes.ListRoutes)         // GET /routes
	mux.HandleFunc("/routes/create", routes.CreateRoute) // POST /routes/create

	// Deliveries (dominio)
	mux.HandleFunc("/deliveries", deliveries.ListDeliveries)        // GET /deliveries
	mux.HandleFunc("/deliveries/create", deliveries.CreateDelivery) // POST /deliveries/create

	return mux
}
