package v1

import (
	"encoding/json"
	"net/http"

	"github.com/Klef99/bhs-task/internal/entity"
	"github.com/Klef99/bhs-task/internal/usecase"
	"github.com/Klef99/bhs-task/pkg/jwtgenerator"
	"github.com/Klef99/bhs-task/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type userRoutes struct {
	t   usecase.User
	l   logger.Interface
	jtg jwtgenerator.Interface
}

func NewUserRoutes(handler chi.Router, t usecase.User, l logger.Interface, jtg jwtgenerator.Interface) {
	rt := &userRoutes{t: t, l: l, jtg: jtg}
	router := chi.NewRouter()
	tokenAuth := rt.jtg.GetJWTAuth()
	router.Group(func(r chi.Router) {
		r.Post("/register", rt.Register)
		r.Get("/login", rt.Login)
	})
	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))
		r.Post("/deposit/make", rt.Deposit)
		r.Get("/deposit/check", rt.CheckDeposit)
	})
	handler.Mount("/", router)
}

// @Summary     Registration
// @Description Registration in the system
// @ID          register
// @Tags  	    Authentication
// @Accept      json
// @Produce     json
// @Success     200 {object} response
// @Failure     500 {object} response
// @Router      /register [post]
// @Param request body entity.Credentials true "query params"
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
		json.NewEncoder(w).Encode(response{"User registered successfully"})
	}
}

type loginResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
}

// @Summary     Login
// @Description Sign in in the system
// @ID          login
// @Tags  	    Authentication
// @Accept      json
// @Produce     json
// @Success     200 {object} response
// @Failure     500 {object} response
// @Router      /login [get]
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
	user, err := rt.t.Login(r.Context(), crd)
	if err != nil {
		rt.l.Error(err, "http - v1 - login - rt.t.Login")
		errorResponse(w, http.StatusInternalServerError, "error login user")
		return
	}
	if user.Id != 0 {
		token, err := rt.jtg.GenerateToken(user.Username, user.Id)
		if err != nil {
			rt.l.Error(err, "http - v1 - login - rt.jtg.GenerateToken")
			errorResponse(w, http.StatusInternalServerError, "error generating token")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(loginResponse{"Success", token})
	}
}

type depositRequest struct {
	Amount float64 `json:"amount"`
}

// @Summary     Make a deposit
// @Description Make a deposit
// @ID          MakeDeposit
// @Tags  	    Deposit
// @Accept      json
// @Produce     json
// @Success     200 {object} response
// @Failure     500 {object} response
// @Router      /deposit/make [post]
// @Param request body depositRequest true "query params"
func (rt *userRoutes) Deposit(w http.ResponseWriter, r *http.Request) {
	req := depositRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		rt.l.Error(err, "http - v1 - Deposit")
		errorResponse(w, http.StatusInternalServerError, "error decoding request body")
		return
	}
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		rt.l.Error(err, "http - v1 - Deposit - jwtauth.FromContext")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	id, ok := claims["id"].(float64)
	if !ok {
		rt.l.Error(err, "http - v1 - Deposit - .(float64)")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	name, ok := claims["name"].(string)
	if !ok {
		rt.l.Error(err, "http - v1 - Deposit - .(string)")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	usr := entity.User{Username: name, Id: int64(id)}
	status, err := rt.t.MakeDeposit(r.Context(), usr, req.Amount)
	if err != nil {
		rt.l.Error(err, "http - v1 - Deposit - rt.t.Deposit")
		errorResponse(w, http.StatusInternalServerError, "error depositing money")
		return
	}
	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response{"Deposit successful"})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response{"Error depositing money"})
	}
}

type checkDepositResponse struct {
	Amount float64 `json:"amount"`
}

// @Summary     Get current deposit
// @Description Get curret balance
// @ID          CheckDeposit
// @Tags  	    Deposit
// @Accept      json
// @Produce     json
// @Success     200 {object} response
// @Failure     500 {object} response
// @Router      /deposit/check [get]
// @Param request body depositRequest true "query params"
func (rt *userRoutes) CheckDeposit(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		rt.l.Error(err, "http - v1 - CheckDeposit - jwtauth.FromContext")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	id, ok := claims["id"].(float64)
	if !ok {
		rt.l.Error(err, "http - v1 - CheckDeposit - .(float64)")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	name, ok := claims["name"].(string)
	if !ok {
		rt.l.Error(err, "http - v1 - CheckDeposit - .(string)")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	usr := entity.User{Username: name, Id: int64(id)}
	deposit, err := rt.t.CheckDeposit(r.Context(), usr)
	if err != nil {
		rt.l.Error(err, "http - v1 - CheckDeposit - rt.t.CheckDeposit")
		errorResponse(w, http.StatusInternalServerError, "error depositing money")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(checkDepositResponse{deposit})

}
