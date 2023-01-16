package sendmail

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MailerClient interface {
	SendViaSendGrid(sendEmailRequest SendEmailRequest) error
}

type MailerImpl struct {
	SendGridClient *sendgrid.Client
}

func New(sendGridKey string) *MailerImpl {
	sendGridClient := sendgrid.NewSendClient(sendGridKey)
	return &MailerImpl{SendGridClient: sendGridClient}
}

func (mailer *MailerImpl) SendViaSendGrid(sendEmailRequest SendEmailRequest) error {
	from := mail.NewEmail(sendEmailRequest.FromName, sendEmailRequest.From)
	to := mail.NewEmail(sendEmailRequest.ToName, sendEmailRequest.To)
	subject := sendEmailRequest.Subject
	var htmlContent = ""
	var plainTextContent = sendEmailRequest.Content

	var message = mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	//attachment is required
	if sendEmailRequest.File.Content != "" {
		a := mail.NewAttachment()
		a.SetContent(sendEmailRequest.File.Content)
		a.SetType(sendEmailRequest.File.Type)
		a.SetFilename(sendEmailRequest.File.Name)
		a.SetDisposition("attachment")
		a.SetContentID(sendEmailRequest.File.Name)
		message.AddAttachment(a)
	}
	_, err := mailer.SendGridClient.Send(message)
	return err
}
