package example1

import (
	"errors"
	"mocking/example1/messaging_service"
)

var ErrNotify = errors.New("failed to notify")

func SendNotification(ms messaging_service.MessagingService, message string, recipient string) error {
	return ms.SendMessage(message, recipient)
}

type Notifier interface {
	Notify(message string, recipient string) error
}

type notifier struct {
	ms messaging_service.MessagingService
}

func NewNotifier(msgSvr messaging_service.MessagingService) *notifier {
	return &notifier{ms: msgSvr}
}

func (n notifier) Notify(message string, recipient string) error {
	if err := n.ms.SendMessage(message, recipient); err != nil {
		return ErrNotify
	}
	return nil
}
