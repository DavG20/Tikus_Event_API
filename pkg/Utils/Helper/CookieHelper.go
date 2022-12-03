package helper

import "github.com/gin-gonic/gin"

func SetCookie(tokenString string, c *gin.Context) {
	c.SetCookie(
		"token",
		tokenString,
		3600,
		"/user",
		"localhost",
		false,
		true,
	)
}

func RemoveCookie(c *gin.Context) {
	c.SetCookie(
		"token",
		"",
		-24,
		"/user",
		"localhost",
		false,
		true,
	)

}
