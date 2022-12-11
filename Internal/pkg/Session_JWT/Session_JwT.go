package sessionjwt

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var expirationTime time.Time

type Session struct {
	jwt.StandardClaims
	UserName string
}

type CookieHandler struct{}

func NewCookieHandler() *CookieHandler {
	return &CookieHandler{}
}

// get cookis from session object
// session object has username and jwtstandar's
func (cookieHandler *CookieHandler) CreateCookie(session *Session) (generatedTokenString string, err error) {
	expirationTime = time.Now().Add(2400 * time.Hour)

	standardClaim := jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}

	session.StandardClaims = standardClaim

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, session)

	tokenString, err := token.SignedString([]byte("JWTKEY"))
	if err != nil {
		log.Println("error while token signing")
		return "", err
	}
	// cookie = &http.Cookie{
	// 	Name:     "token",
	// 	Value:    tokenString,
	// 	Expires:  expirationTime,
	// 	HttpOnly: true,
	// }
	return tokenString, nil

}

func (cookieHandler *CookieHandler) ValidateCookie(context *gin.Context) (session *Session, isValid bool) {
	tokenCookie, err := context.Cookie("token")
	if err != nil {
		fmt.Println("erorr geting cookie from request , in session_jwt")
		return nil, false
	}
	token := tokenCookie

	session = &Session{}
	tkn, err := jwt.ParseWithClaims(token, session, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return "", nil
		}
		return []byte("JWTKEY"), nil
	})
	fmt.Println(err, "err")
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

func (cookieHandler *CookieHandler) RemoveCookie() (string, error) {
	expirationTime := time.Unix(0, 0)

	session := &Session{
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()},
		UserName:       "",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, session)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTKEY")))
	if err != nil {
		fmt.Println("error while removing cookie")
		return "", err
	}
	return tokenString, nil
}
