package authController

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"

	"strings"

	"../../models/tokenModel"
	. "../../models/userModel"
	. "../../respondFormating"
)

func CreateTokenEndPoint(w http.ResponseWriter, r *http.Request) {
	var user UserModel
	_ = json.NewDecoder(r.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"username": user.Username,
		"password": user.Password,
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error generating token")
	}
	json.NewEncoder(w).Encode(tokenModel.Token{Token: tokenString})
}

func ProtectedEndPoint(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	token, _ := jwt.Parse(params["token"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		var user UserModel
		mapstructure.Decode(claims, &user)
		json.NewEncoder(w).Encode(user)
	} else {
		RespondWithError(w, http.StatusUnauthorized, "Invalid authorization Token")
	}
}

// nie rozumiem tego :/
func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("there was an error")
					}
					return []byte("secret"), nil
				})
				if err != nil {
					RespondWithError(w, http.StatusInternalServerError, err.Error())
					return
				}
				if token.Valid {
					context.Set(req, "decode", token.Claims)
					next(w, req)
					RespondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
				} else {
					RespondWithError(w, http.StatusUnauthorized, "Invalid authorization token")
				}
			}
		} else {
			RespondWithError(w, http.StatusBadRequest, "An authorization header is required")
		}
	})
}

func TestEndpoint(w http.ResponseWriter, req *http.Request){
	decode := context.Get(req, "decode")
	var user UserModel
	err := mapstructure.Decode(decode.(jwt.MapClaims), &user)
	if err != nil {
		fmt.Println(err)
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println(err)
	}
}