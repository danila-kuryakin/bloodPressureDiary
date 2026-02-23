package rest

import (
	"bloodPressureDiary/bp-service/internal/model"
	"bloodPressureDiary/bp-service/internal/service"
	"encoding/json"
	"net/http"
)

type PressureHandler struct {
	svc *service.PressureService
}

func NewPressureHandler(s *service.PressureService) *PressureHandler {
	return &PressureHandler{svc: s}
}

func (h *PressureHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/pressure", h.pressure)
}

func (h *PressureHandler) pressure(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")

	switch r.Method {
	case http.MethodPost:
		var p model.BloodPressure
		json.NewDecoder(r.Body).Decode(&p)
		p.UserID = userID
		h.svc.Create(r.Context(), &p)
		json.NewEncoder(w).Encode(p)

	case http.MethodGet:
		res, _ := h.svc.List(r.Context(), userID)
		json.NewEncoder(w).Encode(res)
	}
}
