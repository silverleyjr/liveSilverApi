package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"silverApi/api"
	"silverApi/internal/tools"

	log "github.com/sirupsen/logrus"
)

func AuthorizationOptions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	fmt.Println("authorization")
}

func GetPasswordDatabase(w http.ResponseWriter, r *http.Request) {
	//var decoder *schema.Decoder = schema.NewDecoder()
	//err = decoder.Decode(&params, r.URL.Query())

	var err error
	user := r.Header.Get("AuthToken")
	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	fmt.Println("user recebido:" + user)
	var tokenDetails *tools.LoginDetails
	tokenDetails = (*database).GetUserLoginDetails(user)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.PasswordResponse{
		Password: (*tokenDetails).Password,
		Code:     http.StatusOK,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	var err error

	name := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")
	authorization := r.URL.Query().Get("authorization")

	if name == "" || password == "" || authorization == "" {
		api.RequestErrorHandler(w, errors.New("Not mactch params"))
	} else {
		newUser := tools.LoginDetails{Authorization: name, Password: password}

		w.Header().Set("Content-Type", "application/json")
		err = tools.NewUser(newUser)
		if err != nil {
			api.RequestErrorHandler(w, err)
		} else {
			response := api.NewUserResponse{
				Code: http.StatusOK,
			}
			err = json.NewEncoder(w).Encode(response)
			if err != nil {
				log.Error(err)
				api.InternalErrorHandler(w)
				return
			}
		}
	}
}
