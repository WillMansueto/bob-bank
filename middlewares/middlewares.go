package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"bob-bank/config"
	
	jwt "github.com/dgrijalva/jwt-go"
)

var secretkey = config.LoadConfigs().Jwt.SecretKey

func IsAuth(handler func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header != "" {
			bearerToken := strings.Split(header, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token)(interface{}, error) {
					_, ok := token.Method.(*jwt.SigningMethodHMAC)
					if !ok {
						return nil, fmt.Errorf("Falha de autenticação")
					}
					return secretkey, nil
				})
				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("Unauthorized"))
					return
				}
				if token.Valid {
					handler(w, r)
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized")
		}
	})
}
