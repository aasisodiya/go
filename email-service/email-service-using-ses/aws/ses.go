package aws

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	// DefaultRegion is Default Region
	DefaultRegion = "us-west-2"
	// CharSet is the character encoding for the email.
	CharSet = "UTF-8"
)

// validateRegion function checks if region has been provided or not, if its not provided then it assign default value
func validateRegion(region string) string {
	if len(strings.ReplaceAll(region, " ", "")) < 1 {
		return DefaultRegion
	}
	return region
}

// ListValidEmails method gives you the list of all validated emails (It will exclude emails with pending status)
func ListValidEmails(region string) ([]string, error) {
	region = validateRegion(region)
	var listValidEmails []string
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	// Create SES service client
	svc := ses.New(sess)

	result, err := svc.ListIdentities(&ses.ListIdentitiesInput{IdentityType: aws.String("EmailAddress")})
	// Sample Data in result
	// {
	// 	Identities: [
	// 	 	 	"akash_sisodiya@test.com",
	// 		]
	// }
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, email := range result.Identities {
		// Use *email to get its value
		// result.Identities is non-interface type []*string
		var e = []*string{email}

		verified, err := svc.GetIdentityVerificationAttributes(&ses.GetIdentityVerificationAttributesInput{Identities: e})
		// Sample Data in verified
		// {
		// 	VerificationAttributes: {
		// 	  akash_sisodiya@test.com: {
		// 		VerificationStatus: "Success"
		// 	   }
		// 	 }
		// }
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		for _, va := range verified.VerificationAttributes {
			if *va.VerificationStatus == "Success" {
				// fmt.Println(*email)
				listValidEmails = append(listValidEmails, *email)
			}
		}
	}
	return listValidEmails, nil
}

// ListAllEmails method gives you the list of all emails (It will include emails with pending status)
func ListAllEmails(region string) ([]string, error) {
	region = validateRegion(region)
	var listEmails []string
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	// Create SES service client
	svc := ses.New(sess)

	result, err := svc.ListIdentities(&ses.ListIdentitiesInput{IdentityType: aws.String("EmailAddress")})
	// Sample Data in result
	// {
	// 	Identities: [
	// 	 	 	"akash_sisodiya@test.com",
	// 		]
	// }
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for _, email := range result.Identities {
		listEmails = append(listEmails, *email)
	}
	return listEmails, nil
}

// VerifyEmailID method sends an verification mail to given EmailID for Verification
func VerifyEmailID(email string, region string) (bool, error) {
	region = validateRegion(region)
	// Create a new session in the us-west-2 region.
	// Replace us-west-2 with the AWS Region you're using for Amazon SES.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	// Create an SES session.
	svc := ses.New(sess)

	// Attempt to send the email.
	_, err = svc.VerifyEmailAddress(&ses.VerifyEmailAddressInput{EmailAddress: aws.String(email)})

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return false, err
	}

	fmt.Println("Verification sent to address: " + email)
	return true, nil
}

// SendEmail function to send an email to given receiver address with given message, subject and from given sender
func SendEmail(Recipient string, Sender string, Subject string, TextBody string, HtmlBody string, region string) (bool, error) {
	region = validateRegion(region)
	// Create a new session in the us-west-2 region.
	// Replace us-west-2 with the AWS Region you're using for Amazon SES.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	// Create an SES session.
	svc := ses.New(sess)
	var input *ses.SendEmailInput

	// Assemble the email.
	if(HtmlBody == "") {
		input = &ses.SendEmailInput{
			Destination: &ses.Destination{
				CcAddresses: []*string{},
				ToAddresses: []*string{
					aws.String(Recipient),
				},
			},
			Message: &ses.Message{
				Body: &ses.Body{
					Text: &ses.Content{
						Charset: aws.String(CharSet),
						Data:    aws.String(TextBody),
					},
				},
				Subject: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(Subject),
				},
			},
			Source: aws.String(Sender),
			// Uncomment to use a configuration set
			//ConfigurationSetName: aws.String(ConfigurationSet),
		}
	} else {
		input = &ses.SendEmailInput{
			Destination: &ses.Destination{
				CcAddresses: []*string{},
				ToAddresses: []*string{
					aws.String(Recipient),
				},
			},
			Message: &ses.Message{
				Body: &ses.Body{
					Html: &ses.Content{
						Charset: aws.String(CharSet),
						Data:    aws.String(HtmlBody),
					},
				},
				Subject: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(Subject),
				},
			},
			Source: aws.String(Sender),
			// Uncomment to use a configuration set
			//ConfigurationSetName: aws.String(ConfigurationSet),
		}
	}

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				fmt.Println(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				fmt.Println(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				fmt.Println(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

		return false, err
	}

	fmt.Println("Email Sent to address: " + Recipient)
	fmt.Println(result)
	return true, nil
}

// RemoveEmail method will remove the email from email list
func RemoveEmail(Recipient string, region string) (bool, error) {
	region = validateRegion(region)
	// Create a new session in the us-west-2 region
	// Replace us-west-2 with the AWS Region you're using for Amazon SES
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	if err != nil {
		fmt.Println("Got error creating SES session:")
		fmt.Println(err.Error())
		return false, err
	}

	// Create an SES session
	svc := ses.New(sess)

	// Remove email address
	_, delErr := svc.DeleteVerifiedEmailAddress(&ses.DeleteVerifiedEmailAddressInput{EmailAddress: aws.String(Recipient)})

	// Display error message if it occurs
	if delErr != nil {
		fmt.Println("Got error attempting to remove email address: " + Recipient)
		fmt.Println(delErr.Error())
		return false, delErr
	}

	// Display success message
	fmt.Println("Removed email address: " + Recipient)
	return true, nil
}

// SESStatistics method gives you SES Stats
func SESStatistics(region string) error {
	region = validateRegion(region)
	// Initialize a session that the SDK uses to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and configuration from the shared configuration file ~/.aws/config.
	// sess := session.Must(session.NewSessionWithOptions(session.Options{
	//     SharedConfigState: session.SharedConfigEnable,
	// }))
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	// Create an SES session.
	svc := ses.New(sess)

	// Attempt to send the email.
	result, err := svc.GetSendStatistics(nil)
	fmt.Println(result)
	// Display any error message
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	dps := result.SendDataPoints

	fmt.Println("Got", len(dps), "datapoints")
	fmt.Println("")

	for _, dp := range dps {
		fmt.Println("Timestamp: ", dp.Timestamp)
		fmt.Println("Attempts:  ", aws.Int64Value(dp.DeliveryAttempts))
		fmt.Println("Bounces:   ", aws.Int64Value(dp.Bounces))
		fmt.Println("Complaints:", aws.Int64Value(dp.Complaints))
		fmt.Println("Rejects:   ", aws.Int64Value(dp.Rejects))
		fmt.Println("")
	}
	return nil
}
