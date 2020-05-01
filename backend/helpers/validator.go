package helpers

import "errors"

var strErrorValidator = "ErrValidator: "

func GenerateErrorValidator(str string) error {
	return errors.New(strErrorValidator + str)
}
