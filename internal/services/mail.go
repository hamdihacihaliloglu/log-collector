package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
	"log-collector/internal/config"
	"log-collector/internal/models"
	"log-collector/internal/logging"
)
func SendMail(entry models.LogEntry) {
	if !config.IsEmailEnabled() {
		log.Println("[Mail] Disabled via config")
		return
	}

	minLevel := config.GetMailLogLevel()
	if !logging.ShouldLog(entry.Level, minLevel) {
		log.Printf("[Mail] Skipped %s log below threshold: %s\n", entry.Service, entry.Level)
		return
	}

	from := os.Getenv("EMAIL_FROM")
	to := os.Getenv("EMAIL_TO")
	smtpHost := os.Getenv("EMAIL_SMTP")
	smtpPort := os.Getenv("EMAIL_PORT")

	subject := fmt.Sprintf("Log Alert: [%s] %s", strings.ToUpper(entry.Level), entry.Service)

	body := fmt.Sprintf(`
	<html>
	<body style="font-family: Arial, sans-serif;">
		<h2 style="color: #333;">Log Notification</h2>
		<p><strong>Service:</strong> %s</p>
		<p><strong>Level:</strong> %s</p>
		<p><strong>Message:</strong> %s</p>
		<p><strong>Timestamp:</strong> %s</p>
	</body>
	</html>
	`, entry.Service, strings.ToUpper(entry.Level), entry.Message, entry.Timestamp)

	message := []byte(fmt.Sprintf("Subject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s", subject, body))
	addr := smtpHost + ":" + smtpPort

	username := os.Getenv("EMAIL_USER")
	password := os.Getenv("EMAIL_PASS")

	var auth smtp.Auth
	if username != "" && password != "" {
		auth = smtp.PlainAuth("", username, password, smtpHost)
	}

	var err error
	if auth != nil {
		err = smtp.SendMail(addr, auth, from, []string{to}, message)
	} else {
		err = smtp.SendMail(addr, nil, from, []string{to}, message)
	}

	if err != nil {
		log.Println("[Mail] Send error:", err)
		return
	}

	log.Println("[Mail] Mail successfully sent for log entry:", entry.Service)
}
