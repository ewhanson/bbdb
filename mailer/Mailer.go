package mailer

import (
	"errors"
	appConfig "github.com/ewhanson/bbdb/config"
	"github.com/google/uuid"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"time"
)

// Mailer handles construction and dispatching of plain text and HTML SMTP email
type Mailer struct {
	from         string
	to           []string
	subject      string
	plainTextMsg []string
	htmlMsg      []string
	config       *appConfig.Config
}

// New populates Mailer struct with initial data.
func New(to []string, subject string) (*Mailer, error) {
	configDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	config, err := appConfig.LoadConfig(configDir)
	if err != nil {
		return nil, errors.New("could not initialize appConfig")
	}

	m := &Mailer{
		from:    "Babygramz <noreply@babygramz.com>",
		to:      to,
		subject: subject,
		config:  &config,
	}

	return m, nil
}

// AddPlainTextMsg sets plain text version of message
func (m *Mailer) AddPlainTextMsg(msg []string) {
	m.plainTextMsg = msg
}

// AddHtmlMsg sets HTML version of message
func (m *Mailer) AddHtmlMsg(msg []string) {
	m.htmlMsg = msg
}

// Send delivers message and headers to SMTP server for dispatch
func (m *Mailer) Send() error {
	if m.plainTextMsg == nil && m.htmlMsg == nil {
		return errors.New("A Mailer must have a plain text or HTML message in order to be sent")
	}

	addr := m.config.SmtpHost + ":" + m.config.SmtpPort
	auth := smtp.PlainAuth("", m.config.SmtpUsername, m.config.SmtpPassword, m.config.SmtpHost)
	msg := m.composeMessage()

	err := smtp.SendMail(addr, auth, m.from, m.to, msg)
	if err != nil {
		log.Print("Unable to send email: ", err)
	}

	return err
}

// composeMessage adds all required headers and elements to a slice then inserts carriage return/line feeds between
// each element
func (m *Mailer) composeMessage() []byte {
	var msg []string

	msg = append(msg, "To: "+m.to[0])
	msg = append(msg, "From: "+m.from)
	msg = append(msg, "Date: "+time.Now().Format(time.RFC1123Z))
	msg = append(msg, "Subject: "+m.subject)
	msg = append(msg, "Message-Id: <"+strconv.Itoa(os.Getpid())+"+"+strconv.Itoa(int(time.Now().UnixMilli()))+"+"+uuid.New().String()+"@babygramz.com>")
	msg = append(msg, "MIME-Version: 1.0")
	msg = append(msg, "Content-Type: multipart/alternative; boundary=\"boundary-string\"")

	msg = append(msg, "")

	if m.plainTextMsg != nil {
		msg = append(msg, "--boundary-string")
		msg = append(msg, "Content-Type: text/plain; charset=\"utf-8\"")
		msg = append(msg, "Content-Transfer-Encoding: quoted-printable")
		msg = append(msg, "Content-Disposition: inline")

		msg = append(msg, "")

		for _, ptItem := range m.plainTextMsg {
			msg = append(msg, ptItem)
		}

		msg = append(msg, "")
	}

	if m.htmlMsg != nil {
		msg = append(msg, "--boundary-string")
		msg = append(msg, "Content-Type: text/html; charset=\"utf-8\"")
		msg = append(msg, "Content-Transfer-Encoding: quoted-printable")
		msg = append(msg, "Content-Disposition: inline")

		msg = append(msg, "")

		for _, htmlItem := range m.htmlMsg {
			msg = append(msg, htmlItem)
		}

		msg = append(msg, "")
	}

	msg = append(msg, "--boundary-string--")

	return []byte(strings.Join(msg, "\r\n"))
}
