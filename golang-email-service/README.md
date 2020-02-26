# AWS SES Vs SendGrid Service

## Comparison

Each tool is targeted to a different audience. SendGrid is more of an all-inclusive tool, suitable for sending all kinds of emails, while Amazon SES does mainly transactional emails (but it does them really well). SES works with pay-as-you-go pricing, while SendGrid offers pre-designed plans.

* SES is cheaper compared to SendGrid

    |Emails sent / month|50,000|200,000|1,000,000|
    |-------------------|------|-------|---------|
    |SendGrid Email API |$14.95|$79.95 |$399.95  |
    |Amazon SES w/o discounts|$5|$20   |$100     |

* SES is only for transactional mails, whereas SendGrid is for all purpose mails

## Choose Amazon SES if you…

* Already work with AWS and are familiar with an environment
* Have an application hosted on EC2 or are considering moving there
* Need a reliable tool just for sending lots of emails and can live without all the additional features
* Are looking for the most cost-effective solution and have developers ready to work with this tool

## Choose SendGrid if you…

* Need to send all kinds of emails – newsletters, transactional messages, and everything in-between
* Require advanced analytical tools to see the progress of your campaigns and improve them on the go
* Want non-technical team members to be able to handle the tool with little to no involvement of the development team
* Have doubts how to integrate your application with Amazon SES and SendGrid has a library built specifically for this purpose*

---

[Source](https://blog.mailtrap.io/amazon-ses-vs-sendgrid/)