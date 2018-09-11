package email

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type SESEmailSender struct {
	session *session.Session
	ses     *ses.SES
	charset string
}

func NewSESEmailSender(credentials *credentials.Credentials) (*SESEmailSender, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: credentials,
	})
	if err != nil {
		return nil, err
	}
	svc := ses.New(sess)
	return &SESEmailSender{session: sess, ses: svc, charset: "UTF-8"}, nil
}

func (s SESEmailSender) Send(settings MessageSettings, content string) error {
	emailInput := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(settings.recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Text: &ses.Content{
					Charset: aws.String(s.charset),
					Data:    aws.String(content),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(s.charset),
				Data:    aws.String(settings.subject),
			},
		},
		Source: aws.String(settings.sender),
	}

	_, err := s.ses.SendEmail(emailInput)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				return fmt.Errorf("error code:'%s',message:'%s'", ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				return fmt.Errorf("error code:'%s',message:'%s'", ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				return fmt.Errorf("error code:'%s',message:'%s'", ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				return fmt.Errorf("error code:'%s',message:'%s'", "unknown", aerr.Error())
			}
		}
		return err
	}
	return nil
}
