package base

import "errors"

type Sender interface {
    Send(recipient, text string) error
}

type NotificationService struct {
    sender Sender
}

func NewNotificationService(sender Sender) (*NotificationService, error) {
    if sender == nil {
        return nil, errors.New("The sender must not be nil")
    }

    return &NotificationService{sender: sender}, nil
}

func (s *NotificationService) Notify(recipient, text string) error {
    if (recipient == "") {
        return errors.New("Recipient must not be empty")
    }

    if (text == "") {
        return errors.New("Text must not be empty")
    }
    return s.sender.Send(recipient, text)
}

type Message struct {
    Recipient string
    Text      string
}

type MemorySender struct {
    Messages []Message
    Err      error
}

func (s *MemorySender) Send(recipient, text string) error {
    if s.Err != nil {
        return s.Err
    }
    s.Messages = append(s.Messages, Message{Recipient: recipient, Text: text})
    return nil
}

type CountSender struct {
    Calls     int
    recipient string
    text      string
    err       error   // если задана — Send вернёт её
}

func (s *CountSender) Send(recipient, text string) error {
    s.Calls++
    s.recipient = recipient
    s.text = text
    return s.err
}

var _ Sender = (*MemorySender)(nil)
var _ Sender = (*CountSender)(nil)