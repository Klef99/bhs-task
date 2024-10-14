package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Klef99/bhs-task/internal/entity"
	"github.com/Klef99/bhs-task/internal/usecase"
	"github.com/Klef99/bhs-task/pkg/jwtgenerator"
	"github.com/Klef99/bhs-task/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
)

type assetRoutes struct {
	t   usecase.Asset
	l   logger.Interface
	jtg jwtgenerator.Interface
}

func NewAssetRoutes(handler chi.Router, t usecase.Asset, l logger.Interface, jtg jwtgenerator.Interface) {
	rt := &assetRoutes{t: t, l: l, jtg: jtg}
	tokenAuth := rt.jtg.GetJWTAuth()
	router := chi.NewRouter()
	router.Use(jwtauth.Verifier(tokenAuth))
	router.Use(jwtauth.Authenticator(tokenAuth))
	router.Group(func(r chi.Router) {
		r.Post("/", rt.CreateAsset)
		r.Delete("/{id}", rt.DeleteAsset)
		r.Get("/{id}", rt.GetAssetById)
		r.Get("/", rt.UserAssetsList)
		r.Get("/market", rt.AssetsToBuying)
		r.Get("/{id}/buy", rt.BuyAsset)
		r.Get("/purchased", rt.GetPurchasedAsset)
	})
	handler.Mount("/asset", router)
}

type createAssetRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

// @Summary     Create Asset
// @Description Adds a new asset to the system with the specified details.
// @ID          CreateAsset
// @Security    ApiKeyAuth
// @Tags        Asset
// @Accept      json
// @Produce     json
// @Success     200 {object} response "Asset added successfully"
// @Failure     500 {object} response "Internal server error or asset creation failed"
// @Router      /asset [post]
// @Param       request body createAssetRequest true "Asset details (name, description, price)"
func (rt *assetRoutes) CreateAsset(w http.ResponseWriter, r *http.Request) {
	car := createAssetRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&car)
	if err != nil {
		rt.l.Error(err, "http - v1 - CreateAsset")
		errorResponse(w, http.StatusInternalServerError, "error decoding request body")
		return
	}
	ast := entity.Asset{
		Name:        car.Name,
		Description: car.Description,
		Price:       car.Price,
	}
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		rt.l.Error(err, "http - v1 - CreateAsset - jwtauth.FromContext")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	f, ok := claims["id"].(float64)
	if !ok {
		rt.l.Error(err, claims["id"], "http - v1 - CreateAsset - .(int64)")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	ast.Owner_id = int64(f)
	status, err := rt.t.CreateAsset(r.Context(), ast)
	if err != nil {
		rt.l.Error(err, "http - v1 - CreateAsset")
		errorResponse(w, http.StatusInternalServerError, "error creating asset")
		return
	}
	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response{"Asset added successfully"})
	}
}

type deleteAssetRequest struct {
	Id int64 `json:"id"`
}

// @Summary     Delete Asset
// @Description Removes an asset from the system based on the provided asset ID.
// @ID          DeleteAsset
// @Security    ApiKeyAuth
// @Tags        Asset
// @Accept      json
// @Produce     json
// @Success     200 {object} response "Asset removed successfully"
// @Failure     404 {object} response "Asset not found"
// @Failure     500 {object} response "Internal server error"
// @Router      /asset/{id} [delete]
// @Param       id path int true "Asset ID to be deleted"
func (rt *assetRoutes) DeleteAsset(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	idAsset, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		rt.l.Error(err, "http - v1 - DeleteAsset")
		errorResponse(w, http.StatusInternalServerError, "error decoding request parameters")
		return
	}

	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		rt.l.Error(err, "http - v1 - DeleteAsset - jwtauth.FromContext")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	id, ok := claims["id"].(float64)
	if !ok {
		rt.l.Error(err, "http - v1 - DeleteAsset - .(float64)")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	name, ok := claims["name"].(string)
	if !ok {
		rt.l.Error(err, "http - v1 - DeleteAsset - .(string)")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	usr := entity.User{Username: name, Id: int64(id)}
	status, err := rt.t.DeleteAsset(r.Context(), usr, int64(idAsset))
	if err != nil {
		rt.l.Error(err, "http - v1 - DeleteAsset - rt.t.DeleteAsset")
		errorResponse(w, http.StatusInternalServerError, "error deleting asset")
		return
	}
	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response{"Asset remove successfully"})
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response{"Asset not found"})
	}
}

type listOfAssetResponse struct {
	Assets []entity.Asset `json:"assets"`
}

// @Summary     List User Assets
// @Description Retrieves all assets belonging to the currently authenticated user.
// @ID          MyAssets
// @Security    ApiKeyAuth
// @Tags        Asset
// @Accept      json
// @Produce     json
// @Success     200 {object} listOfAssetResponse "List of user's assets"
// @Failure     404 {object} response "No assets found"
// @Failure     500 {object} response "Internal server error"
// @Router      /asset [get]
func (rt *assetRoutes) UserAssetsList(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		rt.l.Error(err, "http - v1 - AssetsList - jwtauth.FromContext")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	id, ok := claims["id"].(float64)
	if !ok {
		rt.l.Error(err, "http - v1 - AssetsList - .(float64)")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	name, ok := claims["name"].(string)
	if !ok {
		rt.l.Error(err, "http - v1 - AssetsList - .(string)")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	usr := entity.User{Username: name, Id: int64(id)}
	assets, err := rt.t.UserAssetsList(r.Context(), usr)
	if err != nil {
		rt.l.Error(err, "http - v1 - AssetsList - rt.t.AssetsList")
		errorResponse(w, http.StatusInternalServerError, "error getting asset")
		return
	}
	if len(assets) != 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(listOfAssetResponse{assets})
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response{"Asset not found"})
	}
}

