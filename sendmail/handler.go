package sendmail

import (
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/shishanksingh2015/email-sample/model"
	"net/http"
)

type Handler struct {
	Mailer MailerClient
}

func (handler *Handler) SendEmail(c echo.Context) error {
	id, err := uuid.NewUUID()
	if err != nil {
		log.Errorj(map[string]interface{}{"method": "traceId", "message": "Error while generating trace id", "error": "Please contact support", "errorCode": 500, "trace_id": ""})
		return c.JSON(http.StatusBadRequest, nil)
	}
	traceId := id.String()
	sendEmailRequest := SendEmailRequest{}
	response := sendEmailRequest.bind(c)
	if response.Message != "" {
		response.TraceId = traceId
		log.Errorj(map[string]interface{}{"method": "bind", "message": "Error while binding request", "error": response.Message, "errorCode": response.ErrorCode, "trace_id": traceId})
		return c.JSON(http.StatusBadRequest, response)
	}
	if sendEmailRequest.FromName == "" {
		sendEmailRequest.FromName = sendEmailRequest.From
	}
	if sendEmailRequest.ToName == "" {
		sendEmailRequest.ToName = sendEmailRequest.To
	}

	err = handler.Mailer.SendViaSendGrid(sendEmailRequest)
	if err != nil {
		log.Errorj(map[string]interface{}{"method": "sendEmail", "message": "Error while sending the email", "error": err.Error(), "trace_id": traceId})
		return c.JSON(http.StatusInternalServerError, model.ApiStatusResponse{ErrorCode: -103, Message: UnableToSendEmail, TraceId: traceId})
	}
	return c.JSON(http.StatusOK, model.ApiStatusResponse{ErrorCode: 0, Message: EmailSent, TraceId: traceId})
}
