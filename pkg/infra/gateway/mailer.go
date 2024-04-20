package gateway

import "fmt"

type MailerMemory struct {
}

func NewMailerGatewayMemory() *MailerMemory {
	return &MailerMemory{}
}

func (MailerMemory) Send(ricipient, subject, content string) error {
	fmt.Println("Mailer gateway send", ricipient, subject, content)
	return nil
}
