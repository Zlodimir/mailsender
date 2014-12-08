package main

import (
        "bytes"
        "log"
        "net/smtp"
        "fmt"
        "strconv"
	"html/template"
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

var doc bytes.Buffer
var err error

func main() {


	const emailTemplate = `From: &#123;&#123;.From&#125;&#125;
	To: &#123;&#123;.To&#125;&#125;
	Subject: &#123;&#123;.Subject&#125;&#125;
	&#123;&#123;.Body&#125;&#125;
	With the Best Regards,
	&#123;&#123;.From&#125;&#125;`
	
	context := &SmtpTemplateData{"ZlodimirBot", "sslupinos@gmail.com", "This is the e-mail subject line!", "Hello, this is a test e-mail body."	}

	t := template.New("emailTemplate")
	t, err = t.Parse(emailTemplate)
	if err != nil {
		log.Print("error trying to parse mail template")
		fmt.Println("Error trying ti parse mail template...")
	}
	err = t.Execute(&doc, context)
	if err != nil {
	    	log.Print("error trying to execute mail template")
		fmt.Println("Error trying to execute mail template...")
	}

        emailUser := &EmailUser{"ZlodimirBot@gmail.com", "ZlodimirBotPassword", "smtp.gmail.com", 587}
	fmt.Println("Preparing auth object...")
        auth := smtp.PlainAuth("", emailUser.Username, emailUser.Password, emailUser.EmailServer)

        // and send the email all in one step.
        err := smtp.SendMail(emailUser.EmailServer + ":" + strconv.Itoa(emailUser.Port), auth, "ZlodimirBot@gmail.com",
                []string{"sslupinos@gmail.com"},
               	doc.Bytes())
        if err != nil {
                log.Fatal(err)
		fmt.Println("Error sending message...")
	} else {
		fmt.Println("Message sent...")
        }

}


