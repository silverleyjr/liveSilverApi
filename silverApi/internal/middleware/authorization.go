package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"silverApi/api"
	"silverApi/internal/tools"

	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token.")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)
			return
		}

		//var username string = r.URL.Query().Get("username")

		username := r.Header.Get("Username")
		var token = r.Header.Get("Authorization")

		//token := "SilverleyFodao"
		fmt.Println("token=" + token)
		fmt.Println("user=" + username)
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		fmt.Println("depois de if not null inicial")

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.User
		loginDetails = (*database).GetAllUserDetails(username)

		fmt.Println("antes do if informaçao valida")
		if loginDetails == nil || (token != (*loginDetails).Authorization) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
