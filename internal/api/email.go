package api

import (
	"context"
	"ewallet-notification/cmd/proto/notification"
	"ewallet-notification/helpers"
	"ewallet-notification/internal/interfaces"
	"ewallet-notification/internal/models"
)

type EmailAPI struct {
	EmailService interfaces.IEmailService
	notification.UnimplementedNotificationServiceServer
}

func (api *EmailAPI) SendNotification(ctx context.Context, req *notification.SendNotificationRequest) (*notification.SendNotificationResponse, error) {
	var (
		log = helpers.Logger
	)
	internalReq := models.InternalNotificationRequest{
		TemplateName: req.TemplateName,
		Recipient:    req.Recipient,
		Placeholder:  req.Placeholders,
	}

	if err := internalReq.Validate(); err != nil {
		log.Error("failed to validate request: ", err)
		return &notification.SendNotificationResponse{
			Message: "failed to validate request",
		}, nil
	}

	err := api.EmailService.SendEmail(ctx, internalReq)
	if err != nil {
		log.Error("failed to send email: ", err)
		return &notification.SendNotificationResponse{
			Message: "failed to send email",
		}, nil
	}

	return &notification.SendNotificationResponse{
		Message: "success",
	}, nil
}
