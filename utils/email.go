package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/alonelegion/api_golang_mongodb/config"
	"github.com/alonelegion/api_golang_mongodb/models"
	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL       string
	FirstName string
	Subject   string
}

func SendEmail(user *models.DBResponse, data *EmailData, temp *template.Template, templateName string) error {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load config", err)
	}

	// Sender data
	from := config.EmailFrom
	smtpUser := config.SMTPUser
	smtpPass := config.SMTPPass
	to := user.Email
	smtpHost := config.SMTPHost
	smtpPort := config.SMTPPort

	var body bytes.Buffer
	if err := temp.ExecuteTemplate(&body, templateName, &data); err != nil {
		log.Fatal("could not execute template", err)
	}

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("test/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send email
	if err := d.DialAndSend(m); err != nil {
		return nil
	}

	return nil
}

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	fmt.Println("Am parsing templates...")
	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}
