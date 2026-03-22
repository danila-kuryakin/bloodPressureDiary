package server_rest

import (
	"bp-service/internal/model"
	"bp-service/internal/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

type PressureHandler struct {
	svc    *service.Service
	buffer map[string]map[string]any
}

func NewPressureHandler(s *service.Service) *PressureHandler {
	return &PressureHandler{svc: s, buffer: make(map[string]map[string]any)}
}

func (h *PressureHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/pressure", h.Pressure)
	mux.HandleFunc("/pressure/", h.PressureById)
}

func (h *PressureHandler) Pressure(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("user_id")
	tokenApi := r.Header.Get("token_api")

	if tokenApi != os.Getenv("SERVICE_TOKEN_API") {
		return
	}

	switch r.Method {
	case http.MethodPost:
		var p model.BloodPressure
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			return
		}
		p.UserID = userID
		err = h.svc.Pressure.Create(r.Context(), &p)
		if err != nil {
			log.Println("pressure create error", err)
			return
		}
		err = json.NewEncoder(w).Encode(p)
		if err != nil {
			return
		}
		return

	case http.MethodGet:
		pressures, err := h.svc.Pressure.List(r.Context(), userID)
		if err != nil {
			log.Println("List. Pressure list error", err)
			return
		}

		err = json.NewEncoder(w).Encode(pressures)
		if err != nil {
			log.Println("Encode. Pressure list error", err)
			return
		}
		if h.buffer[userID] == nil {
			h.buffer[userID] = map[string]any{}
		}

		for id, pressure := range pressures {
			keyStr := fmt.Sprintf("LastPressureList%v", id)
			h.buffer[userID][keyStr] = pressure
			//fmt.Println(keyStr, pressure)
		}
		return
	}
}

func (h *PressureHandler) PressureById(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("user_id")
	tokenApi := r.Header.Get("token_api")

	if tokenApi != os.Getenv("SERVICE_TOKEN_API") {
		return
	}

	pressureId := r.URL.Path[len("/pressure/"):]
	switch r.Method {
	case http.MethodDelete:
		pressureIdInt, err := strconv.Atoi(pressureId)
		if err != nil {
			log.Println("Atoi. Pressure id error:", err)
			return
		}
		//pressureIdInt -= 1
		fmt.Println("MethodDelete", userID, pressureIdInt)

		keyStr := fmt.Sprintf("LastPressureList%v", pressureIdInt)

		pressure := h.buffer[userID][keyStr].(*model.BloodPressure)
		fmt.Println("pressure", pressure)

		if h.svc.Pressure.Delete(r.Context(), pressure.ID, userID) != nil {
			return
		}
		return
	}
}
