package gateway

//go:generate mockery
type MailerGW interface {
	Send(recipient, subject, content string) error
}

//go:generate mockery
type CreditCardGW interface {
	ProcessPayment(token, amount string) error
}
