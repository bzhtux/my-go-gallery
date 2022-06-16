package email

import (
	"os"
)

var (
	SMTPUSer     = os.Getenv("SMTP_USER")
	SMTPPassword = os.Getenv("SMTP_PASSWORD")
	SMTPHost     = os.Getenv("SMTP_HOST")
	SMTPPort     = os.Getenv("SMTP_PORT")
	SMTPAPIKey   = os.Getenv("SMTP_API_KEY")
)

// func SendRegistrationEmail(fname, lname, to, regtoken string) {
// 	var ctx context.Context
// 	cfg := sendinblue.NewConfiguration()
// 	//Configure API key authorization: api-key
// 	cfg.AddDefaultHeader("api-key", SMTPAPIKey)
// 	//Configure API key authorization: partner-key
// 	cfg.AddDefaultHeader("partner-key", SMTPAPIKey)

// 	sib := sendinblue.NewAPIClient(cfg)
// 	result, resp, err := sib.AccountApi.GetAccount(ctx)
// 	if err != nil {
// 		fmt.Println("Error when calling AccountApi->get_account: ", err.Error())
// 		return
// 	}
// 	fmt.Println("GetAccount Object:", result, " GetAccount Response: ", resp)

// 	// var m = sendinblue.SendSmtpEmailReplyTo{}
// 	// m.Email = to
// 	p := map[string]interface{}{
// 		"FirstName":          fname,
// 		"LastName":           lname,
// 		"BaseURL":            "http://0.0.0.0:8080/user/register/",
// 		"RegistrationTToken": regtoken,
// 		"Email":              to,
// 	}
// 	body := sendinblue.SendSmtpEmail{
// 		TemplateId: 3,
// 		Sender: &sendinblue.SendSmtpEmailSender{
// 			Email: "no-reply@gallery.bzhtux-lab.net",
// 		},
// 		To: []sendinblue.SendSmtpEmailTo{
// 			{
// 				Email: to,
// 			},
// 		},
// 		ReplyTo: &sendinblue.SendSmtpEmailReplyTo{
// 			Name:  "Registration bot",
// 			Email: "no-reply@gallery.bzhtux-lab.net",
// 		},
// 		Params: map[string]interface{}{
// 			"FirstName":          fname,
// 			"LastName":           lname,
// 			"BaseURL":            "http://0.0.0.0:8080/user/register/",
// 			"RegistrationTToken": regtoken,
// 			"Email":              to,
// 		},
// 	}
// 	obj, resp, err := sib.TransactionalEmailsApi.SendTransacEmail(ctx, body)
// 	if err != nil {
// 		//
// 	}
// }
