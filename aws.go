// Package ses_go is a simple wrapper around Amazon's AWS library
// that provides a function to send email with SES.
package ses_go

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Session that is used. Should be initialized with Init().
var sess *session.Session

// Initializes aws session with specified key id, key secret and region.
// This function has to be called before sending an email.
// Returns error from aws library if session could not be created.
func Init(keyID, keySecret, region string) error {
	var err error
	sess, err = session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(keyID, keySecret, ""),
		Region:      aws.String(region),
	})

	return err
}
