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

---

##  Troubleshooting

### 1. `/usr/bin/env: 'bash\r': No such file or directory`

**Cause:**  
This error occurs when a `.sh` script file is saved with Windows-style line endings (`CRLF`) instead of Unix-style (`LF`).  
Linux-based Docker containers require `LF` line endings.

**Solution:**  
- Open the `.sh` file (such as `wait-for-it.sh`) in your text editor.
- Change line endings from `CRLF` to `LF`.
- Save the file.

Alternatively, you can fix it via terminal:

```bash
dos2unix wait-for-it.sh
```
If dos2unix is not installed, you can install it:
```bash
sudo apt install dos2unix
```
### 2. Redis or MySQL Connection Errors

Cause:
Docker services may not be fully ready when the application tries to connect.

Solution:

    Ensure all Docker services (db, redis, app) are properly started.

    Run:
```bash
docker-compose up --build
```
to rebuild and start everything cleanly.
### 3. Slack or Email Notifications Not Working

Cause:
Missing or incorrect environment variables.

Solution:

    Check your .env file.

    Make sure variables like SLACK_WEBHOOK_URL, EMAIL_SMTP, EMAIL_USER, and EMAIL_PASS are properly set.

    If you don't want to use notifications, you can disable them:
```bash
SLACK_ENABLED=false
EMAIL_ENABLED=false
```
