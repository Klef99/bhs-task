package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Klef99/bhs-task/internal/entity"
	"github.com/Klef99/bhs-task/internal/usecase"
	"github.com/Klef99/bhs-task/pkg/jwtgenerator"
	"github.com/Klef99/bhs-task/pkg/logger"
	"github.com/go-chi/chi/v5"
)

type userRoutes struct {
	t   usecase.User
	l   logger.Interface
	jtg jwtgenerator.Interface
}

func NewUserRoutes(handler chi.Router, t usecase.User, l logger.Interface, jtg jwtgenerator.Interface) {
	rt := &userRoutes{t: t, l: l, jtg: jtg}
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Post("/register", rt.Register)
		r.Get("/login", rt.Login)
	})
	handler.Mount("/user", router)
}

type registerResponse struct {
	Status string `json:"status"`
}

// @Summary     Registration
// @Description Registration in the system
// @ID          register
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Success     200 {object} registerResponse
// @Failure     500 {object} response
// @Router      /register [post]
func (rt *userRoutes) Register(w http.ResponseWriter, r *http.Request) {
	crd := entity.Credentials{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&crd)
	if err != nil {
		rt.l.Error(err, "http - v1 - register")
		errorResponse(w, http.StatusInternalServerError, "error decoding request body")
		return
	}
	status, err := rt.t.Register(r.Context(), crd)
	if err != nil {
		rt.l.Error(err, "http - v1 - register")
		errorResponse(w, http.StatusInternalServerError, "error registering user or user already exists")
		return
	}
	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(registerResponse{"User registered successfully"})
	}
}

type loginResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

// @Summary     Login
// @Description Sign in in the system
// @ID          login
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Success     200 {object} loginResponse
// @Failure     500 {object} response
// @Router      /user/login [get]
// @Param request body entity.Credentials true "query params"
func (rt *userRoutes) Login(w http.ResponseWriter, r *http.Request) {
	crd := entity.Credentials{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&crd)
	if err != nil {
		rt.l.Error(err, "http - v1 - login - decoder.Decode")
		errorResponse(w, http.StatusInternalServerError, "error decoding request body")
		return
	}
	status, err := rt.t.Login(r.Context(), crd)
	if err != nil {
		rt.l.Error(err, "http - v1 - login - rt.t.Login")
		errorResponse(w, http.StatusInternalServerError, "error login user")
		return
	}
	if status {
		token, err := rt.jtg.GenerateToken(crd.Username)
		if err != nil {
			rt.l.Error(err, "http - v1 - login - rt.jtg.GenerateToken")
			errorResponse(w, http.StatusInternalServerError, "error generating token")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(loginResponse{"Success", token})
	}
}
