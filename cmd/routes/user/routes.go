package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/user/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/user/register", h.handleRegistration).Methods("GET")
	// router.HandleFunc("/user/{id}", h.getUser).Methods("GET")
	// router.HandleFunc("/user/{id}", h.updateUser).Methods("PUT")
	// router.HandleFunc("/user/{id}", h.deleteUser).Methods("DELETE")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) handleRegistration(w http.ResponseWriter, r *http.Request) {
}
