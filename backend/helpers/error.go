package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Descriptions Used to store errors payload
type Descriptions struct {
	FR  string `json:"fr"`
	ENG string `json:"eng"`
}

type ErrorKey string

const (
	// URL : error is about bad URL
	URL ErrorKey = "URL"
	// PAYLOAD : error is about bad data payload formatting
	PAYLOAD ErrorKey = "PAYLOAD"
	// HEADER : error is in the header of the request
	HEADER ErrorKey = "HEADER"
	// DATA : generic error about bad data
	DATA ErrorKey = "DATA"
)

// ErrorData is used to describe an error and send it to frontend
type ErrorData struct {
	Title      string                   `json:"title"`
	HttpStatus int                      `json:"httpStatus"`
	Descr      Descriptions             `json:"descriptions"`
	Details    map[ErrorKey]interface{} `json:"details"`
}

// MakeEasyDetail maps errors with unique key
func MakeEasyDetail(key ErrorKey, value interface{}) map[ErrorKey]interface{} {
	return map[ErrorKey]interface{}{
		key: value,
	}
}

func CheckErrorNotFound(c *gin.Context, details map[ErrorKey]interface{}, err error) {
	if gorm.IsRecordNotFoundError(err) {
		SendError(c, ErrNotFound(), details, err)
		return
	}
	SendError(c, ErrInternalServerError(), nil, err)
}

func CheckErrorValidator(c *gin.Context, err error) {
	if strings.HasPrefix(err.Error(), strErrorValidator) {
		errStr := strings.Trim(err.Error(), strErrorValidator)
		SendError(c, ErrBadRequest(), MakeEasyDetail(PAYLOAD, errStr), err)
		return
	}
	SendError(c, ErrInternalServerError(), nil, err)
}

func NewGenErrorToError(errData ErrorData) error {
	strErr, err := json.Marshal(errData)
	if err != nil {
		panic(err)
	}
	return errors.New(string(strErr))
}

func ErrAuth() ErrorData {
	return ErrorData{
		Title:      "ErrAuth",
		HttpStatus: http.StatusBadRequest,
		Descr: Descriptions{
			FR:  "Impossible de se connecter",
			ENG: "Cannot login",
		},
	}
}

func ErrBadCredentials() ErrorData {
	return ErrorData{
		Title:      "ErrBadCredentials",
		HttpStatus: http.StatusBadRequest,
		Descr: Descriptions{
			FR:  "Mauvais login ou mot de passe",
			ENG: "Bad credentials",
		},
	}
}

func ErrInternalServerError() ErrorData {
	return ErrorData{
		Title:      "ErrInternalServerError",
		HttpStatus: http.StatusInternalServerError,
		Descr: Descriptions{
			FR:  "Erreur interne du serveur",
			ENG: "Internal server error",
		},
	}
}
func ErrDuplicate() ErrorData {
	return ErrorData{
		Title:      "ErrDuplicate",
		HttpStatus: http.StatusConflict,
		Descr: Descriptions{
			FR:  "La ressource existe déjà",
			ENG: "The existing ressource already exists",
		},
	}
}
func ErrInvalidInput() ErrorData {
	return ErrorData{
		Title:      "ErrInvalidInput",
		HttpStatus: http.StatusBadRequest,
		Descr: Descriptions{
			FR:  "Les donnees entrées sont incorrectes",
			ENG: "Input data are incorrect",
		},
	}
}
func ErrInvalidInputJSON() ErrorData {
	return ErrorData{
		Title:      "ErrInvalidInputJSON",
		HttpStatus: http.StatusBadRequest,
		Descr: Descriptions{
			FR:  "Les donnees entrées dans le JSON sont incorrectes",
			ENG: "Input data are misformatted",
		},
	}
}
func ErrInvalidInputURL() ErrorData {
	return ErrorData{
		Title:      "ErrInvalidInputURL",
		HttpStatus: http.StatusBadRequest,
		Descr: Descriptions{
			FR:  "Les donnees entrées dans l'url sont incorrectes",
			ENG: "Bad URL data",
		},
	}
}
func ErrInvalidInputCookies() ErrorData {
	return ErrorData{
		Title:      "ErrInvalidInputCookies",
		HttpStatus: http.StatusBadRequest,
		Descr: Descriptions{
			FR:  "Le cookie est incorrect",
			ENG: "Cookie are corrupted",
		},
	}
}
func ErrBadRequest() ErrorData {
	return ErrorData{
		Title:      "ErrBadRequest",
		HttpStatus: http.StatusBadRequest,
		Descr: Descriptions{
			FR:  "La requête n'est pas correcte",
			ENG: "Invalid request",
		},
	}
}
func ErrBadTokenConnection() ErrorData {
	return ErrorData{
		Title:      "ErrBadTokenConnection",
		HttpStatus: http.StatusBadRequest,
		Descr: Descriptions{
			FR:  "Le token de connexion est incorrect",
			ENG: "Invalid session",
		},
	}
}
func ErrNotFound() ErrorData {
	return ErrorData{
		Title:      "ErrNotFound",
		HttpStatus: http.StatusNotFound,
		Descr: Descriptions{
			FR:  "Impossible de trouver la ressource spécifiquement demandé",
			ENG: "Ressource not found",
		},
	}
}
func ErrUnauthorized() ErrorData {
	return ErrorData{
		Title:      "ErrUnauthorized",
		HttpStatus: http.StatusUnauthorized,
		Descr: Descriptions{
			FR:  "Vous n'êtes pas autorisés à faire cette requête",
			ENG: "Not authorized",
		},
	}
}
func ErrForbidden() ErrorData {
	return ErrorData{
		Title:      "ErrForbidden",
		HttpStatus: http.StatusForbidden,
		Descr: Descriptions{
			FR:  "Il est interdit de faire cette requête dans cet état",
			ENG: "Request can't be completed in this state",
		},
	}
}

func ErrBadPassword() ErrorData {
	return ErrorData{
		Title:      "ErrBadPassword",
		HttpStatus: http.StatusBadRequest,
		Descr: Descriptions{
			FR:  "Le mot de passe renseigné est incorrect",
			ENG: "Bad password",
		},
	}
}

func ErrPasswordStrength() ErrorData {
	return ErrorData{
		Title:      "ErrPasswordStrength",
		HttpStatus: http.StatusBadRequest,
		Descr: Descriptions{
			FR:  "Le mot de passe renseigné ne respect pas les règles de sécurité minimales",
			ENG: "Bad password strength",
		},
	}
}

func ErrInvalidInputImage() ErrorData {
	return ErrorData{
		Title:      "ErrInvalidInputImage",
		HttpStatus: http.StatusBadRequest,
		Descr: Descriptions{
			FR:  "L'image envoyée est incorrecte ou manquante",
			ENG: "Bad image",
		},
	}
}
