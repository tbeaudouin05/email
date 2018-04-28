package sendemail

import (
	"log"
	"net/mail"
	"net/smtp"

	"github.com/scorredoira/email"
)

// SendEmail function is good
func SendEmail(attachPath string, toEmail []string, title string, message string, senderName string, senderEmail string, senderPw string, smtpAddr string, smtpPort string) {

	// compose the message
	m := email.NewMessage(title, message)
	m.From = mail.Address{Name: senderName, Address: senderEmail}
	m.To = toEmail

	// add attachments
	if attachPath != "NA" {
		if err := m.Attach(attachPath); err != nil {
			log.Fatal(err)
		}
	}

	// send it
	auth := smtp.PlainAuth("", senderEmail, senderPw, smtpAddr)
	if err := email.Send(smtpAddr+":"+smtpPort, auth, m); err != nil {
		log.Fatal(err)
	}
}
