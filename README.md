# PM_Golang - Email Service API

A robust email service API built with Go that handles sending emails with RabbitMQ message queuing, failure tracking, and automatic retry mechanisms.

## Overview

PM_Golang is a production-ready email service that provides:
- **REST API** for email submission
- **RabbitMQ Integration** for asynchronous email processing
- **Database Layer** for tracking failed emails
- **Automatic Retry Logic** for failed email delivery
- **Structured Logging** using zerolog
- **Health Checks** for monitoring

## Features

- 📧 Email sending via Resend API
- 🔄 Asynchronous processing with RabbitMQ
- 🔁 Automatic retry mechanism for failed emails
- 📊 Failed email tracking and storage
- 🩺 Health check endpoint
- 📝 Structured logging
- ✅ Comprehensive test coverage

## Prerequisites

- Go 1.23.4 or higher
- RabbitMQ running locally (or configured host)
- Database setup (configured in infrastructure)
- Resend API key

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Wahbi8/PM_Golang.git
cd PM_Golang
```

2. Install dependencies:
```bash
go mod download
go mod tidy
```

3. Configure environment variables:
Create a `.env` file in the project root:
```
Resend_api_key=your_resend_api_key_here
```

## Quick Start

1. **Start RabbitMQ** (if not already running):
```bash
rabbitmq-server
```

2. **Run the application**:
```bash
go run main.go
```

The server will start on `http://localhost:1212`

## API Endpoints

### Send Email
```
POST /email/invoice
Content-Type: application/json

{
  "Sender": "sender@example.com",
  "RecipientEmail": "recipient@example.com",
  "Body": "Email body content",
  "subject": "Email Subject",
  "InvoiceId": "uuid-string",
  "InvoiceType": 0,
  "retry": 0
}
```

**Response:** `200 OK` - Data received successfully!

### Health Check
```
GET /health
```

**Response:** `200 OK`

## Project Structure

```
PM_Golang/
├── apis/              # HTTP endpoints and request handlers
├── DTO/               # Data Transfer Objects (EmailInfo)
├── Services/          # Business logic (SendEmail service)
├── repository/        # Database operations
├── rabbitmq/          # Message queue handling
├── logger/            # Logging configuration
├── infrastructure/    # Infrastructure setup
├── main.go            # Application entry point
└── go.mod             # Go module definition
```

## Configuration

### RabbitMQ Connection
The service connects to RabbitMQ at `amqp://guest:guest@localhost:5672` by default.
This can be configured in the RabbitMQ functions.

### Email Provider
The service uses [Resend](https://resend.com) as the email provider. Configure your API key in the `.env` file.

### Logging
The application uses [zerolog](https://github.com/rs/zerolog) for structured logging. All operations are logged with appropriate levels (Info, Error, Fatal).

## Testing

Run the test suite:
```bash
go test ./...
```

Individual test files:
- `main_test.go` - Main application tests
- `apis/email_test.go` - API endpoint tests
- `Services/SendEmail_test.go` - Email service tests
- `repository/emailRepo_test.go` - Repository layer tests
- `rabbitmq/sendmessagequeue_test.go` - Queue processing tests

## Dependencies

- **github.com/resend/resend-go/v2** - Email sending API client
- **github.com/rabbitmq/amqp091-go** - RabbitMQ client
- **github.com/rs/zerolog** - Structured logging
- **github.com/joho/godotenv** - Environment file management
- **github.com/google/uuid** - UUID generation

## Architecture

### Request Flow
1. HTTP POST request arrives at `/email/invoice`
2. Request is validated and parsed
3. Email data is published to RabbitMQ queue
4. RabbitMQ consumer processes the message asynchronously
5. Email is sent via Resend API
6. Failed emails are stored in the database with retry count
7. Automatic retry mechanism attempts to resend failed emails

### Components

- **APIs Layer** - Handles HTTP requests and responses
- **Services Layer** - Contains business logic for email sending
- **Repository Layer** - Manages database operations for failed emails
- **RabbitMQ Layer** - Handles asynchronous message processing and retry logic
- **Logger** - Centralized logging for debugging and monitoring

## Error Handling

The service implements:
- Request validation with JSON error responses
- Structured error logging
- Automatic retry for failed emails
- Database storage for permanently failed emails

## Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `Resend_api_key` | API key for Resend email service | Yes |

## Performance Considerations

- **RabbitMQ QoS (Quality of Service)** set to 5 - This limits the consumer to process a maximum of 5 messages concurrently, not the queue size itself. The queue can grow indefinitely.
- **No queue size limit** currently configured - Messages are only limited by available RabbitMQ server memory and disk space
- Asynchronous processing prevents blocking HTTP responses
- Database persistence for failed email tracking
- Automatic retry mechanism reduces manual intervention

## License

This project is proprietary. All rights reserved.

## Author

Created by Oussama Wahbi

## Contact

For questions or issues, please contact me.
