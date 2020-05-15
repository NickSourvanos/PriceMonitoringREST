package auth

import (
	"PriceMonitoringService/models"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"net/http"
	"strings"
	"time"
)

const (
	AppKey = "golangcode.com"
)



func TokenHandler(userId string, username, role string) (accessToken string) {
	// We are happy with the credentials, so build a token. We've given it
	// an expiry of 1 hour.
	claims := &models.Claims{
		UserId: userId,
		Username: username,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}

	fmt.Println("Claims: ", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(AppKey))
	if err != nil {
		return ""
	}

	return tokenString
}

// AuthMiddleware is our middleware to check our token is valid. Returning
// a 401 status to the client if it is not valid.
func ValidateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("authorization")
		claims := &models.Claims{}
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.ParseWithClaims(bearerToken[1], claims, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(AppKey), nil
				})
				fmt.Println("token: ", token)

				if error != nil {
					json.NewEncoder(w).Encode(models.ResponseMessage{Message: error.Error()})
					return
				}
				if token.Valid {
					//user := &models.Claims{
					//	UserId: claims.UserId,
					//	Username: claims.Username,
					//	Role: claims.Role,
					//	StandardClaims: claims.StandardClaims,
					//}

					test, _ := StructToMap(claims)

					context.Set(req, "decoded", test)
					next(w, req)
				} else {
					json.NewEncoder(w).Encode(models.ResponseMessage{Message: "Invalid authorization token"})
				}
			} else {
				json.NewEncoder(w).Encode(models.ResponseMessage{Message: "Invalid authorization token"})
			}
		} else {
			json.NewEncoder(w).Encode(models.ResponseMessage{Message: "An authorization header is required"})
		}
	})
}

func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj) // Convert to a json string

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap) // Convert to a map
	return
}