// @Summary     Get List of Assets for Buying
// @Description Retrieves all assets available for purchase in the system.
// @ID          BuyingList
// @Security    ApiKeyAuth
// @Tags        Asset
// @Accept      json
// @Produce     json
// @Success     200 {object} listOfAssetResponse "List of assets available for buying"
// @Failure     404 {object} response "No assets found"
// @Failure     500 {object} response "Internal server error"
// @Router      /asset/market [get]
func (rt *assetRoutes) AssetsToBuying(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		rt.l.Error(err, "http - v1 - AssetsList - jwtauth.FromContext")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	id, ok := claims["id"].(float64)
	if !ok {
		rt.l.Error(err, "http - v1 - AssetsList - .(float64)")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	name, ok := claims["name"].(string)
	if !ok {
		rt.l.Error(err, "http - v1 - AssetsList - .(string)")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	usr := entity.User{Username: name, Id: int64(id)}
	assets, err := rt.t.GetAssetsToBuying(r.Context(), usr)
	if err != nil {
		rt.l.Error(err, "http - v1 - AssetsList - rt.t.GetAllAssets")
		errorResponse(w, http.StatusInternalServerError, "error getting asset")
		return
	}
	if len(assets) != 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(listOfAssetResponse{assets})
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response{"Asset not found"})
	}
}

// @Summary     Buy Asset
// @Description Allows the user to purchase an asset by its ID.
// @ID          BuyAsset
// @Security    ApiKeyAuth
// @Tags        Asset
// @Accept      json
// @Produce     json
// @Success     200 {object} response "Asset purchased successfully"
// @Failure     404 {object} response "Asset not found or purchase failed"
// @Failure     500 {object} response "Internal server error"
// @Router      /asset/{id}/buy [get]
// @Param       id path int true "Asset ID to retrieve"
func (rt *assetRoutes) BuyAsset(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	idAsset, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		rt.l.Error(err, "http - v1 - GetAssetById")
		errorResponse(w, http.StatusInternalServerError, "error decoding request parameters")
		return
	}
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		rt.l.Error(err, "http - v1 - BuyAsset - jwtauth.FromContext")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	id, ok := claims["id"].(float64)
	if !ok {
		rt.l.Error(err, "http - v1 - BuyAsset - .(float64)")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	name, ok := claims["name"].(string)
	if !ok {
		rt.l.Error(err, "http - v1 - BuyAsset - .(string)")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	usr := entity.User{Username: name, Id: int64(id)}
	status, err := rt.t.BuyAsset(r.Context(), usr, idAsset)
	if err != nil {
		rt.l.Error(err, "http - v1 - BuyAsset - rt.t.BuyAsset")
		errorResponse(w, http.StatusInternalServerError, "error buying asset")
		return
	}
	if status {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response{"Asset successfully buying"})
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response{"Failed to buy asset"})
	}
}

// @Summary     Get List of Purchased Assets
// @Description Retrieves a list of all purchased assets.
// @ID          PurchasedAsset
// @Security    ApiKeyAuth
// @Tags        Asset
// @Accept      json
// @Produce     json
// @Success     200 {object} listOfAssetResponse "List of purchased assets"
// @Failure     404 {object} response "No assets found"
// @Failure     500 {object} response "Internal server error"
// @Router      /asset/purchased [get]
func (rt *assetRoutes) GetPurchasedAsset(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		rt.l.Error(err, "http - v1 - GetPurchasedAsset - jwtauth.FromContext")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	id, ok := claims["id"].(float64)
	if !ok {
		rt.l.Error(err, "http - v1 - GetPurchasedAsset - .(float64)")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	name, ok := claims["name"].(string)
	if !ok {
		rt.l.Error(err, "http - v1 - GetPurchasedAsset - .(string)")
		errorResponse(w, http.StatusInternalServerError, "error getting token claims")
		return
	}
	usr := entity.User{Username: name, Id: int64(id)}
	assets, err := rt.t.GetPurchasedAsset(r.Context(), usr)
	if err != nil {
		rt.l.Error(err, "http - v1 - GetPurchasedAsset - rt.t.GetAllAvaliableAsset")
		errorResponse(w, http.StatusInternalServerError, "error getting asset")
		return
	}
	if len(assets) != 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(listOfAssetResponse{assets})
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response{"Asset not found"})
	}
}

// @Summary     Get Asset
// @Description Get an asset from the system based on the provided asset ID.
// @ID          GetAsset
// @Security    ApiKeyAuth
// @Tags        Asset
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.Asset "Asset retrieved successfully"
// @Failure     404 {object} response "Asset not found"
// @Failure     500 {object} response "Internal server error"
// @Router      /asset/{id} [get]
// @Param       id path int true "Asset ID to retrieve"
func (rt *assetRoutes) GetAssetById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	idAsset, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		rt.l.Error(err, "http - v1 - GetAssetById")
		errorResponse(w, http.StatusInternalServerError, "error decoding request parameters")
		return
	}
	asset, err := rt.t.GetAssetById(r.Context(), int64(idAsset))
	if err != nil {
		rt.l.Error(err, "http - v1 - GetAssetById - rt.t.GetAssetById")
		errorResponse(w, http.StatusInternalServerError, "error getting asset")
		return
	}
	if asset.Id != 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(asset)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response{"Asset not found"})
	}
}
