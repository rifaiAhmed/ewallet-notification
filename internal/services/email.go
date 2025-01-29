package services

import (
	"bytes"
	"context"
	"ewallet-notification/external"
	"ewallet-notification/internal/interfaces"
	"ewallet-notification/internal/models"
	"html/template"

	"github.com/pkg/errors"
)

type EmailService struct {
	EmailRepo interfaces.IEmailRepo
}

func (s *EmailService) SendEmail(ctx context.Context, req models.InternalNotificationRequest) error {
	emailTemplate, err := s.EmailRepo.GetTemplate(ctx, req.TemplateName)
	if err != nil {
		return errors.Wrap(err, "failed to get template email")
	}

	tmpl, err := template.New("emailTemplate").Parse(emailTemplate.Body)
	if err != nil {
		return errors.Wrap(err, "failed to parse email template")
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, req.Placeholder)
	if err != nil {
		return errors.Wrap(err, "failed to execute the placeholder")
	}

	email := external.Email{
		To:      req.Recipient,
		Subject: emailTemplate.Subject,
		Body:    tpl.String(),
	}
	err = email.SendEmail()
	if err != nil {
		notifHistory := &models.NotificationHistory{
			Recipient:    req.Recipient,
			TemplateID:   emailTemplate.ID,
			Status:       "failed",
			ErrorMessage: err.Error(),
		}
		s.EmailRepo.InsertNotificationHistory(ctx, notifHistory)

		return errors.Wrap(err, "failed to send email")
	}

	notifHistory := &models.NotificationHistory{
		Recipient:  req.Recipient,
		TemplateID: emailTemplate.ID,
		Status:     "success",
	}
	s.EmailRepo.InsertNotificationHistory(ctx, notifHistory)

	return nil
}
