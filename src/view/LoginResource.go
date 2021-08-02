package view

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sebastianbordi/DataloggerDemo/configuration"
	"github.com/sebastianbordi/DataloggerDemo/controller"
	"github.com/sebastianbordi/DataloggerDemo/security"
)

type userDto struct {
	User     string
	Password string
}

func Login(w http.ResponseWriter, r *http.Request) {
	var userDto userDto
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		log.Println(err)
		BasicResponse(&w, 400, "ivalid body or error decoding json body")
		return
	}
	if userDto.Password == "" || userDto.User == "" {
		log.Println(err)
		BasicResponse(&w, 400, "user or password  empty")
		return
	}
	userController := controller.GetUserController()
	userDb, err := userController.GetUserByName(userDto.User)
	if err != nil {
		message := ""
		statusCode := 0
		if err.Error() == "record not found" {
			message = "user or password incorrect"
			statusCode = 401
		} else {
			message = "internal server error"
			statusCode = 500
		}
		log.Println(err)
		BasicResponse(&w, statusCode, message)
		return
	}
	if userDb.Password != userDto.Password {
		BasicResponse(&w, 401, "user or password incorrect")
		return
	}
	token, err := security.GetToken(*userDb)
	if err != nil {
		log.Println(err)
		BasicResponse(&w, 500, "internal server error")
		return
	}
	w.Header().Add("Authorization", token)
}

func JWTAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			BasicResponse(&rw, 401, "token not found")
			return
		}
		conf, err := configuration.GetInstance()
		if err != nil {
			log.Println(err)
			BasicResponse(&rw, 500, "Internal server error")
			return
		}
		tokenKey := conf.GetTokenKey()

		valid, _, err := security.ValidateToken(authorization[7:], tokenKey)
		if err != nil {
			log.Println(err)
			BasicResponse(&rw, 500, "Internal server error")
			return
		}
		if !valid {
			log.Printf("login atemp with invalid token %s", authorization)
			BasicResponse(&rw, 401, "Invalid token")
			return
		}
		next(rw, r)
	}
}
