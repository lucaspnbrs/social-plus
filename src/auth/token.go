package auth

import (
	"errors"
	"fmt"
	"net/http"
	"social-plus/src/config"
	"strconv"
	"strings"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

//Values it receives as a parameter and returns with strings and errors
func CreateToken(userID uint64) (string, error){
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}

//Verify if the token is valid!
func ValidateToken(r *http.Request) error {
	 tokenString := extractToken(r)
	 token, erro := jwt.Parse(tokenString, returnKeyWithVerification)

	 if erro != nil {
		return erro
	 }

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil 
	}

	return errors.New("Invalid Token")
	
}


func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnKeyWithVerification(token *jwt.Token) (interface {}, error) {
	if _ , ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected Subscription Method! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

func ExtractUserWithID(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, erro := jwt.Parse(tokenString, returnKeyWithVerification)
	if erro != nil {
		return 0, erro
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userID"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return userID, nil
	}

	return 0, errors.New("Invalid Token")
}