package fileProcess

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"html/template"
	"net/smtp"
)

type Repository struct {
	DbClient *gorm.DB
	from     string
	password string
	smtpHost string
	smtpPort string
}

type RepositoryBase interface {
	EmailSender(emailBody []AdditionalData, userEmail string) error
}

func NewRepository(dbClient *gorm.DB, from, password, smtHost, smtPort string) RepositoryBase {
	return Repository{
		DbClient: dbClient,
		from:     from,
		password: password,
		smtpHost: smtHost,
		smtpPort: smtPort,
	}
}

func (r Repository) EmailSender(emailBody []AdditionalData, userEmail string) error {
	t, errTemplate := template.ParseFiles("./action/fileProcess/docs/emailTemplate.html")
	if errTemplate != nil {
		return errTemplate
	}

	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(
		fmt.Sprintf("From: %s\\r\\nTo: %s\\r\\nSubject: This is for filesProcessor testing subject \n%s\n\n",
			r.from, userEmail, mimeHeaders)))

	errExecute := t.Execute(&body, emailBody)
	if errExecute != nil {
		fmt.Println(errExecute)
		return errExecute
	}
	auth := smtp.PlainAuth("", r.from, r.password, r.smtpHost)
	err := smtp.SendMail(r.smtpHost+":"+r.smtpPort, auth, r.from, []string{userEmail}, body.Bytes())
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent Successfully!")
	return nil
}

func (r Repository) DbSave(data UserTransaction) error {
	err := r.DbClient.Save(data).Error
	return err
}
