package base

import (
	"errors"
	"testing"
)

func newServiceWithSpy(t *testing.T) (*NotificationService, *CountSender) {
	t.Helper()
	spy := &CountSender{}
	svc, err := NewNotificationService(spy)
	if err != nil {
		t.Fatalf("неожидали ошибку конструктора: %v", err)
	}
	return svc, spy
}

func TestNotify_MemorySender_Success(t *testing.T) {
    t.Parallel()
	sender := &MemorySender{}
	svc, err := NewNotificationService(sender)
	if err != nil {
		t.Fatalf("неожидали ошибку конструктора: %v", err)
	}

	if err := svc.Notify("alice@example.com", "hello"); err != nil {
		t.Fatalf("неожидали ошибку Notify: %v", err)
	}

	if len(sender.Messages) != 1 {
		t.Fatalf("ожидали 1 сообщение, получили %d", len(sender.Messages))
	}
	if sender.Messages[0].Recipient != "alice@example.com" {
		t.Errorf("recipient: ожидали %q, получили %q",
			"alice@example.com", sender.Messages[0].Recipient)
	}
	if sender.Messages[0].Text != "hello" {
		t.Errorf("text: ожидали %q, получили %q", "hello", sender.Messages[0].Text)
	}
}

func TestNotify_PassesArgumentsAndCallsOnce(t *testing.T) {
    t.Parallel()
	svc, spy := newServiceWithSpy(t)

	if err := svc.Notify("bob@example.com", "hi"); err != nil {
		t.Fatalf("неожидали ошибку Notify: %v", err)
	}

	if spy.Calls != 1 {
		t.Fatalf("ожидали 1 вызов Send, получили %d", spy.Calls)
	}
	if spy.recipient != "bob@example.com" {
		t.Errorf("recipient: ожидали %q, получили %q", "bob@example.com", spy.recipient)
	}
	if spy.text != "hi" {
		t.Errorf("text: ожидали %q, получили %q", "hi", spy.text)
	}
}

func TestNotify_EmptyRecipient_NoSend(t *testing.T) {
    t.Parallel()
	svc, spy := newServiceWithSpy(t)

	if err := svc.Notify("", "hello"); err == nil {
		t.Fatal("ожидали ошибку валидации для пустого recipient")
	}

	if spy.Calls != 0 {
		t.Fatalf("Send не должен вызываться, Calls = %d", spy.Calls)
	}
}

func TestNotify_EmptyText_NoSend(t *testing.T) {
    t.Parallel()
	svc, spy := newServiceWithSpy(t)

	if err := svc.Notify("alice@example.com", ""); err == nil {
		t.Fatal("ожидали ошибку валидации для пустого text")
	}

	if spy.Calls != 0 {
		t.Fatalf("Send не должен вызываться, Calls = %d", spy.Calls)
	}
}

func TestNotify_SendError_ReturnedAsIs(t *testing.T) {
    t.Parallel()
	sentinel := errors.New("boom")
	spy := &CountSender{err: sentinel}
	svc, err := NewNotificationService(spy)
	if err != nil {
		t.Fatalf("неожидали ошибку конструктора: %v", err)
	}

	got := svc.Notify("alice@example.com", "hello")

	if got != sentinel {
		t.Fatalf("ожидали ту же ошибку %v, получили %v", sentinel, got)
	}
	if spy.Calls != 1 {
		t.Fatalf("ожидали 1 вызов Send, получили %d", spy.Calls)
	}
}

func TestNewNotificationService_NilSender(t *testing.T) {
    t.Parallel()
	svc, err := NewNotificationService(nil)
	if err == nil {
		t.Fatal("ожидали ошибку для nil-интерфейса")
	}
	if svc != nil {
		t.Fatalf("ожидали nil сервис, получили %v", svc)
	}
}

func TestTypedNil_PassesNilCheck(t *testing.T) {
    t.Parallel()
	var ms *MemorySender = nil
	var sender Sender = ms

	if sender == nil {
		t.Fatal("ожидали, что typed nil НЕ равен nil-интерфейсу")
	}

	svc, err := NewNotificationService(sender)
	if err != nil {
		t.Fatalf("ожидали, что typed nil пройдёт проверку, получили ошибку: %v", err)
	}
	if svc == nil {
		t.Fatal("ожидали непустой сервис")
	}
}