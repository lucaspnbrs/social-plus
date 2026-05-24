package auth

import (
	"social-plus/src/config"
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
