package helper

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
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

func SaveProfileInFileSystem(file multipart.File, userName, eventId string) string {
	profilePath := "../../pkg/Entity/Static/EventProfile/" + userName + "_" + eventId + ".png"
	filePath, err := os.Create(profilePath)
	if err != nil {
		fmt.Println("error in helper profile")
		return ""
	}
	defer filePath.Close()
	_, err = io.Copy(filePath, file)
	if err != nil {
		return ""
	}
	return profilePath

}
