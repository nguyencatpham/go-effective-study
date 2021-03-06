package helper

import (
	// "fmt"

	"github.com/labstack/echo"
	"github.com/nguyencatpham/go-effective-study/pkg/utl/config"
	"github.com/thoas/go-funk"
)

var appConfig *config.Configuration

func SetAppConfig(data *config.Configuration) {
	appConfig = data
}

func GetAppConfig() *config.Configuration {
	return appConfig
}

var errorDict *config.ErrorDict

func SetErrorList(data *config.ErrorDict) {
	errorDict = data
}

func GetError(status int, messageType string) error {
	var err = echo.NewHTTPError(status, "problem to perform your action")
	if errorDict != nil {
		if errorDict.ErrorList != nil && len(errorDict.ErrorList) > 0 {
			temp := funk.Find(errorDict.ErrorList, func(item config.ErrorMessage) bool {
				if item.Type == messageType {
					return true
				}
				return false
			})
			if temp != nil {
				temp := temp.(config.ErrorMessage)
				err = echo.NewHTTPError(status, temp.Text)
			}
		}

	}
	return err
}
func HandleError(message string) error {
	return echo.NewHTTPError(400, message)
}
