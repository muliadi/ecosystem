package ecosystem

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"
)

//sendEmail is used internally by ECOSystem modules to send transactional emails
func sendEmail(host string, port string, userName string, password string, to []string, subject string, data map[string]string, templateToUse string) (err error) {

	//Prepare the date for the email template
	parameters := struct {
		From    string
		To      string
		Subject string
		Data    map[string]string
	}{
		userName,
		strings.Join([]string(to), ","),
		subject,
		data,
	}

	buffer := new(bytes.Buffer)
	t, err := template.New(templateToUse).ParseGlob("templates/email/*")
	err = t.Execute(buffer, &parameters)

	auth := smtp.PlainAuth("", userName, password, host)

	err = smtp.SendMail(
		fmt.Sprintf("%s:%s", host, port),
		auth,
		userName,
		to,
		buffer.Bytes())

	return err
}
