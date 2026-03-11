package Services

import (
	"testing"
)

type MockResendClient struct {
	SentEmails []MockEmail
}

type MockEmail struct {
	To      []string
	From    string
	Html    string
	Subject string
}

func (m *MockResendClient) Send(req interface{}) (interface{}, error) {
	return map[string]string{"id": "test-123"}, nil
}

func TestSendEmail_Success(t *testing.T) {
	t.Skip("Requires actual Resend API key - run with real credentials")

	err := SendEmail("test@example.com", "Test Subject", "Test Body")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestSendEmail_MissingAPIKey(t *testing.T) {
	t.Setenv("Resend_api_key", "")

	err := SendEmail("test@example.com", "Test Subject", "Test Body")
	if err == nil {
		t.Error("Expected error when API key is missing")
	}
}
