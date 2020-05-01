package helpers

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
)

func prepareData(data interface{}) gin.H {
	return gin.H{"data": data}
}

// SendData oui
func SendData(c *gin.Context, status int, data interface{}) {
	c.JSON(status, prepareData(data))
}

// SendError Permet de renvoyer au front une erreur
func SendError(c *gin.Context, errorNewGen ErrorData, details map[ErrorKey]interface{}, errorToLog error) {
	logError(errorToLog)

	errorNewGen.Details = details

	strErr, err := json.Marshal(errorNewGen)
	if err != nil {
		logError(err)
		c.Error(err)
		panic(err)
	}
	c.AbortWithStatusJSON(errorNewGen.HttpStatus, errorNewGen)
	c.Error(errors.New(string(strErr)))
}
