package sendmail

import (
	"github.com/labstack/echo"
	"github.com/shishanksingh2015/email-sample/model"
	"gopkg.in/go-playground/validator.v9"
)

type (
	File struct {
		Type    string `json:"type"`
		Content string `json:"content"`
		Name    string `json:"name"`
	}

	SendEmailRequest struct {
		From     string `json:"from" validate:"required,email" form:"from"`
		FromName string `json:"from_name" form:"from_name"`
		To       string `json:"to" form:"to" validate:"required,email"`
		ToName   string `json:"to_name" form:"to_name"`
		Subject  string `json:"subject" form:"subject" validate:"required"`
		Content  string `json:"content" form:"content"`
		File     File   `json:"file"`
	}
	VarValues struct {
		Variable string `json:"variable"`
		Value    string `json:"value"`
	}
)

func (request *SendEmailRequest) bind(c echo.Context) *model.ApiStatusResponse {
	if err := c.Bind(request); err != nil {
		return &model.ApiStatusResponse{ErrorCode: -101,
			Message: TypeError}
	}
	v := validator.New()
	if err := v.Struct(request); err != nil {
		return &model.ApiStatusResponse{ErrorCode: -102,
			Message: ValidationError}
	}
	return &model.ApiStatusResponse{}
}
