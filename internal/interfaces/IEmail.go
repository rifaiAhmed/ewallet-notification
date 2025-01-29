package interfaces

import (
	"context"
	"ewallet-notification/cmd/proto/notification"
	"ewallet-notification/internal/models"
)

type IEmailExternal interface {
	SendEmail() error
}

type IEmailRepo interface {
	GetTemplate(ctx context.Context, templateName string) (models.NotificationTemplate, error)
	InsertNotificationHistory(ctx context.Context, notif *models.NotificationHistory) error
}

type IEmailService interface {
	SendEmail(ctx context.Context, req models.InternalNotificationRequest) error
}

type IEmailAPI interface {
	SendNotification(ctx context.Context, req *notification.SendNotificationRequest) (*notification.SendNotificationResponse, error)
}
