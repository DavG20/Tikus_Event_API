package helper

import (
	"encoding/base64"
	"time"
)

type PasswordResetData struct {
	ResetURL string    `json:"reseturl"`
	Message  string    `json:"message"`
	ExpireAt time.Time `json:"expireat"`
}

func EncodeRestPassword(s string) string {
	data := base64.StdEncoding.EncodeToString([]byte(s))
	return string(data)
}

func DecodeRestPassword(s string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	return string(data), nil

}
