package helper

import (
	"strings"

	constants "github.com/DavG20/Tikus_Event_Api/pkg/Utils/Constants"
)

func CheckExstension(exstension string) bool {
	exstension = strings.ToLower(exstension)
	for _, ext := range constants.Extension {
		if ext == exstension {
			return true
		}
	}
	return false

}
