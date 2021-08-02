package security

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sebastianbordi/DataloggerDemo/configuration"
	"github.com/sebastianbordi/DataloggerDemo/model"
)

func ValidateToken(tokenString, tokenKey string) (bool, *jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		key := []byte(tokenKey)
		return key, nil
	})
	return token.Valid, token, err
}

func GetToken(user model.User) (string, error) {
	config, err := configuration.GetInstance()
	if err != nil {
		return "", err
	}
	token, err := generataJwt(user, config.GetTokenKey())

	return token, err
}

func generataJwt(user model.User, tokenKey string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usr": user.User,
		"iss": "Emconsol",
		"web": "www.emconsol.com",
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Duration(50) * time.Minute).Unix(),
	})

	signedToken, err := token.SignedString([]byte(tokenKey))

	return "Barear " + signedToken, err
}
