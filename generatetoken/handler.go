package generatetoken

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/shishanksingh2015/email-sample/model"
	"net/http"
)

type TokenHandler struct {
	Secret string
}

func (tokenHandler *TokenHandler) GenerateToken(c echo.Context) error {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Errorj(map[string]interface{}{"method": "traceId", "message": "Error while generating trace id", "error": "Please contact support", "errorCode": 500, "trace_id": ""})
		return c.JSON(http.StatusBadRequest, nil)
	}
	traceId := id.String()
	request := &Claims{}
	request.TraceId = traceId
	response := request.bind(c)
	if response.Message != "" {
		response.TraceId = traceId
		log.Errorj(map[string]interface{}{"method": "bind", "message": "Error while binding request", "error": response.Message, "errorCode": response.ErrorCode, "trace_id": traceId})
		return c.JSON(http.StatusBadRequest, response)
	}
	claims, err := json.Marshal(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ApiStatusResponse{Message: ValidationError, ErrorCode: -102, TraceId: traceId})
	}
	token, err := CreateToken(string(claims), tokenHandler.Secret)
	if err != nil {
		response.TraceId = traceId
		log.Errorj(map[string]interface{}{"method": "GenerateToken", "message": "Error while creating token", "error": err.Error(), "trace_id": traceId})
		return c.JSON(http.StatusInternalServerError, model.ApiStatusResponse{Message: SomethingWentWrong, TraceId: traceId})
	}
	return c.JSON(http.StatusOK, &JWT{Token: token, TraceId: traceId})
}
