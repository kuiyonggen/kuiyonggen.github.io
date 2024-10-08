package main

import (
    "fmt"
    "os"
    "net/smtp"
)

func main() {

    from := os.Getenv("FROM")

    password := os.Getenv("PASSWORD")

    to := []string{ os.Getenv("TO") }

    smtpHost := os.Getenv("SMTPHOST")

    smtpPort := os.Getenv("SMTPPORT")

    msg := []byte("This is a test email message.")

    auth := smtp.PlainAuth("", from, password, smtpHost)

    err := smtp.SendMail(smtpHost + smtpPort, auth, from, to, msg)

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("Email sent successfully")
}
