package mailsender

import (
        "log"
        "net/smtp"
        "fmt"
        "strconv"
)

type EmailUser struct {
    Username    string
    Password    string
    EmailServer string
    Port        int
}

type SmtpTemplateData struct {
    From    string
    To      string
    Subject string
    Body    string
}


func Sendmail(from string, to []string, subject string, body string) string {

	emailUser := &EmailUser{"ZlodimirBot@gmail.com", "ZlodimirBotPassword", "smtp.gmail.com", 587}
	fmt.Println("Preparing auth object...")
        auth := smtp.PlainAuth("", emailUser.Username, emailUser.Password, emailUser.EmailServer)
	
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n";
	subj := "Subject: " +  subject + "\n"
	msg := []byte(subj + mime + "<html><body><h3>" + body + "</h3></body></html>")
        // and send the email all in one step.
        err := smtp.SendMail(emailUser.EmailServer + ":" + strconv.Itoa(emailUser.Port), auth, from,
                to,
               	msg)
        if err != nil {
                log.Fatal(err)
		fmt.Println("Error sending message...")
	} else {
		fmt.Println("Message sent...")
        }
	
	return "Message sent ..."
}


