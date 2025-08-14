package deliveries

import (
	"net/http"

	"backend_service/internal/common"
)

func ListDeliveries(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		common.JSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	common.JSON(w, http.StatusOK, []map[string]any{
		{"id": 101, "status": "pending"},
		{"id": 102, "status": "delivered"},
	})
}

func CreateDelivery(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		common.JSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	// TODO: leer body, validar y crear
	common.JSON(w, http.StatusCreated, map[string]string{"status": "delivery created"})
}
