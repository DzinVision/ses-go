package ses_go

import (
	"github.com/aws/aws-sdk-go/service/ses"
)

// Sends email with HTML content to recipient from sender email address. The email will have
// specified subject. Content should be HTML text that is sent as the body of the email.
// Returns AWS email ID and error. If email ID could not be got, empty string is returned.
func SendHTMLEmail(sender, recipient, subject, htmlContent string) (string, error) {
	// Create SES object from session.
	svc := ses.New(sess)

	// Create input with specified parameters.
	input := &ses.SendEmailInput{
		Source: &sender,
		Destination: &ses.Destination{
			ToAddresses: []*string{&recipient},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: &htmlContent,
				},
			},
			Subject: &ses.Content{
				Data: &subject,
			},
		},
	}

	// Send email.
	res, err := svc.SendEmail(input)

	// Get AWS email ID.
	awsID := ""

	if res != nil {
		if res.MessageId != nil {
			awsID = *res.MessageId
		}
	}

	return awsID, err
}
