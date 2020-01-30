package controllers

import(
	"net/http"
	"io/ioutil"
	"encoding/json"

	"bob-bank/models"
	"bob-bank/utils"
	"bob-bank/auth"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var user models.User
	err := json.Unmarshal(body, &user)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnauthorized)
		return
	}
	userAuth, err := auth.SignIn(user)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnauthorized)
		return
	}
	utils.ToJson(w, userAuth)
}