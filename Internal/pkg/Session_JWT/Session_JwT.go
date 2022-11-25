package sessionjwt

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Session struct {
	jwt.StandardClaims
	UserName string
}

type CookieHandler struct{}

func NewCookieHandler() *CookieHandler {
	return &CookieHandler{}
}

func (cookieHandler *CookieHandler) CreateCookie(session *Session) (cookie *http.Cookie, err error) {
	expirationTime := time.Now().Add(2400 * time.Hour)

	standardClaim := jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}

	session.StandardClaims = standardClaim

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, session)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTKEY")))
	if err != nil {
		log.Println("error while token signing")
		return nil, err
	}
	cookie = &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	}
	return cookie, nil

}

func (cookieHandler *CookieHandler) ValidateCookie(request *http.Request) (session *Session, isValid bool) {
	tokenCookie, err := request.Cookie("token")
	if err != nil {
		fmt.Println("erorr geting cookie from request , in session_jwt")
		return nil, false
	}
	token := tokenCookie.Value
	fmt.Println(token, "token values")

	tkn, err := jwt.ParseWithClaims(token, session, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return "", nil
		}
		return []byte(os.Getenv("JWTKEY")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, false
		}
		return nil, false
	}
	if !tkn.Valid {
		return nil, false
	}
	return session, true

}

func (cookieHandler *CookieHandler) RemoveCookie() (*http.Cookie, error) {
	expirationTime := time.Unix(0, 0)

	session := &Session{
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()},
		UserName:       "",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, session)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTKEY")))
	if err != nil {
		fmt.Println("error while removing cookie")
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  expirationTime,
		HttpOnly: true,
	}
	return cookie, nil
}
