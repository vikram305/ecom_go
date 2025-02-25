package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vikram305/ecom/services/auth"
	"github.com/vikram305/ecom/types"
	"github.com/vikram305/ecom/utils"
)
type Handler struct{
    store types.UserStore
}

func NewHandler (store types.UserStore) *Handler{
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router){
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/login", h.handleRegister).Methods("POST")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request){

}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request){
	// get payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err!=nil{
		utils.WriteJSON(w, http.StatusBadRequest,err)
	}

	// check if user already exist
	_, err := h.store.GetUserByEmail(payload.Email)
	if err!=nil{
		utils.WriteJSON(w, http.StatusBadRequest, fmt.Errorf("User with email %s already exist", payload.Email))
		return
	}

	
	hashedPassword,err := auth.HashPassword(payload.Password)
	if err!=nil{
		utils.WriteError(w, http.StatusInternalServerError,err)
		return
	}

	
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email: payload.Email,
		Password: hashedPassword,
	})


	if err!=nil{
		utils.WriteError(w, http.StatusInternalServerError,err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated,nil)
	
}

