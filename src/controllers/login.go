package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repos"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repos.NewUserRepo(db)
	dbUser, err := repo.SearchByEmail(user.Email)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.CheckPass(dbUser.Password, user.Password); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, err := authentication.CreateToken(dbUser.ID, dbUser.Role)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	tokenResp := struct {
		Token string `json:"token"`
	}{token}

	responses.JSON(w, http.StatusOK, tokenResp)

}

func ForgotPass(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	resp := struct {
		EmailMessage string `json:"message"`
	}{fmt.Sprintf("Caso o e-mail %s esteja cadastrado, você receberá uma mensagem de recuperação.", user.Email)}

	responses.JSON(w, http.StatusOK, resp)
}
