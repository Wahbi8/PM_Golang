package apis

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// 1. Create a "Mock" sender that doesn't actually send emails
type MockEmailSender struct {
	WasCalled bool
	SentTo    string
}

func (m *MockEmailSender) Send(recipient, subject, body string) error {
	m.WasCalled = true
	m.SentTo = recipient
	return nil // Simulate success
}

func TestSendEmailApi(t *testing.T) {
	// 2. Setup the API with the Mock sender
	mockSender := &MockEmailSender{}
	api := &EmailAPI{Sender: mockSender}

	// 3. Create a fake JSON request (like the one from C#)
	requestBody, _ := json.Marshal(map[string]string{
		"RecipientEmail": "test@example.com",
		"Subject":        "Test Invoice",
		"Body":           "Hello World",
	})

	// 4. Simulate an HTTP request
	req := httptest.NewRequest("POST", "/email/invoice", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// 5. Create a response recorder (to "catch" the output)
	w := httptest.NewRecorder()

	// 6. Run the function
	api.SendEmailApi(w, req)

	// 7. Assertions (Check if it worked)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	if !mockSender.WasCalled {
		t.Error("The email sender was never called!")
	}

	if mockSender.SentTo != "test@example.com" {
		t.Errorf("Expected recipient test@example.com, got %s", mockSender.SentTo)
	}
}