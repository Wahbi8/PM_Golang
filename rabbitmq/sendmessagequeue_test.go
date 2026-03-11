package rabbitmq

import (
	"encoding/json"
	"testing"

	"github.com/Wahbi8/PM_Golang/DTO"
	"github.com/google/uuid"
)

func TestQueueMsg(t *testing.T) {
	emailInfo := dto.EmailInfo{
		Recipient:   "test@example.com",
		Message:     "Test message body",
		Subject:     "Test Subject",
		InvoiceId:   uuid.New(),
		InvoiceType: 2,
		Retry:       0,
	}

	result := QueueMsg(emailInfo)

	if result == nil {
		t.Fatal("Expected non-nil result")
	}

	var parsed map[string]interface{}
	err := json.Unmarshal(result, &parsed)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if parsed["RecipientEmail"] != "test@example.com" {
		t.Errorf("Expected RecipientEmail to be test@example.com, got %v", parsed["RecipientEmail"])
	}

	if parsed["Body"] != "Test message body" {
		t.Errorf("Expected Body to be Test message body, got %v", parsed["Body"])
	}

	if parsed["subject"] != "Test Subject" {
		t.Errorf("Expected subject to be Test Subject, got %v", parsed["subject"])
	}

	if parsed["retry"] != float64(0) {
		t.Errorf("Expected retry to be 0, got %v", parsed["retry"])
	}
}

func TestQueueMsg_WithRetry(t *testing.T) {
	emailInfo := dto.EmailInfo{
		Recipient:   "retry@example.com",
		Message:     "Retrying message",
		Subject:     "Retry Subject",
		InvoiceId:   uuid.New(),
		InvoiceType: 1,
		Retry:       2,
	}

	result := QueueMsg(emailInfo)

	var parsed map[string]interface{}
	err := json.Unmarshal(result, &parsed)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if parsed["retry"] != float64(2) {
		t.Errorf("Expected retry to be 2, got %v", parsed["retry"])
	}
}

func TestFailOnError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Got expected panic: %v", r)
		}
	}()

	err := &testError{"test error"}
	failOnError(err, "Test panic message")
}

type testError struct {
	message string
}

func (e *testError) Error() string {
	return e.message
}
