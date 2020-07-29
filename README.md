# SES Go
Simple library for sending emails with AWS SES.

## Usage
```go
package main

import (
    "github.com/DzinVision/ses-go"
    "log"
)

const (
    keyID     = "YOUR_AWS_KEY_ID"
    keySecret = "YOUR_KEY_SECRET"
    region    = "AWS_REGION"
    
    sender    = "sender@mail.com"
    recipient = "recipient@mail.com"
    subject   = "subject"
    content   = "<h2>content</h2>"
)

func main() {
    // Initialize AWS session. 
    err := ses_go.Init(keyID, keySecret, region)
    if err != nil {
        log.Fatal(err)    
    }   
    
    // Send email.
    awsID, err := ses_go.SendHTMLEmail(sender, recipient, subject, content)
    if err != nil {
        log.Fatal(awsID)
    } 
}
```