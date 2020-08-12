package ses_go

import (
	"bytes"
	"github.com/aws/aws-sdk-go/service/ses"
	"gopkg.in/mail.v2"
)

// constructMail constructs new email object with specified data.
func constructMail(sender, recipient, subject, bodyType, body string) *mail.Message {
	msg := mail.NewMessage()
	msg.SetHeader("From", sender)
	msg.SetHeader("To", recipient)
	msg.SetHeader("Subject", subject)
	msg.SetBody(bodyType, body)

	return msg
}

// sendMail sends message `msg`.
// Returns AWS email ID and error. If email ID could not be got, empty string is returned.
func sendMail(msg *mail.Message) (string, error) {
	// Create SES object from session.
	svc := ses.New(sess)

	// Convert message to bytes.
	var emailRaw bytes.Buffer
	_, _ = msg.WriteTo(&emailRaw)

	// Create input instance for AWS.
	input := &ses.SendRawEmailInput{
		RawMessage: &ses.RawMessage{Data: emailRaw.Bytes()},
	}

	// Send email.
	res, err := svc.SendRawEmail(input)

	// Get AWS email ID.
	awsID := ""

	if res != nil {
		if res.MessageId != nil {
			awsID = *res.MessageId
		}
	}

	return awsID, err
}

// SendHTMLEmail sends email with HTML content to recipient from sender email address. The email will have
// specified subject. Content should be HTML text that is sent as the body of the email.
// Returns AWS email ID and error. If email ID could not be got, empty string is returned.
func SendHTMLEmail(sender, recipient, subject, htmlContent string) (string, error) {
	// Construct new message.
	msg := constructMail(sender, recipient, subject, "text/html", htmlContent)

	return sendMail(msg)
}

// SendHTMLEmailWithAttachment sends email with HTML content and attachment to recipient from sender email address.
// The email will have  specified subject. Content should be HTML text that is sent as the body of the email.
// `attachment` represents path to the attachment file.
// Returns AWS email ID and error. If email ID could not be got, empty string is returned.
func SendHTMLEmailWithAttachment(sender, recipient, subject, htmlContent, attachmentPath, attachmentName string) (string, error) {
	// Construct new message.
	msg := constructMail(sender, recipient, subject, "text/html", htmlContent)
	msg.Attach(attachmentPath, mail.Rename(attachmentName))

	return sendMail(msg)
}
