package rest

import (
	"bloodPressureDiary/bp-service/internal/model"
	"bloodPressureDiary/bp-service/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
)

type TagHandler struct {
	svc *service.TagService
}

func NewTagHandler(s *service.TagService) *TagHandler {
	return &TagHandler{svc: s}
}

func (h *TagHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/tags", h.tags)
	mux.HandleFunc("/tags/", h.tagByID)
}

func (h *TagHandler) tags(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")

	switch r.Method {
	case http.MethodPost:
		var t model.UserTag
		json.NewDecoder(r.Body).Decode(&t)
		t.UserID = userID
		h.svc.Create(r.Context(), &t)
		json.NewEncoder(w).Encode(t)

	case http.MethodGet:
		res, _ := h.svc.List(r.Context(), userID)
		json.NewEncoder(w).Encode(res)
	}
}

func (h *TagHandler) tagByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("X-User-ID")
	id, _ := strconv.ParseInt(r.URL.Path[len("/tags/"):], 10, 64)

	switch r.Method {
	case http.MethodPut:
		var body struct{ Name string }
		json.NewDecoder(r.Body).Decode(&body)
		h.svc.Rename(r.Context(), id, userID, body.Name)

	case http.MethodDelete:
		h.svc.Delete(r.Context(), id, userID)
	}
}
