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
		r.Post("/login", rt.Login)
	})
	router.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))
		r.Post("/deposit", rt.Deposit)
		r.Get("/deposit", rt.CheckDeposit)
	})
	handler.Mount("/", router)
}

// @Summary     User Registration
// @Description Handles user registration by accepting credentials and registering a new user in the system.
// @ID          register
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Success     200 {object} response "User registered successfully"
// @Failure     500 {object} response "Internal server error or user already exists"
// @Router      /register [post]
// @Param       request body entity.Credentials true "User credentials (e.g., username, password)"
func (rt *userRoutes) Register(w http.ResponseWriter, r *http.Request) {
	crd := entity.Credentials{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&crd)
	if err != nil {
		rt.l.Error(err, "http - v1 - register")
		errorResponse(w, http.StatusInternalServerError, "error decoding request body")
		return
	}
	if err != nil {
		rt.l.Error(err, "http - v1 - register - crd.Validate")
		errorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	status, err := rt.t.Register(r.Context(), crd)
	if err != nil || !status {
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

// @Summary     User Login
// @Description Authenticates the user by verifying credentials and returns a JWT token on success.
// @ID          login
// @Tags        Authentication
// @Accept      json
// @Produce     json
// @Success     200 {object} loginResponse "Success message and JWT token"
// @Failure     500 {object} response "Internal server error or invalid credentials"
// @Router      /login [post]
// @Param       request body entity.Credentials true "User credentials (e.g., username, password)"
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

func (r depositRequest) Validate() bool {
	return r.Amount > 0
}

type depositResponse struct {
	Status  string  `json:"status"`
	Balance float64 `json:"balance"`
}

// @Summary     Make a Deposit
// @Description Allows a user to make a deposit to their account and returns the updated balance.
// @ID          MakeDeposit
// @Security    ApiKeyAuth
// @Tags        Deposit
// @Accept      json
// @Produce     json
// @Success     200 {object} depositResponse "Deposit successful and updated balance"
// @Failure     500 {object} response "Internal server error or deposit failed"
// @Router      /deposit [post]
// @Param       request body depositRequest true "Amount to be deposited"
func (rt *userRoutes) Deposit(w http.ResponseWriter, r *http.Request) {
	req := depositRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		rt.l.Error(err, "http - v1 - Deposit")
		errorResponse(w, http.StatusInternalServerError, "error decoding request body")
		return
	}
	status := req.Validate()
	if !status {
		rt.l.Error(err, "http - v1 - Deposit - Validate")
		errorResponse(w, http.StatusInternalServerError, "amount should be positive")
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
	balance, err := rt.t.MakeDeposit(r.Context(), usr, req.Amount)
	if err != nil {
		rt.l.Error(err, "http - v1 - Deposit - rt.t.Deposit")
		errorResponse(w, http.StatusInternalServerError, "error depositing money")
		return
	}
	if balance != -1 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(depositResponse{"Deposit successful", balance})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response{"Error depositing money"})
	}
}

// @Summary     Get Current Deposit
// @Description Retrieves the current balance for the authenticated user.
// @ID          CheckDeposit
// @Security    ApiKeyAuth
// @Tags        Deposit
// @Accept      json
// @Produce     json
// @Success     200 {object} depositResponse "Current balance retrieved successfully"
// @Failure     500 {object} response "Internal server error or user not found"
// @Router      /deposit [get]
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
	json.NewEncoder(w).Encode(depositResponse{"OK", deposit})

}
