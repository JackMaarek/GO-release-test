package mailer

// SBDetails contains the necessary informations to connect to the SendinBlue api.
type SBDetails struct {
	Url         string
	ApiKey      string
	SenderEmail string
	SenderName  string
}

// Address describes an email address
type Address struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Message describes an email message that should be sent
type Message struct {

	// Sender takes an Address which is sending the email message.  REQUIRED.
	Sender *Address `json:"sender"`

	// To takes a list of Address as recipients of the email.
	To []*Address `json:"to,omitempty"`

	// Headers is the list of email headers which should be sent with the email message.
	Headers map[string]string `json:"headers,omitempty"`

	// TemplateID takes the template stored in SendinBlue as the content of the email.
	TemplateID int64 `json:"templateId,omitempty"`

	// Params is the list of parameters which should be used in the template.
	Params map[string]string `json:"params,omitempty"`

	// Tags are arbitrary labels which are applied to this email in order to ease organizational operations in SendInBlue
	Tags []string `json:"tags,omitempty"`
}
