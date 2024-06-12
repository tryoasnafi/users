package validation

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ValidationErrorResponse struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		errors := err.(validator.ValidationErrors)
		out := make([]ValidationErrorResponse, len(errors))
		for i, fe := range errors {
			out[i] = ValidationErrorResponse{fe.Field(), msgForTag(fe.Tag())}
		}
		return echo.NewHTTPError(http.StatusBadRequest, map[string]any{
			"errors": out,
		})
	}
	return nil
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return ""
}
