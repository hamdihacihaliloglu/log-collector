# Log Collector Service

A lightweight and extensible log collection microservice designed for microservices architectures.  
Supports asynchronous log processing, fallback handling with Redis, and external notifications via Slack and Email.

---

##  Features

- Buffered log queue with concurrent worker pool
- Redis-based fallback queue when the in-memory queue is full
- Slack notification support with log level filtering
- Email notification support (SMTP/Mailpit) with HTML formatted emails
- Dockerized setup for easy deployment
- Configurable log level thresholds for Slack and Mail

---

##  Quick Start

```bash
git clone https://github.com/yourname/log-collector.git
cd log-collector
cp .env.example .env
docker compose up --build
```
Access Mailpit UI at http://localhost:8025
```bash
Environment Variables (.env)
Variable	Description
SLACK_ENABLED	Enable/disable Slack notifications
SLACK_WEBHOOK_URL	Slack Incoming Webhook URL
SLACK_LOG_LEVEL	Minimum log level to send to Slack (info, warning, error)
EMAIL_ENABLED	Enable/disable Email notifications
EMAIL_SMTP	SMTP server address (use mailpit for local testing)
EMAIL_PORT	SMTP port (typically 587)
EMAIL_USER	SMTP username (leave blank for Mailpit)
EMAIL_PASS	SMTP password (leave blank for Mailpit)
EMAIL_FROM	Sender email address
EMAIL_TO	Receiver email address
MAIL_LOG_LEVEL	Minimum log level to send emails (info, warning, error)
DB_HOST	MySQL database hostname
DB_PORT	MySQL port (default 3306)
DB_USER	MySQL username
DB_PASS	MySQL password
DB_NAME	Database name
```
---
##  API Endpoints
POST /logs

Submit a log entry.
```bash
POST http://localhost:8080/logs
Content-Type: application/json

{
  "service": "imagecontrol",
  "level": "error",
  "message": "Failed to fetch CDN image",
  "timestamp": "2025-04-26T00:00:00Z"
}
```
GET /logs

Retrieve all stored log entries.
GET /health

Simple health check endpoint.


---
## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

