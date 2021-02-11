package mailer

// TemplatedMailerConnector is the interface to send mails from the SendinBlue connector.
type TemplatedMailerConnector interface {
	Send(m *Message) error
	CreateEmailMessage(addressList []*Address, templateId int64, params map[string]string) *Message
}
