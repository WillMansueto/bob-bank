package controllers

import(
	"net/http"
	"strconv"
	"io/ioutil"
	"encoding/json"

	"github.com/gorilla/mux"

	"bob-bank/utils"
	"bob-bank/models"
	"bob-bank/validations"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["uid"], 10, 32)
	user, err := models.GetUser(uint32(id))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJson(w, user)
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil{
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	user, err = validations.ValidateNewUser(user)
	if err != nil{
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	_, err = models.NewUser(user)
	if err != nil{
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, utils.DefaultResponse{"Usuário criado com sucesso", http.StatusCreated})
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uid, _ := strconv.ParseUint(params["uid"], 10, 32)
	body, _ := ioutil.ReadAll(r.Body)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	user.UID = uint32(uid)
	rows, err := models.UpdateUser(user)	
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rows)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uid, _ := strconv.ParseUint(params["uid"], 10, 32)
	_, err := models.DeleteUser(uint32(uid))	
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}