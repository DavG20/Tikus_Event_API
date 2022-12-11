package helper

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type PasswordResetData struct {
	jwt.StandardClaims
	Email string `json:"email"`
}

func GenerateResetCode(passReset PasswordResetData) string {
	expireAt := time.Now().Add(30 * time.Minute)
	standardClaim := jwt.StandardClaims{ExpiresAt: expireAt.Unix(), Issuer: "Negarit"}
	passReset.StandardClaims = standardClaim

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, passReset)
	fmt.Println(token)
return ""
}
