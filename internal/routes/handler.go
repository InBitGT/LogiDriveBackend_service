package routes

import (
	"net/http"

	"backend_service/internal/common"
)

func ListRoutes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		common.JSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	common.JSON(w, http.StatusOK, []map[string]any{
		{"id": 1, "name": "Ruta A"},
		{"id": 2, "name": "Ruta B"},
	})
}

func CreateRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		common.JSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	// TODO: leer body, validar y crear
	common.JSON(w, http.StatusCreated, map[string]string{"status": "route created"})
}
