package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

/*
env GOOS=linux go build -o controller
chmod 0777 -R ./controller
*/
func executeSendMail(From, sendUserName, SMTP_PASS, Host, To, subject, body, mailtype string) error {
	hp := strings.Split(Host, ":")
	auth := smtp.PlainAuth("", From, SMTP_PASS, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + To + "\r\nFrom: " + sendUserName + "<" + From + ">" + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(To, ";")
	err := smtp.SendMail(Host, auth, From, send_to, msg)
	return err
}

func sendMail() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	From := os.Getenv("FROM")
	SMTP_PASS := os.Getenv("SMTP_PASS")
	Host := os.Getenv("HOST")
	To := os.Getenv("TO")
	body := readHTML()
	subject := "Puppeteer rebooted successfully."

	sendUserName := "EMAIL SENT BY GOLANG"
	{
		err := executeSendMail(From, sendUserName, SMTP_PASS, Host, To, subject, body, "html")
		if err != nil {
			fmt.Println("Send mail error!")
			fmt.Println(err)
		} else {
			fmt.Println("Send mail success!")
			writeToLogFile("Successfully sent email.")
		}
	}

}

func readHTML() string {
	b, err := ioutil.ReadFile("./index.html")
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	return str
}
