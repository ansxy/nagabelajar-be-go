package custom_validator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/ansxy/nagabelajar-be-go/pkg/constant"
	custom_error "github.com/ansxy/nagabelajar-be-go/pkg/error"

	go_validator "github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *go_validator.Validate
}

func NewValidator(v *go_validator.Validate) Validator {
	return Validator{
		validate: v,
	}
}

func (v *Validator) ValidateStruct(r *http.Request, values interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, values)
	if err != nil {
		fmt.Printf("[Error Parse Body] %v\n", err.Error())
		err = custom_error.SetCostumeError(&custom_error.ErrorContext{
			Code:     constant.DefaultValidationError,
			Message:  fmt.Sprintf("Error Parse Body: %v", err.Error()),
			HTTPCode: constant.ErrorCodeResponseMap[constant.DefaultValidationError],
		})
		return err
	}

	err = v.validate.Struct(values)

	if err != nil {
		var missingFields []string

		for _, err := range err.(go_validator.ValidationErrors) {
			field := err.Field()
			missingFields = append(missingFields, field)
		}

		message := fmt.Sprintf("fields %v are required", strings.Join(missingFields, ", "))
		err := custom_error.SetCostumeError(&custom_error.ErrorContext{
			Code:     constant.DefaultValidationError,
			Message:  message,
			HTTPCode: constant.ErrorCodeResponseMap[constant.DefaultValidationError],
		})
		return err
	}

	return nil
}
