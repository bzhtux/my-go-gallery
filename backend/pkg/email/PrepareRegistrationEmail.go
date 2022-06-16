package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
)

func PrepareRegistrationEmail(token, fname, lname, email string) bytes.Buffer {
	t, _ := template.ParseFiles("/Users/yfoeillet/go/src/github.com/bzhtux/my-go-gallery/backend/templates/registration_email.html")
	var body bytes.Buffer
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: Confirm your email \n%s\n\n", mimeHeaders)))
	t.Execute(&body, struct {
		FirstName         string
		LastName          string
		BaseURL           string
		Email             string
		RegistrationToken string
	}{
		FirstName:         fname,
		LastName:          lname,
		BaseURL:           "http://0.0.0.0:8080",
		Email:             email,
		RegistrationToken: token,
	})

	log.Println(&body)
	return body
}
