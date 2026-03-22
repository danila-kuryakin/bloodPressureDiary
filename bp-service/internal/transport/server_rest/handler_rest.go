package server_rest

import (
	"bp-service/internal/service"
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type PressureHandlerInterface interface {
	Register(mux *http.ServeMux)
}

type TagHandlerInterface interface {
	Register(mux *http.ServeMux)
}

type Handler struct {
	PressureHandlerInterface
	TagHandlerInterface
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		PressureHandlerInterface: NewPressureHandler(svc),
		TagHandlerInterface:      NewTagHandler(svc),
	}
}

func (h *Handler) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("=== FULL REQUEST ===")
		fmt.Println("Method:", r.Method)
		fmt.Println("URL:", r.URL.String())
		fmt.Println("Headers:", r.Header)

		if r.Body != nil {
			bodyBytes, err := io.ReadAll(r.Body)
			if err != nil {
				fmt.Println("Error reading body:", err)
			} else {
				fmt.Println("Body:", string(bodyBytes))

				// ВАЖНО: возвращаем body обратно,
				// иначе handler не сможет его прочитать
				r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) Register(mux *http.ServeMux) {
	//h.LogRequest(mux)
	h.PressureHandlerInterface.Register(mux)
	h.TagHandlerInterface.Register(mux)
}
