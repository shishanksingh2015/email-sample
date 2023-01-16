package generatetoken

import (
	"github.com/labstack/echo"
	"github.com/shishanksingh2015/email-sample/model"
	"gopkg.in/go-playground/validator.v9"
)

type Claims struct {
	RequestorEmail string `json:"requestorEmail" validate:"required,email"`
	TraceId        string `json:"traceId"`
}

func (c *Claims) bind(echo echo.Context) *model.ApiStatusResponse {
	if err := echo.Bind(c); err != nil {
		return &model.ApiStatusResponse{ErrorCode: -101,
			Message: TypeError}
	}
	v := validator.New()
	if err := v.Struct(c); err != nil {
		return &model.ApiStatusResponse{ErrorCode: -102,
			Message: ValidationError}
	}
	return &model.ApiStatusResponse{}
}
