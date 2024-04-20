package gateway

//go:generate mockery
type MailerGW interface {
	Send(recipient, subject, content string) error
}
