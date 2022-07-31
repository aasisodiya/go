# Using SES with Golang

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-email-service.ses&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

## Getting Started

What I have done here is, compiled some easy to use functions for SES that will perform required operations

- List all Emails (Even the one with pending status)
- List Validated Emails
- Verify Email
- Send Email
- Remove Email
- Get SES Statistics

**Important:** You need to make sure that you have created proper roles with required access rights and assigned properly, may it be lambda/ EC2/ Local Machine.

## How to use the package

### Step 1: Select the AWS Region

Since some of the AWS services are region specific, first thing you need to do is finalize the region you ant to use for SES. Note down the region code and keep it handy as you will have to pass it along in every Function. In case you pass region as blank, default region (i.e us-west-2) will be considered

In example - main.go we have kept region as `us-west-2`

### Step 2: Import Package

Import package for aws with ses - here I have kept package name as aws (feel free to modify)

### Step 3: Call Required Function

Now using the imported package you can call any required function you want and execute the code

## Function Description

1. ListValidEmails Function

   ```golang
   func ListValidEmails(region string) ([]string, error)
   ```

   ListValidEmails Function gives you the list of all validated emails (It will exclude emails with pending status)

1. ListAllEmails Function

   ```golang
   func ListAllEmails(region string) ([]string, error)
   ```

   ListAllEmails Function gives you the list of all emails (It will include emails with pending status)

1. VerifyEmailID Function

   ```golang
   func VerifyEmailID(email string, region string) (bool, error)
   ```

   VerifyEmailID Function sends an verification mail to given EmailID for Verification, function will return true if there are no errors

1. SendEmail Function

   ```golang
   func SendEmail(Recipient string, Sender string, Subject string, TextBody string, HtmlBody string, region string) (bool, error)
   ```

   SendEmail function to send an email to given receiver address with given message, subject and from given sender, function will return true if there are no errors. You can even send HTML message in HtmlBody.

   > :red_circle: For emails which doesn't support HTML, TextBody will be displayed else HtmlBody :red_circle:

1. RemoveEmail Function

   ```golang
   func RemoveEmail(Recipient string, region string) (bool, error)
   ```

   RemoveEmail function will remove the email from email list, function will return true if there are no errors

1. SESStatistics Function

   ```golang
   func SESStatistics(region string) error
   ```

   SESStatistics Function gives you SES Stats like Timestamp, Attempts, Bounces, Complaints and Rejects

## Troubleshooting

- **Please Make Sure That You Use Correct AWS Region**
- **Access Denied:** Make sure to provide right User/roles/policy with required access to your EC2/Lambda/Local Machine
- **Email is not received:** Make sure your mail is verified if you are in sandbox environment
- **Can't send mail to other people:** AWS doesn't allow SES to send mail to unverified emails on Sandbox environment (which is default). But if you still want to send to other people without verifying the emails in SES then you will have to switch to AWS prod environment
<!--
<a href="http://www.youtube.com/watch?feature=player_embedded&v=Be2xHx0A2yI
" target="_blank"><img src="https://img.youtube.com/vi/Be2xHx0A2yI/0.jpg"
alt="IMAGE ALT TEXT HERE" width="240" height="180" border="10" /></a> -->

## Reference

1. [List Email](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/ses-example-list-emails.html)
2. [Send Verification](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/ses-example-send-verification.html)
3. [Send Email](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/ses-example-send-email.html)
4. [Delete Email](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/ses-example-delete-address.html)
5. [Get Statistics](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/ses-example-get-statistics.html)
6. [Monitor Usage Statistics](https://github.com/awsdocs/amazon-ses-developer-guide/blob/master/doc-source/monitor-usage-statistics-api.md)

[![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go&label=aasisodiya/go&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)](https://visitorbadge.io/status?path=aasisodiya.go)
