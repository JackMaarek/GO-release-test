package mailer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// SBConnector handles the email sending to SendinBlue API.
type SBConnector struct {
	SBDetails *SBDetails
}

// Send transmits the email message to SendInBlue
func (sbc *SBConnector) Send(m *Message) error {
	data, err := json.Marshal(m)
	if err != nil {
		return errors.New("failed to encode message: " + err.Error())
	}
	req, _ := http.NewRequest("POST", sbc.SBDetails.Url, bytes.NewReader(data))
	req.Header.Add("api-key", sbc.SBDetails.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to transmit message: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode > 202 {
		return fmt.Errorf("send failed: %d %s", resp.StatusCode, resp.Status)
	}
	return nil
}

// CreateEmailMessage create the prerequisite for the email message.
func (sbc *SBConnector) CreateEmailMessage(addressList []*Address, templateId int64, params map[string]string) *Message {
	message := Message{
		Sender: &Address{
			Name:  sbc.SBDetails.SenderName,
			Email: sbc.SBDetails.SenderEmail,
		},
		To:         addressList,
		Headers:    nil,
		TemplateID: templateId,
		Params:     params,
		Tags:       nil,
	}

	return &message
}
