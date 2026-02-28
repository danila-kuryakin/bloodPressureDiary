package rest

import (
	"bloodPressureDiary/bp-service/internal/model"
	"bloodPressureDiary/bp-service/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

type TagHandler struct {
	svc *service.Service
}

func NewTagHandler(s *service.Service) *TagHandler {
	return &TagHandler{svc: s}
}

func (h *TagHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/tags", h.tags)
	mux.HandleFunc("/tags/", h.tagByID)
}

func (h *TagHandler) tags(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("user_id")

	switch r.Method {
	case http.MethodPost:
		var t model.UserTag
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			return
		}

		t.UserID = userId

		err = h.svc.Tag.Create(r.Context(), &t)
		if err != nil {
			log.Println("Create tags:", err)
			return
		}
		err = json.NewEncoder(w).Encode(t)
		if err != nil {
			return
		}

	case http.MethodGet:
		res, _ := h.svc.Tag.List(r.Context(), userId)
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			return
		}
	}
}

func (h *TagHandler) tagByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("user_id")
	tagName := r.URL.Path[len("/tags/"):]
	switch r.Method {
	case http.MethodPut:
		var body struct{ Name string }
		json.NewDecoder(r.Body).Decode(&body)
		h.svc.Tag.Rename(r.Context(), tagName, userID, body.Name)

	case http.MethodDelete:
		err := h.svc.Tag.Delete(r.Context(), tagName, userID)
		if err != nil {
			return
		}
	}
}
